// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/workflow/api/graphql/model"
	"github.com/woocoos/workflow/codegen/entgen/types"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _WorkflowRun_id(ctx context.Context, field graphql.CollectedField, obj *types.WorkflowRun) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_WorkflowRun_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_WorkflowRun_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "WorkflowRun",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _WorkflowRun_runID(ctx context.Context, field graphql.CollectedField, obj *types.WorkflowRun) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_WorkflowRun_runID(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.RunID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_WorkflowRun_runID(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "WorkflowRun",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDeployDiagramInput(ctx context.Context, obj interface{}) (model.DeployDiagramInput, error) {
	var it model.DeployDiagramInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "appID", "resourceKey", "resourceID"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "appID":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("appID"))
			data, err := ec.unmarshalNID2int(ctx, v)
			if err != nil {
				return it, err
			}
			it.AppID = data
		case "resourceKey":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("resourceKey"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.ResourceKey = data
		case "resourceID":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("resourceID"))
			data, err := ec.unmarshalOID2ᚕintᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.ResourceID = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputStartProcessInput(ctx context.Context, obj interface{}) (model.StartProcessInput, error) {
	var it model.StartProcessInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"procDefKey", "businessKey", "variables"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "procDefKey":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("procDefKey"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ProcDefKey = data
		case "businessKey":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("businessKey"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.BusinessKey = data
		case "variables":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("variables"))
			data, err := ec.unmarshalOVariableInput2ᚕᚖgithubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐVariableInputᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Variables = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputVariableInput(ctx context.Context, obj interface{}) (model.VariableInput, error) {
	var it model.VariableInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "type", "value"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Type = data
		case "value":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputWorkflowRunInput(ctx context.Context, obj interface{}) (types.WorkflowRun, error) {
	var it types.WorkflowRun
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "runID"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "runID":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("runID"))
			data, err := ec.unmarshalOString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.RunID = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var workflowRunImplementors = []string{"WorkflowRun"}

func (ec *executionContext) _WorkflowRun(ctx context.Context, sel ast.SelectionSet, obj *types.WorkflowRun) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, workflowRunImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WorkflowRun")
		case "id":
			out.Values[i] = ec._WorkflowRun_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "runID":
			out.Values[i] = ec._WorkflowRun_runID(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNDeployDiagramInput2githubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐDeployDiagramInput(ctx context.Context, v interface{}) (model.DeployDiagramInput, error) {
	res, err := ec.unmarshalInputDeployDiagramInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNStartProcessInput2githubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐStartProcessInput(ctx context.Context, v interface{}) (model.StartProcessInput, error) {
	res, err := ec.unmarshalInputStartProcessInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNVariableInput2ᚖgithubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐVariableInput(ctx context.Context, v interface{}) (*model.VariableInput, error) {
	res, err := ec.unmarshalInputVariableInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNWorkflowRun2githubᚗcomᚋwoocoosᚋworkflowᚋcodegenᚋentgenᚋtypesᚐWorkflowRun(ctx context.Context, sel ast.SelectionSet, v types.WorkflowRun) graphql.Marshaler {
	return ec._WorkflowRun(ctx, sel, &v)
}

func (ec *executionContext) marshalNWorkflowRun2ᚖgithubᚗcomᚋwoocoosᚋworkflowᚋcodegenᚋentgenᚋtypesᚐWorkflowRun(ctx context.Context, sel ast.SelectionSet, v *types.WorkflowRun) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._WorkflowRun(ctx, sel, v)
}

func (ec *executionContext) unmarshalNWorkflowRunInput2githubᚗcomᚋwoocoosᚋworkflowᚋcodegenᚋentgenᚋtypesᚐWorkflowRun(ctx context.Context, v interface{}) (types.WorkflowRun, error) {
	res, err := ec.unmarshalInputWorkflowRunInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOVariableInput2ᚕᚖgithubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐVariableInputᚄ(ctx context.Context, v interface{}) ([]*model.VariableInput, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.VariableInput, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNVariableInput2ᚖgithubᚗcomᚋwoocoosᚋworkflowᚋapiᚋgraphqlᚋmodelᚐVariableInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// endregion ***************************** type.gotpl *****************************
