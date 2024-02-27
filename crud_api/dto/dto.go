package dto

type Movie struct {
	Id       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Id        string `json: "id"`
	Firstname string `json: "first_name"`
	Lastname  string `json: "last_name"`
}
