package engine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"testing"
)

func Test_crossProduct(t *testing.T) {
	type args struct {
		left  bpmn.Mappings
		right bpmn.Mappings
	}
	tests := []struct {
		name    string
		args    args
		want    []bpmn.Mappings
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "single",
			args: args{
				left:  bpmn.Mappings{"a": 1, "b": "2"},
				right: bpmn.Mappings{"c": 2, "d": "3"},
			},
			want: []bpmn.Mappings{
				{
					"a": 1,
					"b": "2",
					"c": 2,
					"d": "3",
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multi",
			args: args{
				left: bpmn.Mappings{
					"a": bpmn.Mappings{"b1": 1, "b2": 2},
					"b": bpmn.Mappings{"b1": 2, "b2": 3},
				},
				right: bpmn.Mappings{"c": 3, "d": 4},
			},
			want: []bpmn.Mappings{
				{
					"b1": 1,
					"b2": 2,
					"c":  3,
					"d":  4,
				},
				{
					"b1": 2,
					"b2": 3,
					"c":  3,
					"d":  4,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multi",
			args: args{
				left: bpmn.Mappings{
					"a": bpmn.Mappings{"a1": 1, "a2": 2},
					"b": bpmn.Mappings{"a1": 2, "a2": 3},
				},
				right: bpmn.Mappings{
					"c": bpmn.Mappings{"b1": 3, "b2": 4},
					"d": bpmn.Mappings{"b1": 5, "b2": 6},
				},
			},
			want: []bpmn.Mappings{
				{"a1": 1, "a2": 2, "b1": 3, "b2": 4},
				{"a1": 1, "a2": 2, "b1": 5, "b2": 6},
				{"a1": 2, "a2": 3, "b1": 3, "b2": 4},
				{"a1": 2, "a2": 3, "b1": 5, "b2": 6},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := crossMap(tt.args.left, tt.args.right)
			if !tt.wantErr(t, err, fmt.Sprintf("crossMap() error = %v, wantErr %v", tt.args.left, tt.args.right)) {
				return
			}
			assert.Equalf(t, tt.want, got, "crossMap() got = %v, want %v", tt.args.left, tt.args.right)
		})
	}
}
