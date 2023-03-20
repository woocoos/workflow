package dmn

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpec_LoadDecisions(t *testing.T) {
	b := `
<definitions>
	<decision id="decision_1" name="decision_1">
		<decisionTable>
			<output label="decision_1">
				<outputValues>1</outputValues>
			</output>
		</decisionTable>
	</decision>
</definitions>
`
	var dmn *Definitions
	err := xml.Unmarshal([]byte(b), &dmn)
	require.NoError(t, err)
	assert.Len(t, dmn.Decisions, 1)
}

func TestSpec_LoadTypeRef(t *testing.T) {
	b := `
<definitions>
	<decision id="decision_1" name="decision_1">
		<decisionTable>
			 <input id="Input_1" label="Invoice Amount">
				<inputExpression id="InputExpression_1" typeRef="string">
					<text>amount</text>
				</inputExpression>
			</input>
			<input id="Input_1" label="Invoice Amount">
				<inputExpression id="InputExpression_1" typeRef="stringx">
					<text>amount</text>
				</inputExpression>
			</input>
			<output label="decision_1" typeRef="number">
				<outputValues>1</outputValues>
			</output>
		</decisionTable>
	</decision>
</definitions>
`
	var dmn *Definitions
	err := xml.Unmarshal([]byte(b), &dmn)
	assert.Equal(t, TypeRefNumber, dmn.Decisions[0].DecisionTable.Outputs[0].TypeRef)
	assert.Equal(t, TypeRefString, dmn.Decisions[0].DecisionTable.Inputs[0].InputExpression.TypeRef)
	assert.Equal(t, TypeRef("stringx"), dmn.Decisions[0].DecisionTable.Inputs[1].InputExpression.TypeRef)
	require.NoError(t, err)
}
