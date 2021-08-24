package funcsRoutes

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

type User struct {
	Id           string
	Username     string
	Email        string
	Password     string
	Picture      string
	Role         string
	Notification string
	History      string
}

type UserResponse struct {
	Err  bool
	Msg  string
	User User
}

func ValideEmail(email string) bool {
	reg := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+`)
	return reg.MatchString(email)
}

func UserSendResponse(result UserResponse, res http.ResponseWriter) {
	jsonRes, err := json.Marshal(result)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonRes)
}

func parseTags(stringParse string) []string {

	var res []string = strings.Split(stringParse, "|")
	return res
}
