package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"net/url"
	"time"
)

// MakeBucket create a bucket with a specified name
func (m *MyMinio) MakeBucket(ctx context.Context, bucketName string) error {
	exists, err := m.mc.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

// PutToBucket put the file into the bucket by *multipart.FileHeader
func (m *MyMinio) PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	info, err = m.mc.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	fileObj.Close()
	return info, err
}

// GetObjURL get the original link of the file in minio
func (m *MyMinio) GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	u, err = m.mc.PresignedGetObject(ctx, bucketName, filename, exp, make(url.Values))
	return u, err
}

// PutToBucketByBuf put the file into the bucket by *bytes.Buffer
func (m *MyMinio) PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = m.mc.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByFilePath put the file into the bucket by filepath
func (m *MyMinio) PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = m.mc.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}

// DelObject delete file in bucket
func (m *MyMinio) DelObject(ctx context.Context, bucketName, fileName string) error {
	// 执行删除操作
	err := m.mc.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	return err
}
