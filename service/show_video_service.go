package service

import (
	"github.com/lirubao/ozone/model"
	"github.com/lirubao/ozone/serializer"
)

// ShowVideoService 投稿详情的服务
type ShowVideoService struct {
}

// Show 视频详情
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.Find(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
