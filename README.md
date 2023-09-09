# CloudWeGo大作业

> 队名：家人们谁懂啊
>
> 队员：dsy chy fzy

## Apigw接口文档

### 简要描述

该接口用于从 URL 中提取服务名称、方法名称和 IDL 版本，并调用对应的服务方法，以及使用相应版本的idl进行更新

### URL地址

apigw/ServiceName/MethodName/IdlVersion

### 请求方法

POST

### 请求参数

相应请求方法的Body

### 响应参数

参数名	类型	描述
message	string	接口执行结果描述
resp	any	接口执行结果

### 响应状态码

| 状态码 | 描述           |
| ------ | -------------- |
| 200    | 请求成功       |
| 400    | 请求参数错误   |
| 500    | 服务器内部错误 |

### 响应示例

#### 请求StudentServieA的Register方法-成功

请求

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceA/Register -d '{"id": 103, "name":"Easma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"mm"}' -w "\n"
```

响应

``` bash
{
	"message": "Apigw Success!",
	"resp": "{success:true,message:added success}"
}
```



#### 请求StudentServieA的Register方法-已有插入

请求

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceA/Register -d '{"id": 103, "name":"Easma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"mm"}' -w "\n"
```

响应

``` bash
{
	"message": "Apigw Success!",
	"resp": "{success:false,message:User Already Exist}"
}
```



#### 请求错误：Too few parameters!

请求

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/1 -d '{"id":103}' -w "\n"
```

响应

```  bash
{
	"message": "Too few parameters!"
}
```



#### 请求错误：Too many parameters!

请求

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/ds/ds/studentserviceA/Query -d '{"id":103}' -
w "\n"
```

响应

```  bash
{
	"message": "Too many parameters!"
}
```



#### 请求StudentServieA的Query方法-成功

请求

``` bash
 curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceA/Query -d '{"id":103}' -w "\n"
```

响应

```  bash
{
	"message": "Apigw Success!",
	"resp": "{email:[emma@nju.com],id:103,name:Easma,college:{name:software college,address:逸夫}}"
}
```



#### 请求StudentServieB的Query方法-成功(多一个gender字段)

请求

``` bash
 curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceB/Query -d '{"id":10}' -w "\n"
```

响应

```  bash
{
	"message": "Apigw Success!",
	"resp": "{id:10,name:Easma,college:{name:software college,address:逸夫},email:[emma@nju.com],gender:man}"
}
```



#### 请求StudentServieB的Query方法-record not found

请求

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceB/Query -d '{"id":999}' -w "\n"
```

响应

```  bash
{
	"message": "remote or network error[remote]: biz error: record not found"
}
```
