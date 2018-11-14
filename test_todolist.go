package main

import (
	"github.com/labstack/echo"
	"server/todolistsvr"
)

type todolistService struct {
	items []*todolistsvr.Todo
}

//a function to check if a string exist in an array of string
func stringInSlice(a string, arr []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (s *todolistService) Add(c echo.Context, req *todolistsvr.AddReq) (resp *todolistsvr.AddResp, bizError *todolistsvr.AddError, err error) {
	if len(s.items) > 4 {
		bizError = new(todolistsvr.AddError)
		bizError.Req = req
		bizError.Error = "Can't add more than 5 todos"
		return
	}

	//unique check
	for _, b := range s.items {
        if b.Title == req.Item.Title {
            bizError = new(todolistsvr.AddError)
			bizError.Req = req
			bizError.Error = req.Item.Title+" already exist"
			return
        }
	}

	if req.Item == nil {
		e := new(todolistsvr.CommonError)
		err = e
		return
	}

	s.items = append(s.items, req.Item)
	resp = new(todolistsvr.AddResp)
	resp.Count = len(s.items)
	return
}

func (s *todolistService) List(c echo.Context, req *todolistsvr.Empty) (resp *todolistsvr.ListResp, err error) {
	resp = new(todolistsvr.ListResp)
	resp.Items = s.items
	return
}

func main() {
	e := echo.New()

	srv := &todolistService{}
	todolistsvr.RegisterTodolistService(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}
