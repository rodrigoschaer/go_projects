package dto

import (
	"github.com/rodrigoschaer/go_projects/bookshelf_services/internal/entity"
)

type BookDTO struct {
	Id      [16]byte `json: "id"`
	Isbn    string   `json: "isbn"`
	Title   string   `json: "title"`
	Author  string   `json: "author"`
	Pages   uint16   `json: "pages"`
	Year    uint16   `json: "year"`
	Edition uint8    `json: "edition"`
}
