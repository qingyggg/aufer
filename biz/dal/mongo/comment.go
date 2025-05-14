package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Comment struct {
	ArticleID  string        `bson:"article_id"`
	Content    string        `bson:"content"`
	Degree     int8          `bson:"degree"` //用以请求顶级评论,degree==1
	UserID     string        `bson:"user_id"`
	HashID     string        `bson:"hash_id"`
	ParentID   string        `bson:"parent_id"`
	CreateTime bson.DateTime `bson:"create_time"`
}

type CommentClosure struct {
	AncestorID   string `bson:"ancestor"`
	DescendantID string `bson:"descendant"`
	ArticleID    string `bson:"article_id"`
	Depth        int    `bson:"depth"`
}

// CommentItem 用以请求comment list
type CommentItem struct {
	ArticleID  string        `bson:"article_id"`
	Content    string        `bson:"content"`
	Degree     int8          `bson:"degree"`
	UserID     string        `bson:"user_id"`
	HashID     string        `bson:"hash_id"`
	CreateTime bson.DateTime `bson:"create_time"`
	ChildNum   int64         `bson:"child_num"`
	ParentUID  string        `bson:"parent_uid"` //@xxx reply_to @xxx
}
