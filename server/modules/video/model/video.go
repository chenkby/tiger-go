package model

type Video struct {
	VideoId   int    `json:"video_id" from:"video_id"`
	Name      string `json:"name" form:"name"`
	Image     string `json:"image" from:"image"`
	Video     string `json:"video" form:"video"`
	PlayNum   int    `json:"play_num" form:"play_num"`
	Status    int    `json:"status" form:"status"`
	SizeByte  int    `json:"size_byte" form:"size_byte"`
	Size      string `json:"size" form:"size"`
	Free      int    `json:"free" form:"free"`
	Rank      int    `json:"rank" form:"rank"`
	CreatedAt int    `json:"created_at" form:"created_at"`
	UpdatedAt int    `json:"updated_at" form:"updated_at"`
}

func (*Video) TableName() string {
	return "video"
}
