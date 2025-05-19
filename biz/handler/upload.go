package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/qingyggg/aufer/biz/dal"
	"github.com/qingyggg/aufer/biz/model/http/common"

	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
)

// FileUpload 上传文件接口
//
// @Summary 上传文件
// @Description 用户可以通过该接口上传文件
// @Tags 文件上传接口
// @Accept mpfd
// @Produce json
// @Param file formData file true "上传的文件"
// @Param file_type formData string true  "上传的类型，是头像还是背景图片"
// @Success 200 {object} common.UploadResponse "成功返回文件信息"
// @Failure 400 {object} common.BaseResponse "请求错误"
// @Router /upload/file [post]
func FileUpload(ctx context.Context, c *app.RequestContext) {
	// single file
	file, err := c.FormFile("file")
	//fileType is equal to bucket name
	fileType := c.FormValue("file_type") //avatar,background image
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	// Upload the file to specific dst
	_, err = dal.MyDal.Mio.PutToBucket(ctx, string(fileType), file)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	//url, err := minio.GetObjURL(ctx, string(fileType), file.Filename)
	fullUri := utils.URLconvert(string(fileType) + "/" + file.Filename)
	c.JSON(consts.StatusOK, common.UploadResponse{
		Base:    utils.BaseResp(errno.Success),
		FileUrl: fullUri,
	})
}

// MutiUpload experiment
func MutiUpload(ctx context.Context, c *app.RequestContext) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fmt.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, fmt.Sprintf("./file/upload/%s", file.Filename))
	}
	c.String(consts.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
