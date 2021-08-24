package funcsRoutes

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(res http.ResponseWriter, req *http.Request) {
	var error bool
	_ = CreateDb([]string{"create table users (id text unique, username text unique, email text unique, password text, picture text, role text, notification text, history text)"}, userFilename)

	//error method
	if req.Method != "POST" {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		UserSendResponse(response, res)
		return
	}

	//get body of request
	//don't use body of GET request
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		UserSendResponse(response, res)
		return
	}

	if len(user.Username) < 2 || !ValideEmail(user.Email) || len(user.Password) < 8 {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		UserSendResponse(response, res)
		return
	}

	user.Id = GenerateNumber()
	mdp, _ := (bcrypt.GenerateFromPassword([]byte(user.Password), len(user.Password)))
	user.Password = string(mdp)

	error = AddUser(user)

	if error {
		response := UserResponse{Err: true, Msg: "ce nom d'utilisateur ou cette email existe deja."}
		UserSendResponse(response, res)
		return
	}

	response := UserResponse{Err: error, User: user}

	UserSendResponse(response, res)
}

func Login(res http.ResponseWriter, req *http.Request) {
	error := CreateDb([]string{"create table users (id text unique, username text unique, email text unique, password text, picture text, role text, notification text, history text)"}, userFilename)

	if req.Method != "POST" {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue0."}
		UserSendResponse(response, res)
		return
	}

	var logins, user User
	err := json.NewDecoder(req.Body).Decode(&logins)
	if err != nil {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue1."}
		UserSendResponse(response, res)
		return
	}

	if len(logins.Username) < 2 {
		if !ValideEmail(logins.Email) || len(logins.Password) < 8 {
			response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue2."}
			UserSendResponse(response, res)
			return
		}

		user, error = FindUserByEmail(logins.Email)

		if error || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logins.Password)) != nil {
			response := UserResponse{Err: true, Msg: "Email ou mot de passe incorrect"}
			UserSendResponse(response, res)
			return
		}
	} else {
		if len(logins.Username) < 2 || len(logins.Password) < 8 {
			http.Error(res, "bad username || password", http.StatusForbidden)
			return
		}

		user, error = FindUserByUsername(logins.Username)

		if error || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logins.Password)) != nil {
			response := UserResponse{Err: true, Msg: "Nom d'utilisateur ou mot de passe incorrect"}
			UserSendResponse(response, res)
			return
		}
	}

	response := UserResponse{Err: error, User: user}

	//json conversion
	UserSendResponse(response, res)
}

func GetUserInfo(res http.ResponseWriter, req *http.Request) {
	error := CreateDb([]string{"create table users (id text unique, username text unique, email text unique, password text, picture text, role text, notification text, history text)"}, userFilename)

	if req.Method != "POST" {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		UserSendResponse(response, res)
		return
	}

	var login, user User
	err := json.NewDecoder(req.Body).Decode(&login)
	if err != nil {
		response := UserResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		UserSendResponse(response, res)
		return
	}

	user, error = FindUserById(login.Id)

	response := UserResponse{Err: error, User: user}

	UserSendResponse(response, res)
}
