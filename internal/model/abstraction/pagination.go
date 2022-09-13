package abstraction

type Pagination struct {
	Page  *int `query:"page" json:"page"`
	Limit *int `query:"page_size" json:"page_size"`
}

type PaginationInfo struct {
	Pagination
	Count       int64 `json:"count"`
	TotalPage   int64 `json:"total_page"`
	MoreRecords bool  `json:"more_records"`
}
