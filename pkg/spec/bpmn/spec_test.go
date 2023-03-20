package bpmn

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/workflow"
	"os"
	"testing"
	"time"
)

func TestSpec_LoadInvoice(t *testing.T) {
	// load a BPMN file from test directory
	bytes, err := os.ReadFile("../../../test/testdata/unittest.bpmn")
	require.NoError(t, err)
	// unmarshal the BPMN file
	var bpmn *Definitions
	err = xml.Unmarshal(bytes, &bpmn)
	assert.NoError(t, err)
}

func TestSpec_LoadExtensionElement(t *testing.T) {
	bytes := `<?xml version="1.0" encoding="UTF-8"?>
<bpmn:definitions>
	<bpmn:process id="Process_1" isExecutable="true">
		<bpmn:startEvent id="StartEvent_1">
			<bpmn:extensionElements>
				<zeebe:properties>
	                <zeebe:property name="taskQueue" value="tv" />
					<zeebe:property name="startToCloseTimeout" value="5s" />
				</zeebe:properties>
			</bpmn:extensionElements>
            <bpmn:outgoing>Flow_1y8jegt</bpmn:outgoing>
        </bpmn:startEvent>
    </bpmn:process>
</bpmn:definitions>
`
	// unmarshal the BPMN file
	var bpmn Definitions
	err := xml.Unmarshal([]byte(bytes), &bpmn)
	assert.NoError(t, err)
	eps := bpmn.Process[0].StartEvents[0].ExtensionProperties
	ao := workflow.ActivityOptions{}
	assert.NoError(t, eps.Decode(&ao))
	assert.Equal(t, "tv", ao.TaskQueue)
	assert.Equal(t, time.Second*5, ao.StartToCloseTimeout)
}

func TestSpec_FormatExpression(t *testing.T) {
	bytes := `<?xml version="1.0" encoding="UTF-8"?>
<bpmn:definitions>
	<bpmn:process id="Process_1" isExecutable="true">
		<bpmn:startEvent id="StartEvent_1">
            <bpmn:outgoing>Flow_1y8jegt</bpmn:outgoing>
        </bpmn:startEvent>
        <bpmn:exclusiveGateway id="Gateway_01wr5g0">
            <bpmn:incoming>Flow_1y8jegt</bpmn:incoming>
            <bpmn:outgoing>price-gt-zero</bpmn:outgoing> 
        </bpmn:exclusiveGateway>
        <bpmn:sequenceFlow id="Flow_1y8jegt" sourceRef="StartEvent_1" targetRef="Gateway_01wr5g0" />
        <bpmn:sequenceFlow id="price-gt-zero" name="price &#62; 0" sourceRef="Gateway_01wr5g0" targetRef="task-a">
            <bpmn:conditionExpression xsi:type="bpmn:tFormalExpression">price &gt; 0</bpmn:conditionExpression>
        </bpmn:sequenceFlow>
    </bpmn:process>
</bpmn:definitions>
`
	// unmarshal the BPMN file
	var bpmn Definitions
	err := xml.Unmarshal([]byte(bytes), &bpmn)
	assert.Equal(t, "price > 0", bpmn.Process[0].SequenceFlows[1].ConditionExpression.Content)
	assert.NoError(t, err)
}

func TestSpec_EventBaseMessage(t *testing.T) {
	bytes, err := os.ReadFile("../../../test/case/message-timer.bpmn")
	require.NoError(t, err)
	// unmarshal the BPMN file
	var bpmn Definitions
	err = xml.Unmarshal(bytes, &bpmn)
	assert.NoError(t, err)

}
