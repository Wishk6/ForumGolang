package funcsRoutes

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var idsFile string = "sql/ids.db"
var userFilename string = "sql/users.db"
var topicFile string = "sql/subjects.db"
var CommentFile string = "sql/comments.fb"

func CreateDb(tables []string, filename string) bool {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return true
	}
	defer db.Close()

	if err != nil {
		return true
	}

	for _, val := range tables {
		_, err = db.Exec(val)

		if err != nil {
			return true
		}
	}

	return false
}

func IdExist(_id string) (string, bool) {
	db, _ := sql.Open("sqlite3", idsFile)
	var id string

	row, err := db.Prepare("select id from ids where id = ?")

	if err != nil {
		return id, true
	}

	err = row.QueryRow(_id).Scan(&id)

	if err != nil {
		return id, false
	}

	defer db.Close()
	defer row.Close()

	return id, true
}

//User
func AddUser(user User) bool {

	db, _ := sql.Open("sqlite3", userFilename)
	_, err := db.Exec("insert into users(id, username, email, password, picture, role, notification, history) values(?, ?, ?, ?, ?, ?, ?, ?)", user.Id, user.Username, user.Email, user.Password, "", "", "", "")
	return err != nil
}

func FindUserById(_id string) (User, bool) {
	db, _ := sql.Open("sqlite3", userFilename)
	var user User

	row, err := db.Prepare("select id, username, email, password, picture, role, notification, history from users where id = ?")

	if err != nil {
		return user, true
	}

	err = row.QueryRow(_id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Role, &user.Notification, &user.History)

	if err != nil {
		return user, true
	}

	defer db.Close()
	defer row.Close()

	return user, false
}

func FindUserByUsername(username string) (User, bool) {
	db, _ := sql.Open("sqlite3", userFilename)
	var user User

	row, err := db.Prepare("select id, username, email, password, picture, role, notification, history from users where username = ?")

	if err != nil {
		return user, true
	}

	err = row.QueryRow(username).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Role, &user.Notification, &user.History)

	if err != nil {
		return user, true
	}

	defer db.Close()
	defer row.Close()

	return user, false
}

func FindUserByEmail(email string) (User, bool) {
	db, _ := sql.Open("sqlite3", userFilename)
	var user User

	row, err := db.Prepare("select id, username, email, password, picture, role, notification, history from users where email = ?")

	if err != nil {
		return user, true
	}

	err = row.QueryRow(email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Picture, &user.Role, &user.Notification, &user.History)

	if err != nil {
		return user, true
	}

	defer db.Close()
	defer row.Close()

	return user, false
}

func UpdateUserById(db *sql.DB, _id string, pictureProfile string, pseudo string, email string, password string) bool {

	_, errget := FindUserById(_id)

	if !errget {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare("update users set Picture=?, username=?, email=?, password=? where id=?")
		_, err := stmt.Exec(pictureProfile, pseudo, email, password, _id)

		if err != nil {
			println("error update user")
			return true
		}

		tx.Commit()
		return false
	}
	println("error update user")
	return true

}

// Topic

func AddTopic(topic Topic) bool {

	db, _ := sql.Open("sqlite3", topicFile)
	_, err := db.Exec("insert into topics(idUser, id, name, tags, date, likes, dislikes) values(?, ?, ?, ?, ?, ?, ?)", topic.IdUser, topic.Id, topic.Name, topic.Tags, topic.Date, topic.Likes, topic.Dislikes)

	if err != nil {
		println("error add Topic")
		return true
	}

	return false
}

func FindTopicById(_id string) (Topic, bool) {

	db, _ := sql.Open("sqlite3", topicFile)
	var topic Topic

	row, err := db.Prepare("select idUser, id, name, tags, date, likes, dislikes from topics where id = ?")

	if err != nil {
		println("error find Topic")
		return topic, true
	}

	err = row.QueryRow(_id).Scan(&topic.IdUser, &topic.Id, &topic.Name, &topic.Tags, &topic.Date, &topic.Likes, &topic.Dislikes)

	if err != nil {
		println("error fin Topic during row")
		return topic, true

	}

	defer db.Close()
	defer row.Close()
	return topic, false
}

