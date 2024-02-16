package responses

type File struct {
	ID		string	`json:"id" bson:"_id"`
	Name	string	`json:"name"`
}
