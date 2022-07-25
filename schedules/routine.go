package schedules

import (
	"context"

	"killifish/handlers"
	"killifish/mongodb"
)

func InsertRoutine() {
	_, _ = mongodb.Routines.InsertOne(context.TODO(), handlers.NewRoutine())
}
