package main

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	demo "nju/apigw/kitex_gen/demo"
	"strings"
	"time"
)

type Student struct {
	Id             int32
	Name           string
	Email          string
	CollegeName    string
	CollegeAddress string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	Db *gorm.DB
}

func (s *StudentServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(Student{})

	err = db.Migrator().CreateTable(Student{})
	if err != nil {
		panic(err)
	}
	s.Db = db
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// *TODO: Your code here...
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
	//db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	var stuRes *Student
	fmt.Print(req.Id)
	result := s.Db.Table("students").First(&stuRes, req.Id)
	if result.Error == nil {
		return model2Student(stuRes), nil
	}
	return nil, result.Error
}
