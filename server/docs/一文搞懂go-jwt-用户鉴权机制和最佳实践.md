# 一文搞懂go-jwt-用户鉴权机制和最佳实践
我们在开发后台服务时，面对前端发送过来的请求需要鉴权，证明是某个用户的合法请求。在我们登录验证完密码后，用户就可以拿到jwt，后续的请求携带这个jwt就可以证明其身份的合法性，不需要反复查看数据库。这个jwt会被缓存在redis

## Jwt 
Jwt全称JSON Web Token，为了在网络应用环境中传递声明而执行的一种基于JSON的开放标准。Jwt由·翻分隔成三个部分：
- 头部 / Header
- 负载 / Payload
- 签名 / Signature

头部和负载以JSON的形式存在

## 基于JWT实现认证实践

JWT的token是一种refresh token，具有时效性。会话的管理流程如下：
- 客户端使用用户名密码进行认证
- 服务端生成具有有效时间较短的Access Token，和有效时间较长的Refresh token
- 客户端访问需要鉴权的接口时，携带Access Token
- 如果Access Token没过期，服务端执行操作
- 如果鉴权失败，客户但使用Refresh token 像刷新接口申请新的Access Token
- Refresh token 没有过期，服务端向客户端下发新的Access Token
- 客户端使用新的Access Token访问
- 过期重新登录

## NCX-MALL对Jwt鉴权管理的实践

### Login 获取 Jwt

在完成验证请求接口体合法，账号密码合法后，开始签发jwt。如果多点登录，现在redis缓存中查看是否命中，如果命中直接返回缓存的Jwt，如果没有先向redis缓存后返回