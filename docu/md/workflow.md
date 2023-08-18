---
id: workflow
---

woocoos的流程引擎技术采用了标准化的BPMN2.0协议,结合新一代的工作流引擎Temporal,兼具业界标准及稳定性.

## 流程

流程是一系列的任务的集合,流程引擎根据流程定义,创建流程实例,并且根据流程定义的流转规则,执行流程实例,产生任务.

### BPMN2.0

BPMN的协议支持自定义扩展,我们对比了多家的格式协议,最终采用了基于BPMN2.0 Camunda 8的扩展方式,包括其比Camunda 7之前的扩展更加标准化.
因此在流程元素可以参与 Camunda 8的流程定义文档解释.

> 所有的定义及实现存在差异,因此并不是所有的定义都支持,请参考本文档的支持说明.

### 表达式

在Camunda 8以上的定义中,脚本化的方式已经很少见了,这符合我们对用户要求低的需求.
为了解决动态性,引入了[FEEL language](https://learn-dmn-in-15-minutes.com/learn/the-feel-language.html),但该表达式极易上手,对用户不会造成学习难度

### 流程编辑器

采用开源的Camunda Modeler,可以在[官网](https://camunda.com/download/modeler/)下载到对应的版本.

同时我们也提供Web的流程编辑器,可以通过`webui`访问.