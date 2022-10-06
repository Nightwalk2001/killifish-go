package schedules

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
	d cron.EntryID
	y cron.EntryID
)

func Setup() {
	c = cron.New()
	d, _ = c.AddFunc("@daily", DailyReset)
	y, _ = c.AddFunc("@yearly", func() {})

	c.Start()
}

func CleanUp() {
	c.Remove(d)
	c.Remove(y)
	fmt.Println("定时任务清理完成")
}