func FindAllUsersTopic(database *sql.DB, userId string) ([]Topic, bool) {
	var topicArr []Topic
	query := fmt.Sprintf("SELECT idUser, id, name, tags, date, likes, dislikes FROM topics LIMIT 50")
	rows, err := database.Query(query)
	if err != nil {
		return topicArr, true
	}
	defer rows.Close()

	for rows.Next() {
		var new_idUser string
		var new_id string
		var new_name string
		var new_tags string
		var new_date string
		var new_likes string
		var new_dislikes string

		rows.Scan(&new_idUser, &new_id, &new_name, &new_tags, &new_date, &new_likes, &new_dislikes)

		var testesTopic Topic

		testesTopic.IdUser = new_idUser
		testesTopic.Id = new_id
		testesTopic.Name = new_name
		testesTopic.Tags = new_tags
		testesTopic.Date = new_date
		testesTopic.Likes = new_likes
		testesTopic.Dislikes = new_dislikes

		if new_idUser == userId { // test match avec param
			topicArr = append(topicArr, testesTopic)
		}
	}
	return topicArr, false
}

func UpdateTopicLikesById(db *sql.DB, _id string, newlikes string) bool {

	_, errget := FindTopicById(_id)

	if !errget {
		tx, _ := db.Begin()
		stmt, _ := tx.Prepare("update Topics set likes= ? where id = ?")
		_, err := stmt.Exec(newlikes, _id)

		if err != nil {
			println("error update likes")
			return true
		}

		tx.Commit()
		return false
	}
	println("error update likes")
	return true // already liked
}

func UpdateTopicDislikesById(db *sql.DB, _id string, newDislikes string) bool {

	_, errget := FindTopicById(_id)

	if !errget {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare("update Topics set dislikes=? where id=?")
		_, err := stmt.Exec(newDislikes, _id)

		if err != nil {
			println("error update dislikes")
			return true
		}

		tx.Commit()
		return false
	}
	println("error update dislikes")
	return true
}

// Comment

func AddComment(comment Comment) bool {

	db, _ := sql.Open("sqlite3", CommentFile)
	_, err := db.Exec("insert into comments(idUser, idTopic, id, content, date, likes, dislikes) values(?, ?, ?, ?, ?, ?, ?)", comment.IdUser, comment.IdTopic, comment.Id, comment.Content, comment.Date, comment.Likes, comment.Dislikes)

	if err != nil {
		println("error add comment")
		return true
	}

	return false
}

func FindCommentById(_id string) (Comment, bool) {

	db, _ := sql.Open("sqlite3", CommentFile)
	var comment Comment

	row, err := db.Prepare("select * from comments where id = ?")

	if err != nil {
		println("error find comment")
		return comment, true
	}

	err = row.QueryRow(_id).Scan(&comment.IdUser, &comment.IdTopic, &comment.Id, &comment.Content, &comment.Date, &comment.Likes, &comment.Dislikes)

	if err != nil {
		println("error find comment during row")
		return comment, true

	}

	defer db.Close()
	defer row.Close()
	return comment, false
}

func FindAllCommentOfATopic(database *sql.DB, TopicId string) ([]Comment, bool) {
	var commentArr []Comment
	query := fmt.Sprintf("SELECT idUser, idTopic, id, content, date, likes, dislikes FROM comments LIMIT 50")
	rows, err := database.Query(query)
	if err != nil {
		return commentArr, true
	}
	defer rows.Close()

	for rows.Next() {
		var new_idUser string
		var new_idTopic string
		var new_id string
		var new_content string
		var new_date string
		var new_likes string
		var new_dislikes string

		rows.Scan(&new_idUser, &new_idTopic, &new_id, &new_content, &new_date, &new_likes, &new_dislikes)

		var testedComment Comment

		testedComment.IdUser = new_idUser
		testedComment.IdTopic = new_idTopic
		testedComment.Id = new_id
		testedComment.Content = new_content
		testedComment.Date = new_date
		testedComment.Likes = new_likes
		testedComment.Dislikes = new_dislikes

		if new_idTopic == TopicId { // test match avec param
			commentArr = append(commentArr, testedComment)
		}
	}
	return commentArr, false
}

func UpdateCommentLikesById(db *sql.DB, _id string, newlikes string) bool {

	_, errget := FindCommentById(_id)

	if !errget {
		tx, _ := db.Begin()
		stmt, _ := tx.Prepare("update comments set likes= ? where id = ?")
		_, err := stmt.Exec(newlikes, _id)

		if err != nil {
			println("error update comment likes")
			return true
		}

		tx.Commit()
		return false
	}
	println("error update comment likes")
	return true
}

func UpdateCommentDislikesById(db *sql.DB, _id string, newDislikes string) bool {

	_, errget := FindCommentById(_id)

	if !errget {
		tx, _ := db.Begin()

		stmt, _ := tx.Prepare("update comments set dislikes=? where id=?")
		_, err := stmt.Exec(newDislikes, _id)

		if err != nil {
			println("error update comment dislikes")
			return true
		}

		tx.Commit()
		return false
	}
	println("error update comment dislikes ")
	return true
}
