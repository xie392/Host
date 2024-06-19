# 首次运行准备

- docker 环境
- go 环境

## 1.mysql

首次运行需要拉取 mysql 镜像并运行，命令如下：
```shell
docker run --name host-mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=host -e MYSQL_USER=user -e MYSQL_PASSWORD=123456 -p 3306:3306 -d mysql:latest
```
后续有需要可以直接使用 docker start host-mysql 命令启动 mysql 容器。


# 启动

```shell
go run main.go start
```

# 目录介绍

```
├── apps                                # 存放应用代码          
│   ├── host                            # 存放主机应用代码     
│   │   ├── http                        # 存放HTTP应用代码     
│   │   ├── impl                        # 存放实现应用代码   
│   │   │   ├── host.go                 # 主机应用实现代码
│   │   │   ├── host_test.go            # 主机应用测试代码
│   │   │   ├── mysql.go                # MySQL实现代码
│   ├── interface.go                    # 应用接口定义
│   ├── model                           # 存放模型代码        
│   ├── gateway                         # 存放网关应用代码     
│   ├── micro                           # 存放微服务应用代码      
│   ├── oss                             # 存放对象存储应用代码 
├── cmd                                 # 存放命令行工具                             
│   ├── Dockerfile
├── conf                                # 存放配置文件
│   ├── config.go                       # 配置文件入口文件
│   ├── config_test.go                  # 配置文件测试文件, 用于测试配置是否正确
│   ├── load.go                         # 配置文件加载函数, 用于加载配置文件
│   ├── log.go                          # 日志配置文件
├── etc                                 # 存放配置文件模板
│   ├── config.toml                     # 配置文件模板(默认配置文件)
│   ├── config.env                      # 配置文件模板
├── docs                                # 存放文档
├── pkg                                 # 存放公共库代码
│   ├── host
├── protocol                            # 存放协议定义
│   ├── Dockerfile
├── version                             # 存放版本信息
├── go.mod                              # go module文件
├── main.go                             # 程序入口文件
```