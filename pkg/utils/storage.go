package utils

import (
	"context"
	"fmt"
	"github.com/qingyggg/aufer/biz/dal"
	"github.com/qingyggg/aufer/pkg/constants"
	"log"
	"net/url"
	"strings"
)

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(path string) (fullURL string) {
	ctx := context.Background()
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := dal.MyDal.Mio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	u.Scheme = "https"
	u.Host = constants.Host
	u.Path = "/src" + u.Path
	return u.String()
}

// UrlConvertReverse 从完整URL还原数据库中存储的相对路径
func UrlConvertReverse(fullURL string) (path string) {
	if len(fullURL) == 0 {
		return ""
	}

	// 解析传入的 URL
	u, err := url.Parse(fullURL)
	if err != nil {
		log.Printf("解析URL失败: %s \n", err.Error())
		return ""
	}

	// 假设路径的前缀是 "/src"，需要去掉前缀部分
	urlPath := u.Path
	if strings.HasPrefix(urlPath, "/src/") {
		urlPath = strings.TrimPrefix(urlPath, "/src/")
	} else {
		log.Printf("URL路径无效: %s \n", fullURL)
		return ""
	}

	// 将去掉前缀的路径拆分为 bucket 和 object
	arr := strings.Split(urlPath, "/")
	if len(arr) < 2 {
		log.Printf("URL格式无效: %s \n", fullURL)
		return ""
	}

	// 拼接数据库存储的相对路径 (bucket/object)
	path = fmt.Sprintf("%s/%s", arr[0], arr[1])
	return path
}
