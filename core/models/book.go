package models

import "time"

type Book struct {
	URL             string
	Name            string
	Price           int
	ISBN            string
	PublicationDate time.Time
}
