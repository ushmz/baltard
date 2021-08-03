package models

// SimilarwebPage : Similarweb 2000 information with icon URL.
type SimilarwebPage struct {
	// Id : ID of similarweb pages.
	Id int `db:"id" json:"id"`

	// Title : The title of similarweb page.
	Title string `db:"title" json:"title"`

	// Url : Url of the similarweb page.
	Url string `db:"url" json:"url"`

	// Icon : Favicon url of the page.
	Icon string `db:"icon_path" json:"icon"`

	Category string `db:"category" json:"category"`
}
