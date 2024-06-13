//   File Name:  aps.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 15:48
//    Change Activity:

package Init

import (
	"fmt"
	"github.com/preceeder/apscheduler"
	"github.com/preceeder/apscheduler/events"
	"github.com/preceeder/apscheduler/stores"
)

var Scheduler *apscheduler.Scheduler

// Init 这里需要更具自己的需求修改
func Init() {
	Scheduler = apscheduler.NewScheduler()
	events.RegisterEvent(events.EVENT_JOB_ERROR|events.EVENT_JOB_ADDED|events.EVENT_JOB_REMOVED|events.EVENT_JOB_MODIFIED, func(ei events.EventInfo) {
		fmt.Println("lisenter:--->", ei)
	})

	err := Scheduler.SetStore(stores.DefaultName, &stores.MemoryStore{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Scheduler.Start()

}
