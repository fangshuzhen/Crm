package entity

type DtoApiModel struct {
	Version string `json:"version" binding:"required" example:"1001"`
	Token   string `json:"token" binding:"required" example:"xxxx"`
}

// 分页器数据模板
type Pager struct {
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
	Records interface{} `json:"records"` //分页的数据
}
