package service

import (
	"cyylgo/model"
	"cyylgo/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=200"`
	URL   string `form:"url" json:"url"`
}

// Create 创建视频的函数
func (service *CreateVideoService) Create() serializer.Response {

	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
		URL:   service.URL,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{

		Data: serializer.BuildVideo(video),
	}
}
