package model

type User struct {
	ID       int64  `json:"id"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Nickname string `json:"nick_name"`
	HeadUrl  string `json:"head_url"`
	Birthday int    `json:"birthday"`
	Address  string `json:"address"`
	Desc     string `json:"desc"`
	Gender   uint8  `json:"gender"`
	Role     uint8  `json:"role"`
}
