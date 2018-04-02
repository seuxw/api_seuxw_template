#API_SEUXW_TEMPLATE

## 基础信息
Author | CreateDate | ProjectName | ChineseName
--- | --- | --- | ---
TauWoo | 2018-04-02 | api_seuxw_template | 小微API模板

## OS
OS | Version
--- | ---
DevOS | Linux Deepin-15.5 (Debian 6.3.0-11)

## 项目开始的提示
- 进入seuxw目录后，执行make可以在seuxw/_output/local/bin目录生成可执行程序
- 执行完make后可以在VSCode中按F5进行程序调试
- 新立项目修改的东西请自行探索
- 需要在GOPATH(/data/code/com/go)中添加部分VSCode插件依赖[下载地址](https://share.weiyun.com/5IzgLKh)

## 目录介绍
目录 | 详情
--- | ---
.vscode | vscode启动和环境配置文件
data | 数据库建表语句
seuxw | 项目代码
seuxw/_output | 在./seuxw 根目录下执行make 将会在此处生成软链接
seuxw/bash | shell
seuxw/filter | API项目路径
seuxw/vendor | 由网络下载的模块
seuxw/x      | 外部模块
seuxw/embrice| 内部模块
seuxw/embrice/api | 调用外部API
seuxw/embrice/constant | 常量
seuxw/embrice/entity | 结构体
seuxw/embrice/extension | 扩展
seuxw/embrice/middleware | 中间件
seuxw/embrice/rdb | 数据库操作