package models

type ApidocUserProject struct {
	UserId    string `json:"user_id" xorm:"default '0' comment('用户id') VARCHAR(64)"`
	ProjectId string `json:"project_id" xorm:"default '0' comment('项目id') VARCHAR(64)"`
}
