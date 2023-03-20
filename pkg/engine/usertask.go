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
	err      error
}

func (ul *UserTaskListen) Listen(ctx workflow.Context) {
	claimCh := workflow.GetSignalChannel(ctx, UserTaskClaimChannel)
	reviewCh := workflow.GetSignalChannel(ctx, UserTaskReviewChannel)
	for {
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(claimCh, func(c workflow.ReceiveChannel, more bool) {
			var signal ent.CreateIdentityLinkInput
			c.Receive(ctx, &signal)
			signal.TaskID = ul.task.ID
			signal.OrgID = ul.task.OrgID
			ul.err = ul.exporter.ClaimUserTask(context.Background(), &signal)
			if ul.err != nil {
				ul.task.MemberCount++
				ul.task.UnfinishedCount++
			}
		})
		selector.AddReceive(reviewCh, func(c workflow.ReceiveChannel, more bool) {
			var signal ent.UpdateIdentityLinkInput
			c.Receive(ctx, &signal)
			signal.TaskID = &ul.task.ID
			ul.err = ul.exporter.Review(context.Background(), &signal, int(ul.task.UnfinishedCount))
			if ul.err != nil {
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
