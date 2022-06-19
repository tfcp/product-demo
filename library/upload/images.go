package upload

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"tfpro/library/utils"
)

func GetImageFullUrl(name string) string {
	return g.Config().GetString("upload.imagePrefixUrl") + "/" + GetImageFullPath() + name
}
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}
func GetImagePath() string {
	return g.Config().GetString("upload.imageSavePath")
}
func GetImageFullPath() string {
	return g.Config().GetString("upload.runtimeRootPath") + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := utils.GetExt(fileName)
	for _, allowExt := range g.Config().GetStrings("upload.imageAllowExts") {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}
func CheckImageSize(f multipart.File) bool {
	size, err := utils.GetSize(f)
	if err != nil {
		return false
	}
	fileSize := (g.Config().GetInt("upload.imageMaxSize")) * 1024 * 1024
	return size <= fileSize
}
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = utils.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := utils.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
