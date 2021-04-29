package auth

import (
	"log"
	"net/http"

	"github.com/denil/promethee/src/config/db"
	"github.com/denil/promethee/src/config/helper"
	"github.com/denil/promethee/src/login"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r)
		var validate bool = true

		var user login.User
		user.ID = 0
		user.Name = "no"
		user.Email = "no"
		user.Token = "no"
		token := r.Header.Get("authorization")
		log.Println("token", token)
		q := "SELECT id,name,email,token FROM users WHERE token=?"
		mysqldb := db.Koneksi()
		stmt, err := mysqldb.Prepare(q)

		if err != nil {
			validate = false

		}
		result, err := stmt.Query(token)
		if err != nil {
			validate = false
		}
		for result.Next() {
			err = result.Scan(&user.ID, &user.Name, &user.Email, &user.Token)
			if err != nil {
				validate = false
			}
		}
		if user.Name == "no" {
			validate = false
		}
		if r.RequestURI == "/" {
			validate = true
		}
		if r.RequestURI == "/login" {
			validate = true
		}
		if r.RequestURI == "/login/token" {
			validate = true
		}
		if validate {
			next.ServeHTTP(w, r)
		} else {
			helper.JsonResponse(w, map[string]interface{}{
				"message": "NOT AUTHORIZED",
				"code":    403,
				"r":       r.RequestURI,
			})
			return
		}

	})
}
