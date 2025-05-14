package mongo

// Article 定义 article 数据结构
type Article struct {
	ArticleID string `bson:"article_id"`
	Content   string `bson:"content"`
}
