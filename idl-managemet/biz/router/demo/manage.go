// Code generated by hertz generator. DO NOT EDIT.

package demo

import (
	demo "github.com/SYuan03/idlmanage/biz/handler/demo"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_idl := root.Group("/idl", _idlMw()...)
		_idl.POST("/add", append(_addMw(), demo.Add)...)
		_idl.DELETE("/delete", append(_deleteMw(), demo.Delete)...)
		_idl.GET("/query", append(_queryMw(), demo.Query)...)
		_idl.PUT("/update", append(_updateMw(), demo.Update)...)
	}
}