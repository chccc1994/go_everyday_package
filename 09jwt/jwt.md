## JWT
[Go进阶24:Go JWT RESTful身份认证教程](https://mojotv.cn/go/golang-jwt-auth)
### 1 什么是JWT

`JWT（JSON Web Token）`是一个非常轻巧的规范,这个规范允许我们使用JWT在用户和服务器之间传递安全可靠的信息, 一个JWT由三部分组成,   `Header头部`,`Claims载荷`,`Signature签名`,JWT原理类似我们加盖公章或手写签名的的过程,合同上写了很多条款, 不是随便一张纸随便写啥都可以的,必须要一些证明,比如签名, 比如盖章,`JWT就是通过附加签名`,保证传输过来的信息是真的,而不是伪造的,它将用户信息加密到`token`里,服务器不保存任何用户信息,服务器通过使用保存的密钥验证token的正确性,只要正确即通过验证。
### 2 JWT组成
一个JWT由三部分组成,`Header头部`,`Claims载荷`,`Signature签名`,

- Header头部：头部,表明类型和加密算法， 告诉我们使用的算法和 token 类型
-  Payload: 声明,即载荷（承载的内容）必须使用 sub key 来指定用户 ID, 还可以包括其他信息比如 email, username 等.
- Signature签名：签名,这一部分是将header和claims进行base64转码后,并用header中声明的加密算法加盐(secret)后构成,用来保证 JWT 的真实性. 可以使用不同算法

```go
let tmpstr = base64(header)+base64(claims)
let signature = encrypt(tmpstr,secret)
//最后三者用"."连接,即：
let token = base64(header)+"."+base64(claims)+"."+signature
```
当然您要可以使用在线工具来解析jwt token的payload荷载 [JWT在线解析工具](https://jwt.io/)