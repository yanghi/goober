package post

import (
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
)

type CreatePostService struct {
	Content     string `json:"content"`
	Title       string `json:"title"`
	AuthorId    int    `json:"authorId"`
	Description string `json:"desription"`
	Tags        []int  `json:"tags"`
}
type CreatePostResult struct {
	Id int64 `json:"id"`
}
type Tag struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func (srv *CreatePostService) Run() *rep.Response {

	// tagQs := tagQuerys(srv.Tag)

	stm, _ := mysql.DB.Prepare("INSERT INTO gb_post (title,content,author_id,description) VALUES(?,?,?,?)")

	res, er := stm.Exec(srv.Title, srv.Content, srv.AuthorId, srv.Description)
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "创建文章失败")
	}
	id, er := res.LastInsertId()
	if er != nil {
		return rep.Build(nil, gerr.ErrUnExpect, "创建文章失败")
	}

	if er != nil {
		return rep.BuildFatalResponse(er)
	}

	return rep.BuildOkResponse(CreatePostResult{Id: id})
}
