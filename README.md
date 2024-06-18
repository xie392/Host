# 目录介绍

```
├── apps                                # 存放应用代码          
│   ├── host                            # 存放主机应用代码     
│   │   ├──   http                      # 存放HTTP应用代码     
│   │   ├──   impl                      # 存放实现应用代码   
│   ├── host.go                         # 主机应用入口文件    
│   ├── gateway                         # 存放网关应用代码     
│   ├── micro                           # 存放微服务应用代码      
│   ├── oss                             # 存放对象存储应用代码 
├── cmd                                 # 存放命令行工具                             
│   ├── Dockerfile
├── conf                                # 存放配置文件
│   ├── Dockerfile
├── pkg                                 # 存放公共库代码
│   ├── host
├── protocol                            # 存放协议定义
│   ├── Dockerfile
├── version                             # 存放版本信息
├── go.mod                              # go module文件
├── main.go                             # 程序入口文件
```