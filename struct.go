package belajargolangweb

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street   string
	City     string
	Province string
}
