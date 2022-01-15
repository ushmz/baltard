package model

// LinkedPage : Linked page information with icon URL.
type LinkedPage struct {
	// Id : ID of linked page.
	Id int `db:"id" json:"id"`

	// Title : The title of linked page.
	Title string `db:"title" json:"title"`

	// Url : Url of the linked page.
	Url string `db:"url" json:"url"`

	// Icon : Favicon url of the page.
	Icon string `db:"icon_path" json:"icon"`

	// Category : Category name of linked page.
	Category string `db:"category" json:"category"`
}

type LinkedPageRatio struct {
	PageId        int    `db:"page_id"`
	Category      string `db:"category"`
	CategoryCount int    `db:"category_count"`
}

// CategoryCount : Distribution information for each categories.
type CategoryCount struct {
	// Category : Category name.
	Category string `json:"category"`
	// Count : Total number of pages.
	Count int `json:"count"`
}
