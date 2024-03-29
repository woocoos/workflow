// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/woocoos/workflow/codegen/entgen/schema"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/deployment"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/orguser"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/ent/task"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	decisiondefMixin := schema.DecisionDef{}.Mixin()
	decisiondefMixinHooks1 := decisiondefMixin[1].Hooks()
	decisiondefMixinHooks2 := decisiondefMixin[2].Hooks()
	decisiondef.Hooks[0] = decisiondefMixinHooks1[0]
	decisiondef.Hooks[1] = decisiondefMixinHooks2[0]
	decisiondefMixinInters2 := decisiondefMixin[2].Interceptors()
	decisiondef.Interceptors[0] = decisiondefMixinInters2[0]
	decisiondefMixinFields0 := decisiondefMixin[0].Fields()
	_ = decisiondefMixinFields0
	decisiondefMixinFields1 := decisiondefMixin[1].Fields()
	_ = decisiondefMixinFields1
	decisiondefFields := schema.DecisionDef{}.Fields()
	_ = decisiondefFields
	// decisiondefDescCreatedAt is the schema descriptor for created_at field.
	decisiondefDescCreatedAt := decisiondefMixinFields1[1].Descriptor()
	// decisiondef.DefaultCreatedAt holds the default value on creation for the created_at field.
	decisiondef.DefaultCreatedAt = decisiondefDescCreatedAt.Default.(func() time.Time)
	// decisiondefDescID is the schema descriptor for id field.
	decisiondefDescID := decisiondefMixinFields0[0].Descriptor()
	// decisiondef.DefaultID holds the default value on creation for the id field.
	decisiondef.DefaultID = decisiondefDescID.Default.(func() int)
	decisionreqdefMixin := schema.DecisionReqDef{}.Mixin()
	decisionreqdefMixinHooks1 := decisionreqdefMixin[1].Hooks()
	decisionreqdefMixinHooks2 := decisionreqdefMixin[2].Hooks()
	decisionreqdef.Hooks[0] = decisionreqdefMixinHooks1[0]
	decisionreqdef.Hooks[1] = decisionreqdefMixinHooks2[0]
	decisionreqdefMixinInters2 := decisionreqdefMixin[2].Interceptors()
	decisionreqdef.Interceptors[0] = decisionreqdefMixinInters2[0]
	decisionreqdefMixinFields0 := decisionreqdefMixin[0].Fields()
	_ = decisionreqdefMixinFields0
	decisionreqdefMixinFields1 := decisionreqdefMixin[1].Fields()
	_ = decisionreqdefMixinFields1
	decisionreqdefFields := schema.DecisionReqDef{}.Fields()
	_ = decisionreqdefFields
	// decisionreqdefDescCreatedAt is the schema descriptor for created_at field.
	decisionreqdefDescCreatedAt := decisionreqdefMixinFields1[1].Descriptor()
	// decisionreqdef.DefaultCreatedAt holds the default value on creation for the created_at field.
	decisionreqdef.DefaultCreatedAt = decisionreqdefDescCreatedAt.Default.(func() time.Time)
	// decisionreqdefDescID is the schema descriptor for id field.
	decisionreqdefDescID := decisionreqdefMixinFields0[0].Descriptor()
	// decisionreqdef.DefaultID holds the default value on creation for the id field.
	decisionreqdef.DefaultID = decisionreqdefDescID.Default.(func() int)
	deploymentMixin := schema.Deployment{}.Mixin()
	deploymentMixinHooks1 := deploymentMixin[1].Hooks()
	deploymentMixinHooks2 := deploymentMixin[2].Hooks()
	deployment.Hooks[0] = deploymentMixinHooks1[0]
	deployment.Hooks[1] = deploymentMixinHooks2[0]
	deploymentMixinInters2 := deploymentMixin[2].Interceptors()
	deployment.Interceptors[0] = deploymentMixinInters2[0]
	deploymentMixinFields0 := deploymentMixin[0].Fields()
	_ = deploymentMixinFields0
	deploymentMixinFields1 := deploymentMixin[1].Fields()
	_ = deploymentMixinFields1
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescCreatedAt is the schema descriptor for created_at field.
	deploymentDescCreatedAt := deploymentMixinFields1[1].Descriptor()
	// deployment.DefaultCreatedAt holds the default value on creation for the created_at field.
	deployment.DefaultCreatedAt = deploymentDescCreatedAt.Default.(func() time.Time)
	// deploymentDescDeployTime is the schema descriptor for deploy_time field.
	deploymentDescDeployTime := deploymentFields[3].Descriptor()
	// deployment.DefaultDeployTime holds the default value on creation for the deploy_time field.
	deployment.DefaultDeployTime = deploymentDescDeployTime.Default.(func() time.Time)
	// deploymentDescID is the schema descriptor for id field.
	deploymentDescID := deploymentMixinFields0[0].Descriptor()
	// deployment.DefaultID holds the default value on creation for the id field.
	deployment.DefaultID = deploymentDescID.Default.(func() int)
	identitylinkMixin := schema.IdentityLink{}.Mixin()
	identitylinkMixinHooks1 := identitylinkMixin[1].Hooks()
	identitylink.Hooks[0] = identitylinkMixinHooks1[0]
	identitylinkMixinInters1 := identitylinkMixin[1].Interceptors()
	identitylink.Interceptors[0] = identitylinkMixinInters1[0]
	identitylinkMixinFields0 := identitylinkMixin[0].Fields()
	_ = identitylinkMixinFields0
	identitylinkFields := schema.IdentityLink{}.Fields()
	_ = identitylinkFields
	// identitylinkDescID is the schema descriptor for id field.
	identitylinkDescID := identitylinkMixinFields0[0].Descriptor()
	// identitylink.DefaultID holds the default value on creation for the id field.
	identitylink.DefaultID = identitylinkDescID.Default.(func() int)
	orguserFields := schema.OrgUser{}.Fields()
	_ = orguserFields
	// orguserDescJoinedAt is the schema descriptor for joined_at field.
	orguserDescJoinedAt := orguserFields[2].Descriptor()
	// orguser.DefaultJoinedAt holds the default value on creation for the joined_at field.
	orguser.DefaultJoinedAt = orguserDescJoinedAt.Default.(func() time.Time)
	procdefMixin := schema.ProcDef{}.Mixin()
	procdefMixinHooks1 := procdefMixin[1].Hooks()
	procdefMixinHooks2 := procdefMixin[2].Hooks()
	procdefHooks := schema.ProcDef{}.Hooks()
	procdef.Hooks[0] = procdefMixinHooks1[0]
	procdef.Hooks[1] = procdefMixinHooks2[0]
	procdef.Hooks[2] = procdefHooks[0]
	procdefMixinInters2 := procdefMixin[2].Interceptors()
	procdef.Interceptors[0] = procdefMixinInters2[0]
	procdefMixinFields0 := procdefMixin[0].Fields()
	_ = procdefMixinFields0
	procdefMixinFields1 := procdefMixin[1].Fields()
	_ = procdefMixinFields1
	procdefFields := schema.ProcDef{}.Fields()
	_ = procdefFields
	// procdefDescCreatedAt is the schema descriptor for created_at field.
	procdefDescCreatedAt := procdefMixinFields1[1].Descriptor()
	// procdef.DefaultCreatedAt holds the default value on creation for the created_at field.
	procdef.DefaultCreatedAt = procdefDescCreatedAt.Default.(func() time.Time)
	// procdefDescID is the schema descriptor for id field.
	procdefDescID := procdefMixinFields0[0].Descriptor()
	// procdef.DefaultID holds the default value on creation for the id field.
	procdef.DefaultID = procdefDescID.Default.(func() int)
	procinstMixin := schema.ProcInst{}.Mixin()
	procinstMixinHooks1 := procinstMixin[1].Hooks()
	procinstMixinHooks2 := procinstMixin[2].Hooks()
	procinst.Hooks[0] = procinstMixinHooks1[0]
	procinst.Hooks[1] = procinstMixinHooks2[0]
	procinstMixinInters2 := procinstMixin[2].Interceptors()
	procinst.Interceptors[0] = procinstMixinInters2[0]
	procinstMixinFields0 := procinstMixin[0].Fields()
	_ = procinstMixinFields0
	procinstMixinFields1 := procinstMixin[1].Fields()
	_ = procinstMixinFields1
	procinstFields := schema.ProcInst{}.Fields()
	_ = procinstFields
	// procinstDescCreatedAt is the schema descriptor for created_at field.
	procinstDescCreatedAt := procinstMixinFields1[1].Descriptor()
	// procinst.DefaultCreatedAt holds the default value on creation for the created_at field.
	procinst.DefaultCreatedAt = procinstDescCreatedAt.Default.(func() time.Time)
	// procinstDescStartTime is the schema descriptor for start_time field.
	procinstDescStartTime := procinstFields[3].Descriptor()
	// procinst.DefaultStartTime holds the default value on creation for the start_time field.
	procinst.DefaultStartTime = procinstDescStartTime.Default.(func() time.Time)
	// procinstDescID is the schema descriptor for id field.
	procinstDescID := procinstMixinFields0[0].Descriptor()
	// procinst.DefaultID holds the default value on creation for the id field.
	procinst.DefaultID = procinstDescID.Default.(func() int)
	taskMixin := schema.Task{}.Mixin()
	taskMixinHooks1 := taskMixin[1].Hooks()
	task.Hooks[0] = taskMixinHooks1[0]
	taskMixinInters1 := taskMixin[1].Interceptors()
	task.Interceptors[0] = taskMixinInters1[0]
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescParentID is the schema descriptor for parent_id field.
	taskDescParentID := taskFields[5].Descriptor()
	// task.DefaultParentID holds the default value on creation for the parent_id field.
	task.DefaultParentID = taskDescParentID.Default.(int)
	// taskDescAgreeCount is the schema descriptor for agree_count field.
	taskDescAgreeCount := taskFields[10].Descriptor()
	// task.DefaultAgreeCount holds the default value on creation for the agree_count field.
	task.DefaultAgreeCount = taskDescAgreeCount.Default.(int32)
	// taskDescSequential is the schema descriptor for sequential field.
	taskDescSequential := taskFields[12].Descriptor()
	// task.DefaultSequential holds the default value on creation for the sequential field.
	task.DefaultSequential = taskDescSequential.Default.(bool)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[13].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[14].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskMixinFields0[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() int)
}

const (
	Version = "v0.12.4-0.20230702151415-1ec75238037c"           // Version of ent codegen.
	Sum     = "h1:63w0ILLHBEPwyeMO4en09BA8GLUwCwNfI+3C3MamOhY=" // Sum of ent codegen.
)
