package funcsRoutes

import (
	"encoding/json"
	"net/http"
)

type Comment struct {
	IdUser   string
	IdTopic  string
	Id       string
	Content  string
	Date     string
	Likes    string
	Dislikes string
}

type Topic struct {
	IdUser   string
	Id       string
	Name     string
	Tags     string
	Date     string
	Likes    string
	Dislikes string
	Comments []Comment
}

type TopicResponse struct {
	Err    bool
	Msg    string
	Topic  Topic
	Topics []Topic
}

func TopicSendResponse(result TopicResponse, res http.ResponseWriter) {
	jsonRes, err := json.Marshal(result)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonRes)
}
