package bpmn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConditionExpression_Evaluate(t *testing.T) {
	type fields struct {
		exp string
	}
	type args struct {
		input map[string]any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "one element",
			fields: fields{
				exp: "=value = 1",
			},
			args: args{
				input: map[string]any{
					"value": 1,
				},
			},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name: "long",
			fields: fields{
				exp: "=value = 1 or key = 2",
			},
			args: args{
				input: map[string]any{
					"value": 1,
					"key":   2,
				},
			},
			want:    true,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ConditionExpression{}
			c.Content = tt.fields.exp
			got, err := c.Evaluate(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("Evaluate(%v)", tt.args.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Evaluate(%v)", tt.args.input)
		})
	}
}

func TestEvaluateExpressionToDuration(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		wantDu  time.Duration
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "duration",
			args: args{
				exp: "PT1H30M1S",
			},
			wantDu:  time.Duration(5401 * time.Second),
			wantErr: assert.NoError,
		},
		{
			name: "duration",
			args: args{
				exp: "EPT1S",
			},
			wantDu:  0,
			wantErr: assert.Error,
		},
		{
			name: "zero",
			args: args{
				exp: "PT0S",
			},
			wantDu:  0,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDu, err := EvaluateExpressionToDuration(tt.args.exp)
			if !tt.wantErr(t, err, fmt.Sprintf("EvaluateExpressionToDuration(%v)", tt.args.exp)) {
				return
			}
			assert.Equalf(t, tt.wantDu, gotDu, "EvaluateExpressionToDuration(%v)", tt.args.exp)
		})
	}
}
