package funcsRoutes

import (
	"encoding/json"
	"net/http"
)

type TopicRequest struct {
	Title           string
	Text            string
	Tags            string
	AuthorId        string
	CommentAuthorId string
}

func AddSubject(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		response := TopicResponse{Err: true, Msg: "Une erreur inattendu est survenue."}
		TopicSendResponse(response, res)
		return
	}
	var request TopicRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		response := TopicResponse{Err: true, Msg: "Invalide data."}
		TopicSendResponse(response, res)
		return
	}

	if len(request.Title) < 3 || len(request.AuthorId) < 20 || len(request.Text) < 20 {
		response := TopicResponse{Err: true, Msg: "Invalide data."}
		TopicSendResponse(response, res)
		return
	}

	var topicId string = GenerateNumber()

	comment := Comment{
		IdUser:  request.CommentAuthorId,
		IdTopic: topicId,
		Content: request.Text,
	}

	topic := Topic{
		Id:     topicId,
		IdUser: request.AuthorId,
		Name:   request.Title,
		Tags:   request.Tags,
	}

	AddTopic(topic)
	AddComment(comment)
}
