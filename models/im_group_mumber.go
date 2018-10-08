package models

type ImGroupMumber struct {
	GroupMumberId string `json:"group_mumber_id" xorm:"not null pk comment('群组成员id') VARCHAR(64)"`
	GroupId       string `json:"group_id" xorm:"default '0' comment('群组id') VARCHAR(64)"`
	UserId        string `json:"user_id" xorm:"default '0' comment('用户id') VARCHAR(64)"`
	MumberType    int    `json:"mumber_type" xorm:"default 3 comment('成员类型（1群主，2管理员，3群员）') INT(1)"`
	MemberState   int    `json:"member_state" xorm:"default 1 comment('成员状态（1正常，2禁言）') INT(1)"`
	CreatedAt     int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt     int    `json:"updated_at" xorm:"default 0 comment('更新时间') INT(13)"`
}
