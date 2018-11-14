package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/apisvr"
	"strings"
)

func testHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	fmt.Println("Message:  ", message)
	message = "Hello " + message
	w.Write([]byte(message))
}

func testAPI(w http.ResponseWriter, r *http.Request) {
	tgarr := []*apisvr.Tag{
		&apisvr.Tag{
			Tag_id:   0,
			Tag_name: "TagName0",
		},
		&apisvr.Tag{
			Tag_id:   1,
			Tag_name: "TagName1",
		},
	}

	serv := apisvr.Service{1, "aServiceName", "aProductid", "aProductName", "aDesc", tgarr}

	j, err := json.Marshal(serv)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(j))
	w.Write(j)
}

func main() {
	//http.HandleFunc("/", testHello)
	http.HandleFunc("/test", testAPI)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
