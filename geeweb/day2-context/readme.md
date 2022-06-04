## 今日内容
- 将路由(router)独立出来，方便之后增强。
- 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
  

### 1、设计Context<br>
1) 对于Web服务，主要是根据请求*request，构造响应http.ResponseWriter,因此context要包含这两个类型
2) 对于框架来说，context还要支持额外的功能，比如动态路由、中间件


> 
```api
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}
```


### 2、设计router
1) 提前路由相关内容放到router.go,方便下次扩展动态路由
2) router的handler方法传入context