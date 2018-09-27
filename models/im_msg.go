package models

type ImMsg struct {
	MsgId          string `json:"msg_id" xorm:"not null pk comment('消息id') VARCHAR(64)"`
	MsgType        string `json:"msg_type" xorm:"comment('消息类型(client:连接，p2p：一对一，group：群)') VARCHAR(10)"`
	MsgFromId      string `json:"msg_from_id" xorm:"default '0' comment('消息来源id') VARCHAR(64)"`
	MsgFromContent Content `json:"msg_from_content" xorm:"comment('来源人的信息') TEXT"`
	MsgToId        string `json:"msg_to_id" xorm:"default '0' comment('消息接受人的id') VARCHAR(64)"`
	MsgToContent   Content `json:"msg_to_content" xorm:"comment('消息接收人的信息') TEXT"`
	MsgContent     IMMsgContent `json:"msg_content" xorm:"comment('消息内容') TEXT"`
	MsgContentType string `json:"msg_content_type" xorm:"comment('消息内容的类型（im_text,im_img,im_audio）') VARCHAR(10)"`
	MsgStatus      string `json:"msg_status" xorm:"comment('消息状态') VARCHAR(10)"`
	Status         int    `json:"status" xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(2)"`
	CreatedAt      int    `json:"created_at" xorm:"default 0 comment('创建时间') INT(11)"`
}

type Content struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type IMMsgContent struct{
	Text string `json:"text"`
	FileName string `json:"file_name"`
	FileContent string `json:"file_content"`
	File []byte `json:"file"`
	Length int `json:"length"`
}

type ImGroupMember struct {
	Group_member_id string `json:"group_member_id"`
	Group_id string `json:"group_id"`
	User_id string `json:"user_id"`
	Created_at int `json:"created_at"`
}

type Msg struct {
	Id string `json:"id"`
	MsgId string `json:"msg_id"`
	Msg []byte `json:"msg"`
}