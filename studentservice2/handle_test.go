package main

import (
	"context"
	"nju/apigw/kitex_gen/demo2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// 初始化一个 StudentServiceImpl 对象
	s := &StudentServiceImpl{}
	s.InitDB()

	// 构造测试用例
	testCases := []struct {
		name    string
		student *demo2.Student
		want    *demo2.RegisterResp
	}{
		{
			name: "test case 1 - new student",
			student: &demo2.Student{
				Id:    1,
				Name:  "Alice",
				Email: []string{"alice@example.com", "alice@gmail.com"},
				College: &demo2.College{
					Name:    "Engineering",
					Address: "123 Main St",
				},
				Gender: "man",
			},
			want: &demo2.RegisterResp{
				Success: true,
				Message: "added success",
			},
		},
		{
			name: "test case 2 - existing student",
			student: &demo2.Student{
				Id:    1,
				Name:  "Alice",
				Email: []string{"alice@example.com", "alice@gmail.com"},
				College: &demo2.College{
					Name:    "Engineering",
					Address: "123 Main St",
				},
				Gender: "man",
			},
			want: &demo2.RegisterResp{
				Success: false,
				Message: "User Already Exist",
			},
		},
		{
			name: "test case 3 - missing student name",
			student: &demo2.Student{
				Id:    2,
				Email: []string{"bob@example.com"},
				College: &demo2.College{
					Name:    "Business",
					Address: "456 Main St",
				},
				Gender: "man",
			},
			want: &demo2.RegisterResp{
				Success: false,
				Message: "missing student name",
			},
		},
		{
			name: "test case 4 - missing college name",
			student: &demo2.Student{
				Id:    3,
				Name:  "Charlie",
				Email: []string{"charlie@example.com"},
				College: &demo2.College{
					Address: "789 Main St",
				},
				Gender: "man",
			},
			want: &demo2.RegisterResp{
				Success: false,
				Message: "missing college information",
			},
		},
		{
			name: "test case 5 - empty email",
			student: &demo2.Student{
				Id:    5,
				Name:  "Emily",
				Email: []string{},
				College: &demo2.College{
					Name:    "Business",
					Address: "456 Main St",
				},
				Gender: "man",
			},
			want: &demo2.RegisterResp{
				Success: false,
				Message: "no email provided",
			},
		},
	}

	// 遍历测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用 Register 函数
			got, err := s.Register(context.Background(), tc.student)

			// 验证函数的返回值是否符合预期
			if tc.want == nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	// 初始化一个 StudentServiceImpl 对象
	s := &StudentServiceImpl{}
	s.InitDB()

	// 向数据库中插入一个学生记录
	student := &demo2.Student{
		Id:    1,
		Name:  "Alice",
		Email: []string{"alice@example.com", "alice@gmail.com"},
		College: &demo2.College{
			Name:    "Engineering",
			Address: "123 Main St",
		},
		Gender: "man",
	}
	s.Db.Table("students").Create(student2Model(student))

	// 构造测试用例
	testCases := []struct {
		name string
		id   int32
		want *demo2.Student
	}{
		{
			name: "test case 1 - existing student",
			id:   1,
			want: student,
		},
		{
			name: "test case 2 - non-existing student",
			id:   2,
			want: nil,
		},
	}

	// 遍历测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用 Query 函数
			got, err := s.Query(context.Background(), &demo2.QueryReq{Id: tc.id})

			// 验证函数的返回值是否符合预期
			if tc.want == nil {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
