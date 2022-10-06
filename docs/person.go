package docs

type Person struct {
	Name      string `json:"name" bson:"_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	IsManager bool   `json:"isManager" bson:"isManager"`
	Label     string `json:"label,omitempty" bson:"label,omitempty"`
	Tags      string `json:"tags,omitempty" bson:"tags,omitempty"`
}
