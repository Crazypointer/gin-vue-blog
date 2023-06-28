package common

import (
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// 封装列表查询
// 支持分页

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	// 如果开启了debug模式 则打印sql语句
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 如果没有传入排序规则 则默认按照创建时间往前排
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照创建时间往前排
	}
	query := DB.Where(model)
	// 表中记录数
	count = query.Select("id").Find(&list).RowsAffected
	// 后面的query会首前面query的影响
	// 所以重新查询一次 复位query
	query = DB.Where(model)

	// 分页查询参数：偏移量offset 每页条数limit
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	//如果没有传入limit 则不分页
	if option.Limit == 0 {
		err = query.Offset(offset).Order(option.Sort).Find(&list).Error
	} else {
		err = query.Offset(offset).Limit(option.Limit).Order(option.Sort).Find(&list).Error
	}
	return list, count, err
}
