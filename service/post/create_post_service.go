package post

import (
	"encoding/json"
	"fmt"
	"goober/application/mysql"
	"goober/goober"
	"goober/model/post"
	"goober/serializer"
)

type CreatePostService struct {
	Content     string         `json:"content"`
	Title       string         `json:"title"`
	AuthorId    int            `json:"authorId"`
	Description string         `json:"desription"`
	Tags        []int          `json:"tags"`
	Statu       post.PostStatu `json:"statu"`
}
type CreatePostResult struct {
	Id int64 `json:"id"`
}
type Tag struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func (srv *CreatePostService) Run() *goober.ResponseResult {

	tags, _ := json.Marshal(srv.Tags)

	if srv.Description == "" {
		srv.Description = serializer.Post.ExtractMarkdownDescription(srv.Content)
	}
	fmt.Println("createpost srv", srv)
	fmt.Println("createpost data", srv.Title, srv.Content, srv.AuthorId, srv.Description, string(tags), srv.Statu)
	res, er := mysql.DB().Exec(
		"INSERT INTO gb_post (title,content,author_id,description,tags,statu) VALUES(?,?,?,?,?,?)",
		srv.Title, srv.Content, srv.AuthorId, srv.Description, string(tags), srv.Statu,
	)

	if er != nil {
		fmt.Println("create post error:", er.Error())
		return goober.FailedResult(goober.ErrUnExpect, "创建文章失败")
	}
	id, er := res.LastInsertId()
	if er != nil {
		return goober.FailedResult(goober.ErrUnExpect, "创建文章失败")
	}

	if er != nil {
		return goober.WrongResult(er)
	}

	return goober.OkResult(CreatePostResult{Id: id})
}
