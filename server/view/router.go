//   File Name:  router.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 15:34
//    Change Activity:

package view

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ApiAddJob @BasePath
// @Summary 添加任务
// @Schemes http
// @Description 添加任务
// @Tags Job
// @Accept json
// @Produce json
// @param Data body view.Job true "body"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job [post]
func ApiAddJob(c *gin.Context) {
	jb := Job{}
	err := c.ShouldBindBodyWith(&jb, binding.JSON)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	njb, err := AddJob(c, jb)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      njb,
	})
}

type RequestJob struct {
	JobId string `form:"job_id"`
}

// ApiDelJob @BasePath
// @Summary 删除任务
// @Schemes http
// @Description 删除任务
// @Tags Job
// @Accept json
// @Produce json
// @param job_id query string true "job id"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job [delete]
func ApiDelJob(c *gin.Context) {
	var params RequestJob
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	err = DelJob(c, params.JobId)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
	})
}

// ApiUpdateJob @BasePath
// @Summary 修改任务
// @Schemes http
// @Description 修改任务
// @Tags Job
// @Accept json
// @Produce json
// @param Data body view.Job true "body"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job [put]
func ApiUpdateJob(c *gin.Context) {
	jb := Job{}
	err := c.ShouldBindBodyWith(&jb, binding.JSON)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	njb, err := UpdateJob(c, jb)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      njb,
	})
}

// ApiGetJobById @BasePath
// @Summary 根据job_id 获取job
// @Schemes http
// @Description 根据job_id 获取job
// @Tags Job
// @Accept json
// @Produce json
// @param job_id query string true " job id"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job [get]
func ApiGetJobById(c *gin.Context) {
	var params RequestJob
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}

	jb, err := QueryJobById(c, params.JobId)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      jb,
	})
}

type RequestStore struct {
	StoreName string `form:"store_name"`
}

// ApiGetJobsByStoreName @BasePath
// @Summary 根据tore_name获取jobs
// @Schemes http
// @Description 根据store_name获取jobs
// @Tags Jobs
// @Accept json
// @Produce json
// @param store_name query string true " store_name"
// @Success 200 {object} HttpResponse{data=[]job.Job}  "响应值"
// @Router /scheduler/jobs [get]
func ApiGetJobsByStoreName(c *gin.Context) {
	var params RequestStore
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}

	jbs, err := QueryAllJobsByStoreName(c, params.StoreName)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      jbs,
	})
}

// ApiPauseJob @BasePath
// @Summary 暂定任务
// @Schemes http
// @Description 暂定任务
// @Tags Job
// @Accept json
// @Produce json
// @param job_id query string true " job_id"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job/pause [put]
func ApiPauseJob(c *gin.Context) {
	var params RequestJob
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}

	jbs, err := PauseJob(c, params.JobId)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      jbs,
	})
}

// ApiResumeJob @BasePath
// @Summary 恢复任务
// @Schemes http
// @Description 恢复任务
// @Tags Job
// @Accept json
// @Produce json
// @param job_id query string true " job_id"
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/job/resume [put]
func ApiResumeJob(c *gin.Context) {
	var params RequestJob
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}

	jbs, err := ResumeJob(c, params.JobId)
	if err != nil {
		c.JSON(200, HttpResponse{ErrorCode: 400, Msg: err.Error()})
		return
	}
	c.JSON(200, HttpResponse{
		ErrorCode: 0,
		Data:      jbs,
	})
}

// ApiGetAllStoreName @BasePath
// @Summary 获取所有的 store name
// @Schemes http
// @Description 获取所有的 store name
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/store/name [get]
func ApiGetAllStoreName(c *gin.Context) {
	stores := QueryAllStores(c)
	c.JSON(200, HttpResponse{ErrorCode: 0, Data: stores})
}

// ApiGetAllFuncNames @BasePath
// @Summary 获取所有的 func name
// @Schemes http
// @Description 获取所有的 func name
// @Tags Func
// @Accept json
// @Produce json
// @Success 200 {object} HttpResponse{}  "响应值"
// @Router /scheduler/func/name [get]
func ApiGetAllFuncNames(c *gin.Context) {
	funcNames := QueryAllFuncName(c)
	c.JSON(200, HttpResponse{ErrorCode: 0, Data: funcNames})
}
