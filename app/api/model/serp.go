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
	Snippet string `db:"snippet" json:"snippet"`

	// LeaksSet : Users' behavioral data that probably leaked to third party. For more detail, see `Leaks` type.
	Leaks []SimilarwebPage `json:"leaks"`
}

// SerpWithIconQueryResult : Database select result struct
type SerpWithIconQueryResult struct {
	// PageId : ID of search page.
	PageId int `db:"id"`

	// Title : The title of each search result page.
	Title string `db:"title"`

	// Url : Url of each search result page.
	Url string `db:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet"`

	// SimilarwebId : ID of the similarweb page.
	SimilarwebId int `db:"similarweb_id"`

	// SimilarwebTitle : Title of the similarweb page.
	SimilarwebTitle string `db:"similarweb_title"`

	// SimilarwebUrl : Url of the similarweb page.
	SimilarwebUrl string `db:"similarweb_url"`

	// SimilarwebIcon : Url of the similarweb page favicon.
	SimilarwebIcon string `db:"similarweb_icon"`

	// SimilarwebCategory : Category of the similarweb page.
	SimilarwebCategory string `db:"similarweb_category"`
}

// CategoryCount : Distribution information for each categories.
type CategoryCount struct {
	// Category : Category name.
	Category string `json:"category"`
	// Count : Total number of pages.
	Count int `json:"count"`
	// Percentage : The percentage of this category.
	Percentage float64 `json:"pct"`
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
	Snippet string `db:"snippet" json:"snippet"`

	// Total : Total number of similarweb pages.
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

	// CategoryCount : Total number of similarweb pages in the category.
	CategoryCount int `db:"category_count"`

	// SimilarwebCount : The number of all similarweb pages.
	SimilarwebCount int `db:"similarweb_count"`

	// CategoryDistribution : The percentage of this category.
	CategoryDistribution float64 `db:"category_distribution"`
}
