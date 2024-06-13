//   File Name:  test.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 17:52
//    Change Activity:

package task

import (
	"context"
	"github.com/preceeder/apscheduler/job"
	"log/slog"
)

func init() {
	job.RegisterJobsFunc(job.FuncInfo{Func: Test, Name: "task.test:Test"})
}
func Test(ctx context.Context, j job.Job) {
	slog.Info("run job", "jobName", j.Name)
	//time.Sleep(time.Second * 9)
}
