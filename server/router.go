//   File Name:  router.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 11:37
//    Change Activity:

package server

import (
	"gin-apscheduler/server/view"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.POST("/scheduler/job", view.ApiAddJob)
	r.DELETE("/scheduler/job", view.ApiDelJob)
	r.PUT("/scheduler/job", view.ApiUpdateJob)
	r.GET("/scheduler/job", view.ApiGetJobById)
	r.GET("/scheduler/jobs", view.ApiGetJobsByStoreName)
	r.PUT("/scheduler/job/pause", view.ApiPauseJob)
	r.PUT("/scheduler/job/resume", view.ApiResumeJob)
	r.GET("/scheduler/store/name", view.ApiGetAllStoreName)
	r.GET("/scheduler/func/name", view.ApiGetAllFuncNames)

	return r
}
