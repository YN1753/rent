package oss

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
)

var Client *oss.Client
var Bucket *oss.Bucket

func InitOSS() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("加载env失败")
	}
	oss_access_key_id := os.Getenv("OSS_ACCESS_KEY_ID")
	oss_access_key_secret := os.Getenv("OSS_ACCESS_KEY_SECRET")
	oss_endpoint := os.Getenv("OSS_ACCESS_ENDPOINT")
	oss_bucket := os.Getenv("OSS_BUCKET_NAME")
	Client, err = oss.New(oss_endpoint, oss_access_key_id, oss_access_key_secret)
	if err != nil {
		fmt.Println("连接oss失败")
	}
	Bucket, err = Client.Bucket(oss_bucket)
}

func CreateSignUrl(objectKey string) ([]string, error) {
	ls, err := Bucket.ListObjects(oss.Prefix(objectKey))
	if err != nil {
		return nil, err
	}
	var urls []string
	for _, l := range ls.Objects {
		temp, _ := Bucket.SignURL(l.Key, oss.HTTPGet, 3600)
		urls = append(urls, temp)
	}
	if len(urls) == 0 {
		return urls, errors.New("生成签名url错误")
	}
	return urls, nil
}

// 上传图片或者视频到oss（视频会压缩）
func UploadFile(objectKey string, file *multipart.FileHeader) error {
	tempfile, err := file.Open()
	if err != nil {
		return errors.New("打开文件失败")
	}
	defer tempfile.Close()

	ext := path.Ext(file.Filename)
	if ext == ".jpg" || ext == "png" || ext == "jpeg" {
		err = Bucket.PutObject(objectKey, tempfile)
		if err != nil {
			return errors.New("oss上传失败")
		}
	} else {
		loacltemp, err := os.Create(objectKey)
		if err != nil {
			return errors.New("生成临时文件失败")
		}
		_, err = io.Copy(loacltemp, tempfile)
		if err != nil {
			loacltemp.Close()
			os.Remove(loacltemp.Name())
			return errors.New("写入临时文件失败")
		}
		loacltemp.Close()

		err = loacltemp.Sync()
		if err != nil {
			os.Remove(loacltemp.Name())
			return errors.New("同步临时文件失败")
		}
		tempOutputPath := loacltemp.Name() + "_compressed.mp4"
		ffmpegArgs := []string{
			"-i", loacltemp.Name(),
			"-c:v", "libx264",
			"-c:a", "aac",
			"-b:a", "64k",
			"-vf", "scale=1280:720:force_original_aspect_ratio=decrease",
			"-movflags", "+faststart",
			"-y",
			tempOutputPath,
		}
		cmd := exec.Command("ffmpeg", ffmpegArgs...)
		err = cmd.Run()
		if err != nil {
			return errors.New("视频压缩失败")
		}
		err = Bucket.PutObjectFromFile(objectKey, tempOutputPath)
		if err != nil {
			return errors.New("上传oss失败")
		}
	}
	return nil
}
