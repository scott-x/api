# api
### resful api 设计风格
REST api是一种设计标准，风格（不是任何架构标准），通常采用的是http协议，JSON作为数据格式

### resful api 特点
- 统一接口 （uniform interface）
- 无状态 （stateless）
- 可缓存 （cacheable）
- 分层 （layered system）
- CS模式（Client-server Atchitecture）

### api设计原则
- 以url（统一资源定位符） 风格设计api
- 通过不同的 method（GET POST PUT DELETE）来区分对资源的CRUD（操作）
- 返回码（STATUS CODE）必须符合http资源描述的规定