package models

type ApidocUser struct {
	UserId       string `json:"user_id" xorm:"not null pk comment('用户id') VARCHAR(64)"`
	UserUsername string `json:"user_username" xorm:"comment('用户名') VARCHAR(255)"`
	UserPassword string `json:"user_password" xorm:"comment('密码') VARCHAR(255)"`
	CreatedAt    int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt    int    `json:"updated_at" xorm:"default 0 comment('修改时间') INT(13)"`
	UserState    int    `json:"user_state" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
}
