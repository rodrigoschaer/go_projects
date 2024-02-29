package entity

type Book struct {
	Id      [16]byte
	Isbn    string
	Title   string
	Author  string
	Pages   uint16
	Year    uint16
	Edition uint8
}
