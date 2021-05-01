package api

import (
	"cyylgo/service"
	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(ctx *gin.Context) {

	service := service.CreateVideoService{}
	if err := ctx.ShouldBind(&service); err != nil {
		ctx.JSON(200, ErrorResponse(err))
	} else {
		res := service.Create()
		ctx.JSON(200, res)
	}

}

// ShowVideo 视频详情接口
func ShowVideo(ctx *gin.Context) {
	service := service.ShowVideoService{}
	res := service.Show(ctx.Param("id"))
	ctx.JSON(200, res)

}

// ListVideo 视频列表接口
func ListVideo(ctx *gin.Context) {
	service := service.ListVideoService{}
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.List()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, ErrorResponse(err))
	}

}

// UpdateVideo 更新视频接口
func UpdateVideo(ctx *gin.Context) {
	service := service.UpdateVideoService{}
	if err := ctx.ShouldBind(&service); err != nil {
		ctx.JSON(200, ErrorResponse(err))
	} else {
		res := service.Update(ctx.Param("id"))
		ctx.JSON(200, res)
	}
}

// DeleteVideo 删除视频接口

func DeleteVideo(ctx *gin.Context) {
	service := service.DeleteVideooService{}
	res := service.Delete(ctx.Param("id"))
	ctx.JSON(200, res)

}
