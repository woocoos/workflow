// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/deployment"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/ent/task"
)

// GlobalID returns the global identifier for the given DecisionDef node.
func (dd *DecisionDef) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", decisiondef.Table, dd.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given DecisionReqDef node.
func (drd *DecisionReqDef) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", decisionreqdef.Table, drd.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given Deployment node.
func (d *Deployment) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", deployment.Table, d.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given IdentityLink node.
func (il *IdentityLink) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", identitylink.Table, il.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given ProcDef node.
func (pd *ProcDef) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", procdef.Table, pd.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given ProcInst node.
func (pi *ProcInst) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", procinst.Table, pi.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

// GlobalID returns the global identifier for the given Task node.
func (t *Task) GlobalID(context.Context) (string, error) {
	id := fmt.Sprintf("%s:%d", task.Table, t.ID)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

type ResolvedGlobal struct{ Type, ID string }

func FromGlobalID(s string) (*ResolvedGlobal, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	tid := strings.Split(string(b), ":")
	if len(tid) != 2 {
		return nil, fmt.Errorf("invalid global identifier format %q", b)
	}
	return &ResolvedGlobal{Type: tid[0], ID: tid[1]}, nil
}

// GlobalID returns the global identifier for the given type and id.
func GlobalID(tp, id string) (string, error) {
	switch tp {
	case decisiondef.Table:
		break
	case decisionreqdef.Table:
		break
	case deployment.Table:
		break
	case identitylink.Table:
		break
	case procdef.Table:
		break
	case procinst.Table:
		break
	case task.Table:
		break
	default:
		return "", fmt.Errorf("invalid type %q", tp)
	}
	id = fmt.Sprintf("%s:%s", tp, id)
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}

func (r *ResolvedGlobal) Int() (int, error) {
	return strconv.Atoi(r.ID)
}
func IntFromGlobalID(s string) (int, error) {
	r, err := FromGlobalID(s)
	if err != nil {
		return 0, err
	}
	return r.Int()
}

func (c *Client) NoderEx(ctx context.Context, id string) (Noder, error) {
	g, err := FromGlobalID(id)
	if err != nil {
		return nil, err
	}
	v, err := g.Int()
	if err != nil {
		return nil, err
	}
	return c.Noder(ctx, v, WithNodeType(func(ctx context.Context, i int) (string, error) {
		return g.Type, nil
	}))
}

func (c *Client) NodersEx(ctx context.Context, ids []string, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.NoderEx(ctx, ids[0])
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	for i, id := range ids {
		g, err := FromGlobalID(id)
		if err != nil {
			errors[i] = err
			continue
		}
		intID, err := g.Int()
		if err != nil {
			errors[i] = err
			continue
		}
		tables[g.Type] = append(tables[g.Type], intID)
		id2idx[intID] = append(id2idx[intID], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}