package model

type Post struct {
	PostId      int64
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   int64
	UpdatedAt   int64
	Creator     string
	Dummy       byte // duummy field to just sort by index createdat
}
