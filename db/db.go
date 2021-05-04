package db

type Location struct {
	Country string
	City string
}

type Db interface {
	GetLocation(ip string) Location
}