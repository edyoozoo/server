package main

import (
  	"net/http"
	  //"strings"
	"server/apisvr"
	"github.com/labstack/echo"
)

type aResponse struct {
	response_id int
	response_return string
}

func GetEnv(c echo.Context, req *apisvr.EnvListRequest) (resp *apisvr.EnvListResponse, bizError *apisvr.Error, err error) {
	return 
}

/*func main() {
	e := echo.New()
	
	//if no get parameters output hello world
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	req := &apisvr.EnvListRequest{}
	GetEnv(e.Context, *req)
}*/

