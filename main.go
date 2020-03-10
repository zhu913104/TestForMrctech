package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//會員資料struct
type Member struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

//回傳json struct
type Return_json struct {
	Code    int8   `json:"Code"`
	Message string `json:"Message"`
	Result  *IsOK  `json:"Result"`
}

type IsOK struct {
	IsOK bool `json:"IsOK"`
}

var members []Member

// 取得所有會員資料(測試用)
func getalldata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func memberInMembers(a Member, list []Member) bool {
	for _, item := range list {
		if item.Account == a.Account {
			return true
		}
	}
	return false
}

//新增會員
func creat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	if !memberInMembers(member, members) {
		members = append(members, member)
		reJSON := Return_json{Code: 0, Message: "", Result: &IsOK{IsOK: true}}
		json.NewEncoder(w).Encode(reJSON)
		return
	}

}

//刪除會員
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	for index, item := range members {
		if item.Account == member.Account {
			members = append(members[:index], members[index+1:]...)
			break
		}
	}
	reJSON := Return_json{Code: 0, Message: "", Result: &IsOK{IsOK: true}}
	json.NewEncoder(w).Encode(reJSON)
}

//修改會員密碼
func change(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	for index, item := range members {
		if item.Account == member.Account {
			members = append(members[:index], members[index+1:]...)
			members = append(members, member)
			reJSON := Return_json{Code: 0, Message: "", Result: &IsOK{IsOK: true}}
			json.NewEncoder(w).Encode(reJSON)
			return
		}
	}
}

//驗證帳號密碼
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	for _, item := range members {
		if item.Account == member.Account {
			if item.Password == member.Password {
				reJSON := Return_json{Code: 0, Message: "", Result: nil}
				json.NewEncoder(w).Encode(reJSON)
				return
			}
		}
	}
	reJSON := Return_json{Code: 2, Message: "Login Failed", Result: nil}
	json.NewEncoder(w).Encode(reJSON)
}

func main() {
	r := mux.NewRouter()

	members = append(members, Member{Account: "test123", Password: "123"})
	members = append(members, Member{Account: "test456", Password: "456"})

	r.HandleFunc("/v1/getalldata", getalldata).Methods("GET")
	r.HandleFunc("/v1/user/create", creat).Methods("POST")
	r.HandleFunc("/v1/user/delete", delete).Methods("POST")
	r.HandleFunc("/v1/user/pwd/change", change).Methods("POST")
	r.HandleFunc("/v1/user/login", login).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
