package login

import (
	"encoding/json"
	"net/http"

	"github.com/denil/promethee/src/config/helper"
)

var userActive User

func Tokenku(w http.ResponseWriter, req *http.Request) {
	var interfac interface{}
	if !LoginWithToken(w, req) {
		helper.ThrowError("Session has Ended", 503, w)
		return
	} else {
		interfac = map[string]interface{}{
			"message": "Login Success",
			"code":    200,
			"user": map[string]interface{}{
				"id":    userActive.ID,
				"name":  userActive.Name,
				"email": userActive.Email,
			},
		}
		helper.JsonResponse(w, interfac)
	}
}
func DoLogin(w http.ResponseWriter, req *http.Request) {
	var userPost UserLoginPost
	if req.Method == "POST" {
		jsn := helper.GetParam(req)
		marsel, err := json.Marshal(jsn)
		if err != nil {
			helper.JsonResponse(w, map[string]interface{}{
				"message": "ERROR CONVERTING JSON",
				"code":    500,
			})
		}
		json.Unmarshal(marsel, &userPost)
		login := Attempt(userPost.Email, userPost.Password)
		if login {
			helper.JsonResponse(w, map[string]interface{}{
				"message": "berhasil Login",
				"code":    200,
				"user": map[string]interface{}{
					"id":    userActive.ID,
					"name":  userActive.Name,
					"email": userActive.Email,
					"token": userActive.Token,
				},
			})
		} else {
			helper.JsonResponse(w, map[string]interface{}{
				"message": "Login Failed Invalid Username or Password",
				"code":    200,
			})
		}
	} else {
		helper.JsonResponse(w, map[string]interface{}{
			"message": "Not supported Method",
			"code":    401,
		})
	}
}
