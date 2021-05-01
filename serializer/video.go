package serializer

import "cyylgo/model"

// Video 视频的序列化器
type Video struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Info     string `json:"info"`
	CreateAt int64  `json:"create_at"`
}

// 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:       item.ID,
		Title:    item.Title,
		Info:     item.Info,
		CreateAt: item.CreatedAt.Unix(),
	}
}

// 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos

}
