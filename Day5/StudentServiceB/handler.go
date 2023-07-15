package main

import (
	"context"
	"errors"
	demo "github.com/SYuan03/Day3/KitexServer/kitex_gen/demo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"time"
	"fmt"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	Db *gorm.DB
}

var stu_data = make(map[int32]demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// *TODO: Your code here...
	// db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	var stuRes *Student
	result := s.Db.Table("students").First(&stuRes, student.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = s.Db.Table("students").Create(student2Model(student))
		if result.Error != nil {
			return nil, errors.New("Insert Error")
		}
		return &demo.RegisterResp{
			Success: true,
			Message: "added success",
		}, nil
	}
	return &demo.RegisterResp{
		Success: false,
		Message: "User Already Exist",
	}, nil
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	var stuRes *Student
	fmt.Println(req.Id)
	result := s.Db.Table("students").First(&stuRes, req.Id)
	if result.Error == nil {
		fmt.Println("Student Found")
		return model2Student(stuRes), nil
	}
	return nil, result.Error
}

type Student struct {
	Id             int32
	Name           string
	Email          string
	CollegeName    string
	CollegeAddress string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (s *StudentServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("StudentServiceA-DB.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// drop table，删掉上一次的表
	// db.Migrator().DropTable(Student{})

	// 检查表是否存在
	if !db.Migrator().HasTable(&Student{}) {
		// 如果表不存在，则创建表
		err = db.Migrator().CreateTable(&Student{})
		if err != nil {
			panic(err)
		}
	}

	s.Db = db
}

func student2Model(student *demo.Student) *Student {
	return &Student{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func model2Student(student *Student) *demo.Student {
	return &demo.Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &demo.College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}
