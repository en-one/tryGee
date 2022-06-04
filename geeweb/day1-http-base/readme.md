## 今日内容
- 熟悉net包，构建基础http
- 搭建gee雏形



## Gee 框架的雏形

```api
1    gee/
2      |--gee.go
3      |--go.mod
4    main.go
5    go.mod
```

根目录下go.mod的使用
```api
module example

go 1.13

require gee v0.0.0

replace gee => ./gee
```