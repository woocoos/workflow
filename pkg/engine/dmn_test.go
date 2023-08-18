package engine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/workflow/pkg/spec/vars"
	"testing"
)

func Test_crossProduct(t *testing.T) {
	type args struct {
		left  vars.Mapping
		right vars.Mapping
	}
	tests := []struct {
		name    string
		args    args
		want    []vars.Mapping
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "single",
			args: args{
				left:  vars.Mapping{"a": 1, "b": "2"},
				right: vars.Mapping{"c": 2, "d": "3"},
			},
			want: []vars.Mapping{
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
				left: vars.Mapping{
					"a": vars.Mapping{"b1": 1, "b2": 2},
					"b": vars.Mapping{"b1": 2, "b2": 3},
				},
				right: vars.Mapping{"c": 3, "d": 4},
			},
			want: []vars.Mapping{
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
				left: vars.Mapping{
					"a": vars.Mapping{"a1": 1, "a2": 2},
					"b": vars.Mapping{"a1": 2, "a2": 3},
				},
				right: vars.Mapping{
					"c": vars.Mapping{"b1": 3, "b2": 4},
					"d": vars.Mapping{"b1": 5, "b2": 6},
				},
			},
			want: []vars.Mapping{
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
