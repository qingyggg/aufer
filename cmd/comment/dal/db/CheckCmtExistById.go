package db

import (
	"context"
	mydal "github.com/qingyggg/aufer/cmd/comment/dal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CheckCmtExistById(ctx context.Context, cHashId string) (exist bool, err error) {
	count, err := mydal.MyMongo.Cols.Comment.CountDocuments(ctx, bson.M{
		"hash_id": cHashId,
	})
	if err != nil {
		return false, err
	}
	if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
