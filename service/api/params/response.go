package params

import (
	"github.com/Confialink/wallet-pkg-list_params"
	"math"
)

type Error struct {
	Title   string       `json:"title"`
	Details string       `json:"details"`
	Status  int          `json:"status"`
	Code    *string      `json:"code"`
	Source  *ErrorSource `json:"source"`
}

type ErrorSource struct {
	Pointer string `json:"pointer"`
}

type List struct {
	HasMore bool        `json:"hasMore"`
	Items   interface{} `json:"items"`
}

type Response struct {
	Links      interface{} `json:"links,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Messages   []string    `json:"messages,omitempty"`
	Errors     []*Error    `json:"errors,omitempty"`
}

type Pagination struct {
	TotalRecord uint64 `json:"totalRecord"`
	TotalPage   uint64 `json:"totalPage"`
	Limit       uint64 `json:"limit"`
	PageNumber  uint64 `json:"currentPage"`
}

func New() *Response {
	return new(Response)
}

func NewWithListAndLinksAndPagination(items interface{}, total uint64, listParams *list_params.ListParams) *Response {
	res := New()
	res.SetData(items)

	pagination := paginate((uint64)(listParams.Pagination.PageNumber),
		(uint64)(listParams.Pagination.PageSize), total)
	res.SetPagination(pagination)

	return res
}

func (r *Response) AddMessage(message string) *Response {
	r.Messages = append(r.Messages, message)
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetLinks(links interface{}) *Response {
	r.Links = links
	return r
}

func (r *Response) SetPagination(pagination *Pagination) *Response {
	r.Pagination = pagination
	return r
}

func paginate(number, size, total uint64) *Pagination {
	totalPage := uint64(math.Ceil(float64(total) / float64(size)))

	return &Pagination{
		Limit:       size,
		PageNumber:  number,
		TotalRecord: total,
		TotalPage:   totalPage,
	}
}
