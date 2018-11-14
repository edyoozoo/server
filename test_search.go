package main

import (
	"server/searchsvr"

	"github.com/labstack/echo"
)

type searchService struct{}

// search return search response with id replaced
func (s *searchService) Serach(c echo.Context, req *searchsvr.SearchReq) (resp *searchsvr.SearchResponse, bizError *searchsvr.SearchError, err error) {
	resp = new(searchsvr.SearchResponse)
	if req.UserId > 100 {
		bizError = new(searchsvr.SearchError)
		bizError.Req = req
		bizError.Error = "UserId Exceed 100"
		return
	}
	resp.UserId = req.UserId
	resp.Id = 20
	resp.Title = "FakeTitle"
	resp.Completed = true

	return
}

func main() {
	e := echo.New()

	srv := &searchService{}
	searchsvr.RegisterSearchService(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}
