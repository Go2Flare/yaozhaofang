package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	//	初始化对象
	cap := captcha.New()
	//	设置字体
	cap.SetFont("./conf/comic.ttf")

	//	设置验证码大小
	cap.SetSize(128, 64)
	//	设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//	设置前景色
	cap.SetFrontColor(color.RGBA{255, 245, 247, 255})
	//	设置背景色
	cap.SetBkgColor(color.RGBA{101, 147, 74 , 128}, color.RGBA{69, 137, 148, 255})
//	生成字体 ,将验证码展示到网页
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(4, captcha.NUM)
		png.Encode(w, img)
		println(str)
	})
	http.ListenAndServe(":8085", nil)
}
