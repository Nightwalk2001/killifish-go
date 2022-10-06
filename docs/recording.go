package docs

type Recording struct {
	Tank     string `json:"tank"`
	Owner    string `json:"owner"`
	Genotype string `json:"genotype,omitempty" bson:"genotype,omitempty"`
	Sexual   string `json:"sexual,omitempty" bson:"sexual,omitempty"`
	Birthday string `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Quantity int    `json:"quantity"`
	Trigger  string `json:"trigger"`
	Time     int64  `json:"time"`
	Succeed  bool   `json:"succeed"`
}
