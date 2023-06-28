package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`           // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                                 // 更新时间
}

// 接收批量删除

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

// 分页参数结构体

type PageInfo struct {
	Page  int    `form:"page"`  // 当前页
	Key   string `form:"key"`   // 查询关键字
	Limit int    `form:"limit"` //每页条数
	Sort  string `form:"sort"`
}
