//   File Name:  base.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 14:03
//    Change Activity:

package view

import (
	"errors"
	"github.com/preceeder/apscheduler/triggers"
)

type Job struct {
	// 任务的唯一id.
	Id string `json:"id"` // 一旦设置,不能修改
	// job name
	Name            string                    `json:"name"`
	IntervalTrigger *triggers.IntervalTrigger `json:"interval_trigger,omitempty"` // cron_trigger, date_trigger,interval_trigger 只能有一个生效
	DateTrigger     *triggers.DateTrigger     `json:"date_trigger,omitempty"`     // 同时出现的时候优先使用 interval_trigger > date_trigger > cron_trigger
	CronTrigger     *triggers.CronTrigger     `json:"cron_trigger,omitempty"`
	FuncName        string                    `json:"func_name"` // 必须和注册的函数名一致
	// Arguments for `Func`.
	Args map[string]any `json:"args"`
	// jobStoreName
	StoreName    string `json:"store_name"`     // 存储的store name
	OldStoreName string `json:"old_store_name"` // 更新的时候需要
	Replace      bool   `json:"replace"`        // 任务存在是否更新 默认false
	MaxInstance  int    `json:"max_instance"`   // 改任务可以同时存在的个数 最少1个, 默认 1
	// 任务执行超时时间 ms Default: `3600*1000`
	Timeout int64 `json:"timeout"`
}

func (j Job) GetTrigger() (triggers.Trigger, error) {
	if j.IntervalTrigger != nil {
		return j.IntervalTrigger, nil
	}

	if j.DateTrigger != nil {
		return j.DateTrigger, nil
	}

	if j.CronTrigger != nil {
		return j.CronTrigger, nil
	}

	return nil, errors.New("触发器错误")
}
