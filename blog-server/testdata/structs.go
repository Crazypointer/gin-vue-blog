package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`              // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转图片链接非法" structs:"href"`         // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"请输入一个合法的图片地址" structs:"images"` // 图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"`        // 是否展示
}

func main() {
	v1 := AdvertRequest{
		Title:  "test",
		Href:   "https://www.baidu.com",
		Images: "https://www.baidu.com",
		IsShow: true,
	}
	m2 := structs.Map(&v1)
	fmt.Println(m2)
}
