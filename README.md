# apiend-core

## 目录说明
- `app` 业务逻辑层	所有的业务逻辑存放目录。
  -   `controller` 控制器	接收/解析用户输入参数的入口/接口层。
  -   `lib` 逻辑封装	公共的业务逻辑封装层，可供不同的包调用。
  -   `model`	数据模型	数据管理层，仅用于操作管理数据，如数据库操作。
- `boot` 初始化包	用于项目初始化参数设置。
- `config` 配置管理	所有的配置文件存放目录。
- `core` 一些核心工具包
- `library` 第三方包 插件包
- `router` 路由注册	用于路由统一的注册管理。
- `go.mod`	依赖管理	使用Go Module包管理的依赖描述文件。
- `main.go` 入口文件	程序入口文件。

## badger

- https://github.com/squiidz/geoalt
- https://github.com/aspcartman/pcache

## project 

## kpess 项目结构
cmd 目录存放用于编译可运行程序的 main 源码，它又分成了子级目录，主要是考虑一个项目可能有多种可运行程序。

src 目录放主要源码，集中在这个目录主要是为了方便查找和替换。src 目录下除了 app.go，router.go 这种顶层入口，又细分如下：

    `util`，工具函数，不会依赖本项目的任何其它逻辑，只会被其它源码依赖；
    `service`，对外部服务的封装，如对 mongodb、redis、zipkin 等 client 的封装，也不会依赖本项目 util 之外的任何其它逻辑，只会被其它源码依赖；
    `schema`，数据模型，与数据库无关，也不会依赖本项目 util 之外的任何其它逻辑，只会被其它源码依赖；
    `model`，通常依赖 util，service 和 schema，实现对数据库操作的主要逻辑，各个 model 内部无相互依赖；
    `bll`，Business logic layer，通常依赖 util，schema 和 model，通过组合调用 model 实现更复杂的业务逻辑；
    `api`，API 接口，通常依赖 util，schema 和 bll，挂载于 Router 上，直接受理客户端请求、提取和验证数据，调用 bll 层处理数据，然后响应给客户端；
    `ctl`，Controller，类似 api 层，通常依赖 util，schema 和 bll，挂载于 Router 上，为客户端响应 View 页面；
    其它如 auth、logger 等则是一些带状态的被其它组件依赖的全局性组件。
    
与 cmd、src 平级的目录可能还会有：web 前端源码目录；config 配置文件目录；vendor go 依赖包目录；dist 编译后的可执行文件目录；doc 文档目录；k8s k8s 配置文件目录等。