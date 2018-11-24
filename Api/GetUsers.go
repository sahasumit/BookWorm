package Api

import (
	"log"
	"encoding/json"
	"net/http"
   "strconv"
	"github.com/sahasumit/BookWorm/model"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	//params:=req.Params.Query.Get("Page")
  //if ok
	log.Println("Hello GetUserList " )
 pageNumber,_:=strconv.Atoi(req.URL.Query().Get("Page"))
 pagesize:=10
 pagestart:=pagesize * pageNumber
 var data model.UData
 data.Users = model.GetUserList(pagestart,pagesize)
json.NewEncoder(res).Encode(&data.Users)
}
