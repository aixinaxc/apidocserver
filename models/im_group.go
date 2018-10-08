package models

type ImGroup struct {
	GroupId          string `json:"group_id" xorm:"not null pk comment('群组id') VARCHAR(64)"`
	GroupName        string `json:"group_name" xorm:"default '未知' comment('群组名称') VARCHAR(255)"`
	GroupIcon        string `json:"group_icon" xorm:"comment('群组图标') VARCHAR(255)"`
	GroupDescription string `json:"group_description" xorm:"comment('描述') VARCHAR(255)"`
	GroupState       int    `json:"group_state" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
	CreatedAt        int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt        int    `json:"updated_at" xorm:"default 0 comment('更新时间') INT(13)"`
}
