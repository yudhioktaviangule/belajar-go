package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ThrowError(message string, code int, w http.ResponseWriter) {
	JsonResponse(w, map[string]interface{}{`message`: message, "code": code})
}

func JsonResponse(w http.ResponseWriter, datajson interface{}) {
	allowedHeaders := "Access-Control-Allow-Origin, Access-Control-Allow-Methods, Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datajson)
}

func GetParam(request *http.Request) interface{} {
	var Request interface{}
	vd, err := ioutil.ReadAll(request.Body)
	//log.Println(vd)
	if err != nil {
		log.Fatal("ERROR")
	}
	json.Unmarshal(vd, &Request)

	return Request
}
func SetHeader(w http.ResponseWriter) {

}
func MyEncryptor(plaintext string) string {
	password := []byte(plaintext)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("ERROR IN HASHING")
	}
	//log.Fatal(hashedPassword)
	return string(hashedPassword)
}

func MyDecryptor(chipertext string, password string) bool {
	cipher := []byte(chipertext)
	compared := []byte(password)
	err := bcrypt.CompareHashAndPassword(cipher, compared)
	if err == nil {
		return true
	} else {
		return false
	}
}
