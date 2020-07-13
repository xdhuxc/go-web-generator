### go-web-generator

用于生成一个 go web 工程

生成的项目目录结构如下：
```markdown
.
├── Dockerfile
├── Makefile
├── README.md
├── config.local.yaml
├── deploy
│   └── kubernetes.yaml
├── dist
│   ├── favicon-16x16.png
│   ├── favicon-32x32.png
│   ├── index.html
│   ├── oauth2-redirect.html
│   ├── swagger-ui-bundle.js
│   ├── swagger-ui-bundle.js.map
│   ├── swagger-ui-standalone-preset.js
│   ├── swagger-ui-standalone-preset.js.map
│   ├── swagger-ui.css
│   ├── swagger-ui.css.map
│   ├── swagger-ui.js
│   └── swagger-ui.js.map
├── main.go
├── sql
│   └── xdhuxc_user.ddl.sql
└── src
    ├── apis
    │   ├── base.go
    │   ├── filter.go
    │   ├── hi.go
    │   ├── metrics.go
    │   ├── router.go
    │   ├── swagger.go
    │   └── user.go
    ├── config
    │   └── config.go
    ├── db
    │   └── db.go
    ├── models
    │   ├── hi.go
    │   └── user.go
    ├── pkg
    │   ├── constant.go
    │   └── status_code.go
    ├── service
        ├── base.go
        ├── hi.go
        ├── user.go
        └── user_test.go

```

功能：
* 生成一个完整的，可运行的工程，包含 swagger，sql，deployment，makeFile，.dockerIgnore，.gitIgnore等文件和目录
* 可以单独生成数据模型和 service，controller 文件

少部分内容需要手工修改


### 运行

```markdown
main --dir /x/a --project scmp-cicd --port 8080
```
参数说明：
* dir：项目生成到的目录，默认为当前目录
* project：项目名称，必须指定
* port：Web 项目启动的端口号，默认为 8080


### swagger 的问题


### 注意，

kubernetes.yaml 文件中的 `httpGet.path` 需要修改为当前项目的

### 参考资料

cobra 学习资料：https://github.com/cli/cli
