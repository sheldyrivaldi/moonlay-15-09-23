package abstraction

type Pagination struct {
	Page     *int    `json:"page"`
	PageSize *int    `json:"page_size"`
	SortBy   *string `json:"sort_by"`
	Sort     *string `json:"sort"`
}

type PaginationInfo struct {
	*Pagination
	Count       int  `json:"count"`
	MoreRecords bool `json:"more_records"`
}
