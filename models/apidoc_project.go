package models

type ApidocProject struct {
	ProjectId    string `json:"project_id" xorm:"not null pk comment('项目id') VARCHAR(64)"`
	ProjectName  string `json:"project_name" xorm:"comment('项目名称') VARCHAR(255)"`
	ProjectState int    `json:"project_state" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
	CreatedAt    int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt    int    `json:"updated_at" xorm:"default 0 comment('更新时间') INT(13)"`
}
