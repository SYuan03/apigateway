**idl管理平台**

打算使用idl来生成

### 编写idl文件

### 根据idl生成代码

``` bash
hz new -module github.com/SYuan03/idlmanage -idl idl/manage.thrift
go mod tidy
// 更新可用
hz update -idl idl/manage.thrift
```

### 修改监听端口

修改main.go，端口为6666

### 接口文档（使用说明

#### 增

| url    | http://127.0.0.1:6666/idl/add |
| ------ | ----------------------------- |
| method | POST                          |

body如下

``` json
"body": {
    "mode": "formdata",
    "formdata": [
        {
            "key": "service_name",
            "value": "sad",
            "type": "text"
        },
        {
            "key": "version",
            "value": "23",
            "type": "text"
        },
        {
            "key": "content",
            "value": "namespace go demo\n\n// 实现idl管理平台\n// 提供增删改查接口\n\nstruct AddReq {\n    1: required string serviceName(api.body = \"service_name\")\n    2: required i32 version(api.body = \"version\")\n    3: required string content(api.body = \"content\") \n}\n\nstruct AddResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n}\n\nstruct DeleteReq {\n    1: required string serviceName(api.body = \"service_name\")\n    2: required i32 version(api.body = \"version\")\n}\n\nstruct DeleteResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n}\n\nstruct UpdateReq {\n    1: required string serviceName(api.body = \"service_name\")\n    2: required i32 version(api.body = \"version\")\n    3: required string content(api.body = \"content\") \n}\n\nstruct UpdateResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n}\n\nstruct QueryReq {\n    1: required string serviceName(api.body = \"service_name\")\n    2: required i32 version(api.body = \"version\")\n}\n\nstruct QueryResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n    3: string content(api.body='content'),\n}\n\nservice IdlManager {\n    AddResp Add(1: AddReq addReq)(api.post = \"/idl/add\")\n    DeleteResp Delete(1: DeleteReq deleteReq)(api.delete = \"/idl/delete\")\n    UpdateResp Update(1: UpdateReq updateReq)(api.put = \"/idl/update\")\n    QueryResp Query(1: QueryReq queryReq)(api.get = \"/idl/query\")\n\n}",
            "type": "text"
        }
    ]
},
```



#### 删

| url    | http://127.0.0.1:6666/idl/delete |
| ------ | -------------------------------- |
| method | DELETE                           |

body如下

``` json
"body": {
    "mode": "formdata",
    "formdata": [
        {
            "key": "service_name",
            "value": "sad",
            "type": "text"
        },
        {
            "key": "version",
            "value": "23",
            "type": "text"
        }
    ]
},
```



#### 改

| url    | http://127.0.0.1:6666/idl/update |
| ------ | -------------------------------- |
| method | PUT                              |

body如下

``` json
"body": {
    "mode": "formdata",
    "formdata": [
        {
            "key": "service_name",
            "value": "sad",
            "type": "text"
        },
        {
            "key": "version",
            "value": "23",
            "type": "text"
        },
        {
            "key": "content",
            "value": "sadafasfas",
            "type": "text"
        }
    ]
},
```



#### 查

| url    | http://127.0.0.1:6666/idl/query?service_name=sad&version=23 |
| ------ | ----------------------------------------------------------- |
| method | GET                                                         |



### 数据库内容

| ServiceName     | Version | Content |
| --------------- | ------- | ------- |
| StudentServiceA | 1       | 略      |
| StudentServiceB | 2       | 略      |
|                 |         |         |

``` bash
curl -X GET "http://127.0.0.1:6666/idl/query?service_name=StudentSeviceA&version=1" -w "\n"
```

![image-20230726204013488](https://s2.loli.net/2023/07/26/sP6a4i19ycKZNJb.png)



### 服务器

在服务器端也部署了一个idl-manage的服务

![image-20230726224735538](https://s2.loli.net/2023/07/26/ETk9Jt8b6WgrZ3X.png)
