package post

// func (post *model.Post) BeforeDelete(tx *gorm.DB) (err error) {
// 	var comments []model.Comment
// 	var follows []model.Follow

// 	tx.Where("postid = ?", post.ID).Find(&comments)
// 	tx.Where("postid = ?", post.ID).Find(&follows)
// 	for _, comment := range comments {
// 		_, err := databasecomment.Delete(comment.UserId, comment.Postid)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	for _, follow := range follows {
// 		_, err := databasefollow.DeleteFollow(&follow)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
