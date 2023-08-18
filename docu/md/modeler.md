---
id: modeler
---

## 模型编辑器

基本的操作可参考[Camunda Modeler](https://docs.camunda.io/docs/components/modeler/bpmn/)

本文重点描述差异化的实现.

## Task

一般任务对应为工作流中的Activity,由它来执行Task的逻辑.

### 任务的基本设置

任务的基本设置定义在`Extension properties`中,可根据不同的任务类型设置:

| code | 含义 | 默认值 |
|------|----|-----|
|      |    |     |
|      |    |     |


### 任务类型

[x] ServiceTask  
[x] UserTask  
[ ] ScriptTask
