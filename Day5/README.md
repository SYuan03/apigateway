返回类型强制转换成HTTPResponse

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}' -w "\n"
```

``` bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/query?id=100 -w "\n"
```
StudentServiceA
``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info-A -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}' -w "\n"
```

``` bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/query-A?id=100 -w "\n"
```

StudentServiceB

``` bash
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/add-student-info-B -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}' -w "\n"
```

``` bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/query-B?id=100 -w "\n"
```



