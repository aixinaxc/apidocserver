package models

type ApidocApi struct {
	ApiId          string `json:"api_id" xorm:"not null pk comment('api 的id') VARCHAR(64)"`
	SortId         string `json:"sort_id" xorm:"default '0' comment('分类id') VARCHAR(64)"`
	ProjectId      string `json:"project_id" xorm:"default '0' comment('项目id') VARCHAR(64)"`
	ApiName        string `json:"api_name" xorm:"comment('api的名称') VARCHAR(255)"`
	ApiEditContent string `json:"api_edit_content" xorm:"comment('api的编辑内容') TEXT"`
	ApiShowContent string `json:"api_show_content" xorm:"comment('api的展示内容') TEXT"`
	ApiState       int    `json:"api_state" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
	CreatedAt      int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt      int    `json:"updated_at" xorm:"default 0 comment('更新时间') INT(13)"`
}
