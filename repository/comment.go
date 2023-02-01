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

// CommentActionAdd Add comment to DB
func CommentActionAdd(content string, userid int64, videoid int64) (id int64) {
	comments := Comments{
		VideoId:  videoid,
		AuthorId: userid,
		Content:  content,
	}
	// insert new comments
	if db.Create(&comments).Error != nil {
		panic("failed to insert")
	}
	// video's comment count +1
	if db.Model(&Videolist{}).Where("id = ?", videoid).Update("comment_count", gorm.Expr("comment_count+?", 1)).Error != nil {
		panic("failed to update table video_list")
	}
	// get comment id
	if db.Last(&comments, "author_id=?", videoid).Error != nil {
		panic("Author can't be find")
	}
	return int64(comments.ID)
}
