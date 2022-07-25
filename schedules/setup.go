package schedules

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var (
	c   *cron.Cron
	id1 cron.EntryID
	id2 cron.EntryID
)

func Setup() {
	c = cron.New()
	id1, _ = c.AddFunc("@daily", InsertRoutine)
	id2, _ = c.AddFunc("@weekly", WeeklyReport)

	c.Start()
}

func CleanUp() {
	c.Remove(id1)
	c.Remove(id2)
	fmt.Println("定时任务清理完成")
}
