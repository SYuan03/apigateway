# API Gateway部署文档（一，运行环境的部署）

系统环境：Ubuntu22.04.2

## 1.golang的安装

### 1.下载并解压go1.20.5的tar压缩包

```bash
wget https://golang.google.cn/dl/go1.20.5linux-amd64.tar.gz
sudo tar -xvzf go1.20.5linux-amd64.tar.gz -C /usr/local/
```

### 2.设置系统变量

#### 1.打开bashrc

```bash
vim ~/.bashrc
```

#### 2.配置系统变量，将如下内容复制到bashrc

```bash
export GOPATH=/home/"username"/gocode
 
export GOROOT=/usr/local/go
 
export PATH=$PATH:$GOROOT/bin
```

#### 3.使变更生效

```bash
source ~/.bashrc
```

### 3.检验golang是否配置成功

```bash
go version
```

若显示正确的版本号，则golang配置成功

## 2.运行所需工具的安装

### 1.hz工具的安装

执行以下命令

```bash
go install github.com/cloudwego/hertz/cmd/hz@latest
```

检验hz工具是否配置完成

执行如下命令

```bash
hz
```

若不出现hz command not found且提示命令格式，则配置完成

### 2.thriftgo安装

执行如下命令

```bash
GO111MODULE=on go install github.com/cloudwego/thriftgo@latest
```

将GOBIN添加到GOPATH

```bash
export PATH=$GOPATH/bin:$PATH
```

检验是否安装成功

```bash
thriftgo --version
```

提示版本信息

### 3.kitex安装

执行如下命令

```bash
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

检验是否安装成功

```bash
kitex --version
```

提示版本信息

### 4.etcd的配置

在以下网址选择etcd压缩包进行下载

https://github.com/etcd-io/etcd/releases/tag/v3.4.27

下载后将压缩包解压，将etcd和etcdctl移动到/usr/local/bin目录下

检验是否安装成功（终端1，2均不在无权限更改的目录下进行，建议在~下执行）

在终端1运行服务端

```bash
ectd --log-level debug
```

启动后不要关闭终端1，另起一个终端2

在终端2写入数据

```bash
etcdctl put greeting "Hello, etcd"
```

在终端2读取数据

```bash
etcdctl get greeting
```

## 3.congratulations

你已经完成了运行环境的配置，快去试试运行吧！

请移步《部署文档2-网关的试运行》
