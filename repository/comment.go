package repository

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	VideoId  int64
	AuthorId int64
	Content  string
}

func CommentList(video_id int64) (commentlist []Comment) {
	var comments []Comments
	if db.Find(&comments, "video_id=?", video_id).Error != nil {
		return commentlist
	}
	for _, comment := range comments {
		var user User
		if db.First(&user, comment.AuthorId).Error != nil {
			panic("Author can't be find")
		}
		commentlist = append(commentlist, Comment{
			Id:         int64(comment.ID),
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}
	return commentlist
}
