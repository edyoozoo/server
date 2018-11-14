package main

import (
	"fmt"
	"server/echosvr"

	"github.com/labstack/echo"
)

type echoService struct{}

// Echo just return req
func (s *echoService) Echo(c echo.Context, req *echosvr.Msg) (resp *echosvr.Msg, err error) {
	resp = req

	fmt.Println(req.GetMsg())

	return req, err
}

func main() {
	e := echo.New()
	srv := &echoService{}
	echosvr.RegisterEchoService(e, srv)

	//srv.Echo(e.GetContext(), &echosvr.Msg{msg:"1"})

	//start the server
	e.Logger.Fatal(e.Start(":8080"))
}
