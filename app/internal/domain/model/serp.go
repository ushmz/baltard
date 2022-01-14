package model

// SearchPage : Each of search result pages.
type SearchPage struct {
	// PageId : ID of search page.
	PageId int `db:"id" json:"id"`

	// Title : The title of each search result page.
	Title string `db:"title" json:"title"`

	// Url : Url of each search result page.
	Url string `db:"url" json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet" json:"snippet"`
}

// SerpWithIcon : The list of this type struct will be returned as a response of `serp` endpoint.
type SerpWithIcon struct {
	// PageId : ID of search page.
	PageId int `json:"id"`

	// Title : The title of each search result page.
	Title string `json:"title"`

	// Url : Url of each search result page.
	Url string `json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `json:"snippet"`

	// Linked : Users' behavioral data that probably leaked to third party. For more detail, see `Linked` type.
	Linked []LinkedPage `json:"linked"`
}

// SearchPageWithLinkedPage : `SearchPage` with `LinkedPage` query result row struct
type SearchPageWithLinkedPage struct {
	// PageId : ID of search result page.
	PageId int `db:"page_id" json:"page"`

	// LinkedPage : Linked page information with icon URL.
	LinkedPage
}

// SerpWithRatio : The list of this type struct will be returned as a response of `serp` endpoint.
type SerpWithRatio struct {
	// PageId : ID of search page.
	PageId int `json:"id"`

	// Title : The title of each search result page.
	Title string `json:"title"`

	// Url : Url of each search result page.
	Url string `json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `json:"snippet"`

	// Total : Total number of linked pages.
	Total int `json:"total"`

	// Distribution : Distribution information for each categories.
	Distribution []CategoryCount `json:"distribution"`
}

// SerpWithRatioQueryResult : Database select result struct
type SerpWithRatioQueryResult struct {
	// PageId : ID of search page.
	PageId int `db:"id"`

	// Title : The title of each search result page.
	Title string `db:"title"`

	// Url : Url of each search result page.
	Url string `db:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet"`

	// Category : Category name.
	Category string `db:"category"`

	// CategoryRank : DESC rank of category.
	CategoryRank int `db:"category_rank"`

	// CategoryCount : Total number of linked pages in the category.
	CategoryCount int `db:"category_count"`

	// LinkedPageCount : The number of all linked pages.
	LinkedPageCount int `db:"linked_page_count"`

	// CategoryRatio : The percentage of this category.
	CategoryRatio float64 `db:"category_ratio"`
}
