package handlers

import "time"

type Routine struct {
	Id            string `json:"id" bson:"_id"`
	Juveniles1    bool   `json:"juveniles1"`
	AdultFish1    bool   `json:"adultFish1" bson:"adultFish1"`
	AddSanitizer  bool   `json:"addSanitizer" bson:"addSanitizer"`
	Collecting1   bool   `json:"collecting1"`
	Hatching1     bool   `json:"hatching1"`
	WaterMaking   bool   `json:"waterMaking" bson:"waterMaking"`
	Standalone    bool   `json:"standalone"`
	CheckPh       bool   `json:"checkPh" bson:"checkPh"`
	ReplaceFilter bool   `json:"replaceFilter" bson:"replaceFilter"`
	Cleaning2     bool   `json:"cleaning2" bson:"cleaning2"`
	Disinfection5 bool   `json:"disinfection5" bson:"disinfection5"`
	Collecting2   bool   `json:"collecting2"`
	Hatching2     bool   `json:"hatching2"`
	Juveniles2    bool   `json:"juveniles2"`
	AdultFish2    bool   `json:"adultFish2" bson:"adultFish2"`
	FlowRate      bool   `json:"flowRate" bson:"flowRate"`
	Newborn       int    `json:"newborn"`
	Killed        int    `json:"killed"`
	Current       int    `json:"current"`
	All           int    `json:"all"`
}

func NewRoutine() *Routine {
	now := time.Now()
	weekday := now.Weekday()

	r := &Routine{
		Id:            now.Format("2006-01-02"),
		Juveniles1:    false,
		AdultFish1:    false,
		AddSanitizer:  false,
		Collecting1:   false,
		Hatching1:     false,
		WaterMaking:   false,
		Standalone:    false,
		CheckPh:       false,
		ReplaceFilter: false,
		Cleaning2:     false,
		Disinfection5: false,
		Collecting2:   false,
		Hatching2:     false,
		Juveniles2:    false,
		AdultFish2:    false,
		FlowRate:      false,
		Newborn:       0,
		Killed:        0,
		Current:       0,
		All:           0,
	}

	if weekday == time.Tuesday {
		r.Cleaning2 = false
	}
	if weekday == time.Monday {
		r.Disinfection5 = false
	}

	return r
}
