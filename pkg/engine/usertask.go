package engine

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/workflow"
)

const (
	UserTaskClaimChannel  = "ClaimChannel"
	UserTaskReviewChannel = "ReviewChannel"
)

// UserTaskListen listen user task signal
//
// task.MemberCount = instances count, task.UnfinishedCount = unfinished instances count, task.AgreeCount = agree count
type UserTaskListen struct {
	task     *ent.Task
	exporter Exporter
}

func (ul *UserTaskListen) Listen(ctx workflow.Context) {
	claimCh := workflow.GetSignalChannel(ctx, UserTaskClaimChannel)
	reviewCh := workflow.GetSignalChannel(ctx, UserTaskReviewChannel)
	logger := workflow.GetLogger(ctx)
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(claimCh, func(c workflow.ReceiveChannel, more bool) {
			var signal ent.CreateIdentityLinkInput
			c.Receive(ctx, &signal)
			signal.TaskID = ul.task.ID
			signal.TenantID = ul.task.TenantID
			added, err := ul.exporter.ClaimUserTask(context.Background(), &signal)
			if err != nil {
				logger.Error("user task receive error:", "error", err, "userID", *signal.UserID)
				return
			}
			if added {
				ul.task.MemberCount++
				ul.task.UnfinishedCount++
			}
		})
		selector.AddReceive(reviewCh, func(c workflow.ReceiveChannel, more bool) {
			var signal ent.UpdateIdentityLinkInput
			c.Receive(ctx, &signal)
			signal.TaskID = &ul.task.ID
			err := ul.exporter.Review(context.Background(), &signal, int(ul.task.UnfinishedCount))
			if err != nil {
				logger.Error("user task review error:", "error", err, "userID", *signal.UserID)
				return
			}
			ul.task.UnfinishedCount--
			switch *signal.OperationType {
			case identitylink.OperationTypePass:
				ul.task.AgreeCount++
			case identitylink.OperationTypeReject:
			}
		})
		selector.Select(ctx)
	}
}

// CheckComplete 检查任务是否完成. 如果任务有多实例, 则检查多实例的完成条件
func (ul *UserTaskListen) CheckComplete(ctx context.Context, taskEle *bpmn.UserTask) (cp bool, err error) {
	if mi := taskEle.MultiInstance; mi != nil {
		cp, err = CompletionCondition(mi.CompletionCondition.Content, ul.task.MemberCount, ul.task.UnfinishedCount, ul.task.AgreeCount)
		if err != nil {
			return false, err
		}
	}
	cp = ul.task.UnfinishedCount == 0
	return
}
