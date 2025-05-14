package mongo

import (
	"context"
	"github.com/qingyggg/aufer/pkg/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
)

type Collections struct {
	Notification *mongo.Collection
	Article      *mongo.Collection
	Comment      *mongo.Collection
	CmtClosure   *mongo.Collection
}
type MyMongo struct {
	Db   *mongo.Database
	Cols *Collections
}

func Init() (error, *MyMongo) {
	var err error
	mymongo := new(MyMongo)
	mymongo.Cols = &Collections{}

	// 连接MongoDB
	clientOpts := options.Client().ApplyURI(
		constants.MongoDefaultDSN)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return err, nil
	}

	// 保存全局变量
	Client = client

	// 验证连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err, nil
	}

	// 初始化数据库和表的引用
	DB = client.Database("aufer")
	mymongo.Db = DB

	// 获取集合引用
	mymongo.Cols.Notification = DB.Collection("notification")
	mymongo.Cols.Article = DB.Collection("article")
	mymongo.Cols.Comment = DB.Collection("comment") // 存储文章评论的表
	mymongo.Cols.CmtClosure = DB.Collection("comment_closure")

	// Note: 索引创建已移至Docker容器初始化脚本
	// 详见 pkg/configs/mongo/init-mongo.js

	return nil, mymongo
}

func CreateIndex(collection *mongo.Collection, field bson.M, needUnique bool) error {
	// 创建一个索引模型
	indexModel := mongo.IndexModel{
		Keys:    field,                                 // 升序索引
		Options: options.Index().SetUnique(needUnique), // 可选项，设置唯一索引
	}

	// 使用 Indexes().CreateOne 创建索引
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
