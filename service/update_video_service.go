package service

import (
	"github.com/lirubao/ozone/model"
	"github.com/lirubao/ozone/serializer"
)

// UpdateVideoService 视频更新的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 视频更新
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
