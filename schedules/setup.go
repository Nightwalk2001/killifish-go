package schedules

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var (
	c *cron.Cron
	d cron.EntryID
	w cron.EntryID
	y cron.EntryID
)

func Setup() {
	c = cron.New()
	d, _ = c.AddFunc("@daily", Daily)
	w, _ = c.AddFunc("@weekly", Weekly)
	y, _ = c.AddFunc("@yearly", Yearly)

	c.Start()
}

func CleanUp() {
	c.Remove(d)
	c.Remove(w)
	c.Remove(y)
	fmt.Println("定时任务清理完成")
}
