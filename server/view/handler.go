//   File Name:  handler.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 14:04
//    Change Activity:

package view

import (
	"errors"
	"gin-apscheduler/Init"
	"github.com/gin-gonic/gin"
	"github.com/preceeder/apscheduler/job"
)

func AddJob(c *gin.Context, jb Job) (job.Job, error) {
	trigger, err := jb.GetTrigger()
	if err != nil {
		return job.Job{}, errors.New("不支持的触发器类型")
	}

	sjb := job.Job{
		Id:          jb.Id,
		Name:        jb.Name,
		Trigger:     trigger,
		FuncName:    jb.FuncName,
		Args:        jb.Args,
		Timeout:     jb.Timeout,
		StoreName:   jb.StoreName,
		Replace:     jb.Replace,
		MaxInstance: jb.MaxInstance,
	}
	njb, err := Init.Scheduler.AddJob(sjb)

	return njb, err
}

// DelJob 从所有store 中找到 jobId, 然后删除
func DelJob(c *gin.Context, jobId string) error {
	return Init.Scheduler.DeleteJob(jobId)
}

// UpdateJob 除了JobId其他的job参数都可以修改
// 更新也需要给完整的 job数据
func UpdateJob(c *gin.Context, jb Job) (job.Job, error) {
	trigger, err := jb.GetTrigger()
	if err != nil {
		return job.Job{}, errors.New("不支持的触发器类型")
	}

	sjb := job.Job{
		Id:          jb.Id,
		Name:        jb.Name,
		Trigger:     trigger,
		FuncName:    jb.FuncName,
		Args:        jb.Args,
		Timeout:     jb.Timeout,
		StoreName:   jb.StoreName,
		Replace:     jb.Replace,
		MaxInstance: jb.MaxInstance,
	}
	njb, err := Init.Scheduler.UpdateJob(sjb, jb.OldStoreName)

	return njb, err
}

// QueryJobById 根据id 询job, 在知道store name的情况下优先推荐 QueryJobByIdAndStoreName
func QueryJobById(c *gin.Context, jobId string) (job.Job, error) {
	j, err := Init.Scheduler.QueryJob(jobId)
	return j, err
}

// QueryAllJobsByStoreName 查询指定store 下所有的job
func QueryAllJobsByStoreName(c *gin.Context, storeName string) ([]job.Job, error) {
	jobs, err := Init.Scheduler.GetJobsByStoreName(storeName)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// PauseJob 暂停job
func PauseJob(c *gin.Context, jobId string) (job.Job, error) {
	pauseJob, err := Init.Scheduler.PauseJob(jobId)
	if err != nil {
		return job.Job{}, err
	}
	return pauseJob, nil
}

// ResumeJob 回复job
func ResumeJob(c *gin.Context, jobId string) (job.Job, error) {
	resumeJob, err := Init.Scheduler.ResumeJob(jobId)
	if err != nil {
		return job.Job{}, err
	}
	return resumeJob, nil
}

// QueryAllStores 查询所有的 store name
func QueryAllStores(c *gin.Context) []string {
	return Init.Scheduler.GetAllStoreName()
}

// QueryAllFuncName 查询当前注册的所有 方法
func QueryAllFuncName(c *gin.Context) []string {
	var funcNames = make([]string, 0)
	for name, _ := range job.FuncMap {
		funcNames = append(funcNames, name)
	}
	return funcNames
}
