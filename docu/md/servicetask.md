---
id: serviceTask
---

## ServiceTask

首先[参考文档](https://docs.camunda.io/docs/components/modeler/bpmn/service-tasks/).

### 任务定义

任务类型`type`需要指定以下:

#### `http`  流程处理内置了 http client,可以调用外部的 http 服务.

在header中以 key,value 形式定义http请求中要素:

| key         | 说明     | 值                   |
|-------------|--------|---------------------|
| url         | 目标地址   |                     |
| method      | 请求方法   | 默认 GET              |
| contentType | 请求内容类型 | 如: application/json |



