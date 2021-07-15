package models

// SerpRange : Search page id range.
type SerpRange struct {
	// Min : Minimum Id value
	Min int `db:"min_id"`
	// Max : Maxium Id value
	Max int `db:"max_id"`
}

// SerpSimilarwebRelation : Relation between search page id and similarweb id on SERP.
type SerpSimilarwebRelation struct {
	// PageId : ID of search page.
	PageId int `db:"page_id"`
	// SimilarwebId : ID of similarweb page.
	SimilarwebId int `db:"similarweb_id"`
	// TaskId : ID of task.
	TaskId int `db:"task_id"`
}

// SearchPage : Each of search result pages.
type SearchPage struct {
	// PageId : ID of search page.
	// This is assumed to be used as `key` of search result in front-side app (savitr).
	PageId int `db:"id" json:"id"`

	// Title : The title of each search result page.
	Title string `db:"title" json:"title"`

	// Url : Url of each search result page.
	Url string `db:"url" json:"url"`

	// Snippet : Snippet of each search result page.
	Snippet string `db:"snippet" json:"snippet"`
}

// SearchPageWithLeaks : The list of this type struct will be returned as a response of `serp` endpoint.
type Serp struct {
	// PageId : This is assumed to be used as `key` of search result in front-side app (savitr).
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
