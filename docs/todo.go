package docs

type Todo struct {
	Id       string `json:"id" bson:"_id"`
	Creator  string `json:"creator"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt int64  `json:"createAt" bson:"createAt"`
}
