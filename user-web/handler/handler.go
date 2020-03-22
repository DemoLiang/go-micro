package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	user "github.com/DemoLiang/go-micro/user-srv/proto/user"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/prometheus/common/log"
)

var (
	serviceClient user.UserService
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = user.NewUserService("go.micro.lgm.srv.user", client.DefaultClient)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// call the backend service
	rsp, err := serviceClient.QueryUserByName(context.Background(), &user.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	log.Info("-----%v",r.Form)
	log.Info("====%v",rsp)
	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true

		rsp.User.Pwd = ""
		response["data"] = rsp.User
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func LoginV2(ctx *gin.Context){
	// call the backend service
	userName := ctx.PostForm("userName")
	log.Info("username:",userName)
	rsp, err := serviceClient.QueryUserByName(context.Background(), &user.Request{
		UserName: userName,
	})
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	//log.Info("-----%v",r.Form)
	log.Info("====%v",rsp)
	pwd:=ctx.PostForm("pwd")
	if rsp.User.Pwd == pwd {
		response["success"] = true

		rsp.User.Pwd = ""
		response["data"] = rsp.User
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	ctx.Header("Content-Type", "application/json;charset=utf-8")

	// encode and write the response as json
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	ctx.JSON(200,response)
}