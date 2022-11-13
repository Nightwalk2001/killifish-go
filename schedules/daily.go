package schedules

import (
	"encoding/json"

	"killifish/docs"
	"killifish/redis"
)

func Daily() {
	j, _ := redis.ReJson.JSONGet("state", ".")

	r := docs.State{}
	_ = json.Unmarshal(j.([]byte), &r)

	r.Born = 0
	r.Killed = 0

	_, _ = redis.ReJson.JSONSet("state", ".", r)
}
