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

 