package docs

type Alert struct {
	Id      string `json:"id" bson:"_id"`
	Tank    string `json:"tank"`
	Message string `json:"message"`
	Time    string `json:"time"`
}
