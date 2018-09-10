package models

type ApidocSort struct {
	SortId    string `json:"sort_id" xorm:"not null pk comment('分类id') VARCHAR(64)"`
	SortTopId string `json:"sort_top_id" xorm:"default '0' comment('分类最上级id （默认0）') VARCHAR(64)"`
	SortPid   string `json:"sort_pid" xorm:"default '0' comment('分类上级id （默认0）') VARCHAR(64)"`
	ProjectId string `json:"project_id" xorm:"default '0' comment('项目id') VARCHAR(64)"`
	SortName  string `json:"sort_name" xorm:"default '无' comment('分类名称') VARCHAR(255)"`
	SortSeq   int    `json:"sort_seq" xorm:"default 1000 comment('分类排序') INT(11)"`
	Status    int    `json:"status" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
	CreatedAt int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt int    `json:"updated_at" xorm:"default 0 comment('更新时间') INT(13)"`
}
