# API Gateway部署文档（二，运行）

## 终端一、启动httpsvr

```bash
cd ../httpsvr
go run .
```

观察到信息，httpsvr listening on address = [::]:8888，端口启动即可

## 终端二、启动etcd

```bash 
etcd --log-level debug
```

观察到etcd启动即可

## 终端三、启动studentservice

```bash
cd ../studentsevice
go run .
```

观察到端口启动即可

## 终端四、启动studentservice2

```bash
cd ../studentsevice2
go run .
```

观察到端口启动即可

## 终端五、运行测试

### 1.对studentservice的测试

#### 添加学生信息

```bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceA/Register -d '{"id": 103, "name":"Easma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"mm"}' -w "\n"
```



出现以下响应信息即添加学生信息成功

```bash
{"message":"Apigw Success!","resp":"{success:true,message:added success}"}
```



#### 利用id查询刚刚添加的学生信息

```bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceA/Query -d '{"id":103}' -w "\n"
```

出现以下响应信息即查询成功

```bash
{"message":"Apigw Success!","resp":"{id:103,name:Easma,college:{name:software college,address:逸夫},email:[emma@nju.com]}"}
```



### 2.对studentservice2的测试

#### 添加学生信息

```bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceB/Register -d '{"id": 10, "name":"Easma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"],"gender":"man"}' -w "\n"
```

出现以下响应信息即添加学生信息成功

```bash
{"message":"Apigw Success!","resp":"{success:true,message:added success}"}
```



#### 利用id查询刚刚添加的学生信息

```bash
 curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/apigw/studentserviceB/Query -d '{"id":10}' -w "\n"
```

出现以下响应即查询成功

```bash
{"message":"Apigw Success!","resp":"{id:10,name:Easma,college:{name:software college,address:逸夫},email:[emma@nju.com],gender:man}"}
```

