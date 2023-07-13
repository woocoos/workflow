package bpmn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
