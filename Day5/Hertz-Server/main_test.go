package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	queryURLFmt = "http://127.0.0.1:8888/query?id="
	registerURL = "http://127.0.0.1:8888/add-student-info"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func TestStudentService(t *testing.T) {
	for i := 1; i <= 100; i++ {
		newStu := genStudent(i)
		resp, err := rregister(newStu)

      fmt.Println("err: ", err.Error())
      fmt.Println("resp: ", resp)


		Assert(t, err == nil, err)
		Assert(t, resp.Success)

		stu, err := query(i)
		Assert(t, err == nil, err)
		Assert(t, stu.ID == newStu.ID)
		Assert(t, stu.Name == newStu.Name)
		Assert(t, stu.Email[0] == newStu.Email[0])
		Assert(t, stu.College.Name == newStu.College.Name)
	}
}

func BenchmarkStudentService(b *testing.B) {
	for i := 1; i < b.N; i++ {
		newStu := genStudent(i)
		resp, err := rregister(newStu)
		Assert(b, err == nil, err)
		Assert(b, resp.Success, resp.Message)

		stu, err := query(i)
		Assert(b, err == nil, err)
		Assert(b, stu.ID == newStu.ID)
		Assert(b, stu.Name == newStu.Name, newStu.ID, stu.Name, newStu.Name)
		Assert(b, stu.Email[0] == newStu.Email[0])
		Assert(b, stu.College.Name == newStu.College.Name)
	}
}

func rregister(stu *demo.Student) (rResp *demo.RegisterResp, err error) {
	reqBody, err := json.Marshal(stu)
   // fmt.Println("reqBody: ", reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	var resp *http.Response
	req, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewBuffer(reqBody))
	resp, err = httpCli.Do(req)

   // if err == nil {
   //    fmt.Println("afsssss: ")
   // }
   // fmt.Println("resp in rregister: ", resp)

	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &rResp); err != nil {
		return
	}
	return
}

func query(id int) (student demo.Student, err error) {
	var resp *http.Response
	resp, err = httpCli.Get(fmt.Sprint(queryURLFmt, id))
	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &student); err != nil {
		return
	}
	return
}

func genStudent(id int) *demo.Student {
	return &demo.Student{
		ID:   int32(id),
		Name: fmt.Sprintf("student-%d", id),
		College: &demo.College{
			Name:    "",
			Address: "",
		},
		Email: []string{fmt.Sprintf("student-%d@nju.com", id)},
	}
}

// Assert asserts cond is true, otherwise fails the test.
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}
