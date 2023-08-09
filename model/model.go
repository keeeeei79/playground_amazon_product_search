package model

type Doc struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	BulletPoint string `json:"bullet_point"`
	Brand string `json:"brand"`
	Color string `json:"color"`
}

type Query struct {
	Keyword string
}