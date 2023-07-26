namespace go demo

// 实现idl管理平台
// 提供增删改查接口

struct AddReq {
    1: required string serviceName(api.body = "service_name")
    2: required i32 version(api.body = "version")
    3: required string content(api.body = "content") 
}

struct AddResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct DeleteReq {
    1: required string serviceName(api.body = "service_name")
    2: required i32 version(api.body = "version")
}

struct DeleteResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct UpdateReq {
    1: required string serviceName(api.body = "service_name")
    2: required i32 version(api.body = "version")
    3: required string content(api.body = "content") 
}

struct UpdateResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct QueryReq {
    1: required string serviceName(api.query = "service_name")
    2: required i32 version(api.query = "version")
}

struct QueryResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
    3: string content(api.body='content'),
}

service IdlManager {
    AddResp Add(1: AddReq addReq)(api.post = "/idl/add")
    DeleteResp Delete(1: DeleteReq deleteReq)(api.delete = "/idl/delete")
    UpdateResp Update(1: UpdateReq updateReq)(api.put = "/idl/update")
    QueryResp Query(1: QueryReq queryReq)(api.get = "/idl/query")

}