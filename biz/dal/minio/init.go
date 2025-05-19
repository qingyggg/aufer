package minio

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/qingyggg/aufer/pkg/constants"
	"github.com/qingyggg/aufer/pkg/errno"
)

type MyMinio struct {
	mc *minio.Client
}

// Init 初始化 MinIO 客户端
func Init() (error, *MyMinio) {
	var err error
	myMinio := new(MyMinio)
	mc, err := minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyID, constants.MinioSecretAccessKey, ""),
		Secure: constants.MinioSSL,
	})
	if err != nil {
		return errno.ServiceErr.WithMessage("连接 MinIO 失败: " + err.Error()), nil

	}
	myMinio.mc = mc
	return nil, myMinio
}
func InitForHertz() *MyMinio {
	err, MyMio := Init() //初始存储服务
	if err != nil {
		hlog.Fatal(err)
	}
	MyMio.InitImgBucket()
	return MyMio
}

func (m *MyMinio) InitImgBucket() {
	var err error
	var ctx = context.Background()
	if err = m.MakeBucket(ctx, constants.MinioImgBucketName); err != nil {
		hlog.Fatal("创建 Bucket 失败: ", err)
	}
	hlog.Info("成功创建 Bucket")

	// 上传文件
	imgNames := []string{"mols.jpg", "marisa.jpg"}
	imgPaths := []string{"static/mols.jpg", "static/marisa.jpg"}
	for idx, imgPath := range imgPaths {
		if _, err = m.PutToBucketByFilePath(ctx, constants.MinioImgBucketName, imgNames[idx], imgPaths[idx]); err != nil {
			hlog.Fatal("上传文件失败: ", err)
		}
		hlog.Infof("成功上传文件: %s", imgPath)
	}
}
