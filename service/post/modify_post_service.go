package post

import (
	"encoding/json"
	"fmt"
	"goblog/database/mysql"
	gerr "goblog/error"
	"goblog/rep"
	"goblog/serializer"
	"strings"
	"time"
)

type ModifyPostService struct {
	Id          int             `json:"id"`
	Content     string          `json:"content"`
	Title       string          `json:"title"`
	AuthorId    int             `json:"authorId"`
	Description string          `json:"description"`
	Tag         json.RawMessage `json:"tag"`
}

func (srv *ModifyPostService) Modify() *rep.Response {

	var fields []string
	var values []any

	if srv.Content != "" {
		fields = append(fields, "content=?")
		values = append(values, srv.Content)

		if srv.Description == "" {
			srv.Description = serializer.Post.ExtractMarkdownDescription(srv.Content)
		}
	}
	if srv.Description != "" {
		fields = append(fields, "description=?")
		values = append(values, srv.Description)
	}

	if srv.Title != "" {
		fields = append(fields, "title=?")
		values = append(values, srv.Title)
	}

	tagStr := string(srv.Tag)
	if tagStr != "" && tagStr != "null" {
		fields = append(fields, "tag=?")
		values = append(values, tagStr)
	}

	if len(fields) == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "修改无有效字段")
	}

	// 更新时间
	fields = append(fields, "update_time=?")
	values = append(values, time.Now().Format("2006-01-02 15:04:05"))

	values = append(values, srv.AuthorId, srv.Id)

	rows, er := mysql.DB.Exec("UPDATE gb_post SET "+strings.Join(fields, ",")+" where author_id=? and id=?", values...)

	if er != nil {
		fmt.Println("sss", er)
		return rep.Build(nil, gerr.ErrUnExpect, "修改文章失败")
	}
	rowNum, _ := rows.RowsAffected()

	if rowNum == 0 {
		return rep.Build(nil, gerr.ErrUnExpect, "文章不存在或无权限")
	}

	return rep.BuildOkResponse(nil)
}
