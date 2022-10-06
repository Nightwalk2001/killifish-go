package docs

type FeedTime struct {
	Time     string `json:"time"`
	Quantity int    `json:"quantity"`
}

type Tank struct {
	Id        string     `json:"id" bson:"_id"`
	Owner     string     `json:"owner"`
	Amount    int        `json:"amount"`
	Size      float64    `json:"size"`
	FeedTimes []FeedTime `json:"feedTimes,omitempty" bson:"feedTimes,omitempty"`
	Sexual    string     `json:"sexual,omitempty" bson:"sexual,omitempty"`
	Birthday  string     `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Genotype  string     `json:"genotype,omitempty" bson:"genotype,omitempty"`
	Species   string     `json:"species,omitempty" bson:"species,omitempty"`
	Label     string     `json:"label,omitempty" bson:"label,omitempty"`
}

type Statistics struct {
	Group  string `json:"group" bson:"_id"`
	Tanks  int    `json:"tanks"`
	Fishes int    `json:"fishes"`
}

type Facets struct {
	Size []struct {
		Group  float64 `json:"group" bson:"_id"`
		Tanks  int     `json:"tanks"`
		Fishes int     `json:"fishes"`
	} `json:"size"`
	Sexual   []Statistics `json:"sexual"`
	Genotype []Statistics `json:"genotype"`
	Species  []Statistics `json:"species"`
	Birthday []Statistics `json:"birthday"`
}
