package entity

type DtoHello struct {
	DtoApiModel `json:",inline"`
	Message     string `json:"message" binding:"required" example:"message...."`
}

type DtoUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type DtoFindUser struct {
	Name string `json:"name"`
}

//获取列表
type DtoPageList struct {
	Keyword string `json:"keyword"` //搜索的关键字
	Page    int    `json:"page"`    //页码
	Size    int    `json:"size"`    //每页显示多少条
}
