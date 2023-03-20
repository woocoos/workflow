package conv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EvaluateExpression(t *testing.T) {
	type args struct {
		expressions []string
		vars        map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "const value",
			args: args{
				expressions: []string{"=2"},
				vars: map[string]any{
					"value": 2,
				},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "string value",
			args: args{
				expressions: []string{"='abc'"},
				vars: map[string]any{
					"value": 2,
				},
			},
			want:    "abc",
			wantErr: assert.NoError,
		},
		{
			name: "bool",
			args: args{
				expressions: []string{"=value > 1", "=value>1"},
				vars: map[string]any{
					"value": 2,
				},
			},
			want:    true,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.args.expressions {
				gotDu, err := EvaluateExpression(s, tt.args.vars)
				if !tt.wantErr(t, err, fmt.Sprintf("EvaluateExpressionToDuration(%v)", tt.args.vars)) {
					return
				}
				assert.Equalf(t, tt.want, gotDu, "EvaluateExpressionToDuration(%v)", tt.args.vars)
			}
		})
	}
}

func TestEvaluateExpressionToInt(t *testing.T) {
	type args struct {
		exp  string
		vars any
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "int",
			args: args{
				exp: "=value",
				vars: map[string]any{
					"value": 2,
				},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "string-int",
			args: args{
				exp: "=value",
				vars: map[string]any{
					"value": "2",
				},
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "key charset not allow",
			args: args{
				exp: "=a.value",
				vars: map[string]any{
					"a.value": 2,
				},
			},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvaluateExpressionToInt(tt.args.exp, tt.args.vars)
			if !tt.wantErr(t, err, fmt.Sprintf("EvaluateExpressionToInt(%v, %v)", tt.args.exp, tt.args.vars)) {
				return
			}
			assert.Equalf(t, tt.want, got, "EvaluateExpressionToInt(%v, %v)", tt.args.exp, tt.args.vars)
		})
	}
}
