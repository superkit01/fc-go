# fc-go

#### 介绍
阿里云函数计算基于Golang的CustomRuntime的Demo

#### 项目说明
实践理解：函数计算FC部署后执行的是Runtime中的bootstrap，所以bootstrap可以是一个二进制可执行程序也可以是一个启动脚本

目前函数计算有两种类型分别为事件函数和Http函数，Http函数默认监听的端口为9000

阿里云官网目前尚未提供Golang的Runtime的Http函数的Demo(应该很快会提供，因为没什么难点)，本项目是Golang的Http函数的Demo；

#### 流程

1. build-image 中的Dockerfile提供Golang的编译环境，也可以本地编译（注意：本地如果不是x86_64的cpu架构，要使用交差编译）
2. 编译后的文件名称必须是bootstrap，并且打包成code.zip
3. 可以通过函数计算后台上传，也可以通过fun命令行部署template.yaml