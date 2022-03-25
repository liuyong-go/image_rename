package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/spf13/cast"
)

var imageType = []string{".jpg", ".png", ".jpeg"}
var index = 1
var suffix string
var canRun = "0"

func main() {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}
	dir := filepath.Dir(path)
	fmt.Println("当前执行目录为:", dir, " 回复1继续执行，回复0暂停执行")
	fmt.Scanf("%s", &canRun)
	if canRun != "1" {
		fmt.Println("停止执行")
		return
	}
	// dir, _ := os.Getwd()
	// fmt.Println(dir)
	fmt.Println("请输入数字后缀,文件命名为1{后缀}.jpg形式")
	fmt.Scanf("%s", &suffix)
	fmt.Println("数字后缀为:", suffix)

	fmt.Println("执行目录:", dir)

	fmt.Println("请输入索引起始数字")
	fmt.Scanf("%d", &index)
	fmt.Println("起始索引为:", index)

	files, _ := ioutil.ReadDir(dir)
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		renameFile(dir, fi.Name())
	}
	fmt.Println("执行完成，回车结束")
}

func renameFile(dir string, filename string) error {
	fileExt := ".jpg"
	var isImage = false
	for _, t := range imageType {
		if strings.Contains(filename, t) {
			isImage = true
			fileExt = t
			break
		}
	}
	if !isImage {
		return nil
	}
	newName := cast.ToString(index) + suffix + fileExt
	newPath := dir + "/" + newName
	fmt.Println("file: ", filename, " toName: ", newName)
	os.Rename(dir+"/"+filename, newPath)
	resizeImage(newPath)
	index++
	return nil
}
func resizeImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open", path, err)
	}
	// imgInfo, _, _ := image.DecodeConfig(file)
	fmt.Println(path, "图片宽度>1080执行压缩..")
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("decode", path, err)
		return
	}
	file.Close()
	b := img.Bounds()
	width := b.Max.X
	if width > 1080 {
		fmt.Println("width", width, ">1080,执行压缩...")
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1080, 0, img, resize.Lanczos3)

		out, err := os.Create(path)
		if err != nil {
			fmt.Println("create", path, err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

}
