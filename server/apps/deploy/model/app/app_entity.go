package app

import "time"

type DeployApp struct {
	AppId      int64     `json:"app_id" xorm:"not null pk autoincr comment('应用ID') BIGINT(20)"`
	AppName    string    `json:"app_name" xorm:"not null comment('应用名称') VARCHAR(128)"`
	AppDesc    string    `json:"app_desc" xorm:"comment('应用描述') VARCHAR(1024)"`
	Status     string    `json:"status" xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`

	TenantId   int64     `json:"tenant_id" xorm:"default 0 comment('租户id') BIGINT(20)"`
}


func TableName() string {
	return "d_app"
}

