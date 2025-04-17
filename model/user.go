package model

import (
	"time"
)

// BaseModel 定义了基础模型字段
type BaseModel struct {
	ID        int32     `gorm:"primarykey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
	CreatedBy int32     `gorm:"column:created_by" json:"created_by"` // 创建人ID
	UpdatedBy int32     `gorm:"column:updated_by" json:"updated_by"` // 更新人ID
	DeletedBy int32     `gorm:"column:deleted_by" json:"deleted_by"` // 删除人ID
	IsDeleted bool      `gorm:"column:is_deleted" json:"is_deleted"` // 是否删除
}

// User 定义了用户模型
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;not null;type:varchar(11);column:mobile"`        // 手机号
	NickName string     `gorm:";max_length:20;min_length:6;column:nick_name;type:varchar(100)"`         // 用户名
	Password string     `gorm:"not null;max_length:20;min_length:6;type:varchar(1000);column:password"` // 密码
	Birthday *time.Time `gorm:"type:datetime;column:birthday"`                                          // 生日
	Gender   string     `gorm:"type:varchar(6);default:male;column:gender;comment:'male or female'"`    // 性别
	Role     string     `gorm:"type:varchar(10);default:user;column:role;comment:'user or admin'"`      // 角色
}
