package docs

import "time"

type Duty struct {
	Done     bool   `json:"done"`
	Executor string `json:"executor,omitempty" bson:"executor,omitempty"`
}

type Routine struct {
	Id            string `json:"id" bson:"_id"`
	Juveniles1    Duty   `json:"juveniles1"`
	AdultFish1    Duty   `json:"adultFish1" bson:"adultFish1"`
	AddSanitizer  Duty   `json:"addSanitizer" bson:"addSanitizer"`
	Collecting1   Duty   `json:"collecting1"`
	Hatching1     Duty   `json:"hatching1"`
	WaterMaking   Duty   `json:"waterMaking" bson:"waterMaking"`
	Standalone    Duty   `json:"standalone"`
	CheckPh       Duty   `json:"checkPh" bson:"checkPh"`
	ReplaceFilter Duty   `json:"replaceFilter" bson:"replaceFilter"`
	Cleaning2     *Duty  `json:"cleaning2,omitempty" bson:"cleaning2,omitempty"`
	Disinfection5 *Duty  `json:"disinfection5,omitempty" bson:"disinfection5,omitempty"`
	Collecting2   Duty   `json:"collecting2"`
	Hatching2     Duty   `json:"hatching2"`
	Juveniles2    Duty   `json:"juveniles2"`
	AdultFish2    Duty   `json:"adultFish2" bson:"adultFish2"`
	FlowRate      Duty   `json:"flowRate" bson:"flowRate"`
	//Newborn       int    `json:"newborn"`
	//Killed        int    `json:"killed"`
	//Current       int    `json:"current"`
	//All           int    `json:"all"`
}

func NewRoutine() *Routine {
	now := time.Now()
	weekday := now.Weekday()

	d := Duty{Done: false}

	r := &Routine{
		Id:            now.Format("2006-01-02"),
		Juveniles1:    d,
		AdultFish1:    d,
		AddSanitizer:  d,
		Collecting1:   d,
		Hatching1:     d,
		WaterMaking:   d,
		Standalone:    d,
		CheckPh:       d,
		ReplaceFilter: d,
		Cleaning2:     nil,
		Disinfection5: nil,
		Collecting2:   d,
		Hatching2:     d,
		Juveniles2:    d,
		AdultFish2:    d,
		FlowRate:      d,
		//Newborn:       0,
		//Killed:        0,
		//Current:       0,
		//All:           0,
	}

	if weekday == time.Tuesday {
		r.Cleaning2 = &d
	}
	if weekday == time.Monday {
		r.Disinfection5 = &d
	}

	return r
}
