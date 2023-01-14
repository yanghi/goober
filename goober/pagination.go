package goober

import "strconv"

type PaginationOrder = string

const (
	PaginationOrderAsc     PaginationOrder = "ASC"
	PaginationOrderDesc    PaginationOrder = "DESC"
	paginationOrderDefault PaginationOrder = PaginationOrderDesc
)

const (
	paginationSizeDefault = 10
)

type PaginationParams struct {
	Size  int             `form:"size" json:"size"`
	Page  int             `form:"page" json:"page"`
	Order PaginationOrder `form:"order" json:"order"`
}
type PaginationQuery struct {
	Size  string          `form:"size" json:"size"`
	Page  string          `form:"page" json:"page"`
	Order PaginationOrder `form:"order" json:"order"`
}

type PaginationResult struct {
	Total     int             `json:"total"`
	TotalPage int             `json:"totalPage"`
	Order     PaginationOrder `json:"order"`
	Size      int             `json:"size"`
	Page      int             `json:"page"`
	List      []any           `json:"list"`
}

type Pagination struct {
	params     *PaginationParams
	list       []any
	total      int
	totalPages int
}

func (pg *Pagination) Querys(q PaginationQuery) *Pagination {
	p := pg.params

	page, e := strconv.Atoi(q.Page)

	if e == nil {
		p.Page = page
	}

	size, e := strconv.Atoi(q.Size)
	if e == nil {
		p.Size = size
	}

	p.Order = q.Order
	pg.adjustParams()
	return pg
}

func (pg *Pagination) Params(p *PaginationParams) *PaginationParams {
	pg.params = p
	pg.adjustParams()
	return pg.params
}
func (pg *Pagination) GetParams() *PaginationParams {
	pg.adjustParams()
	return pg.params
}
func (pg *Pagination) adjustParams() *Pagination {
	p := pg.params
	if p.Size == 0 {
		p.Size = paginationSizeDefault
	}
	if p.Order == "" {
		p.Order = paginationOrderDefault
	}

	if p.Page == 0 {
		p.Page = 1
	}

	return pg
}

// Pagination start offset
func (pg *Pagination) Start() int {
	p, s := pg.params.Page, pg.params.Size
	return p*s - s
}

func (pg *Pagination) Size() int {
	if pg.params.Size == 0 {
		return paginationSizeDefault
	}

	return pg.params.Size
}
func (pg *Pagination) Order() PaginationOrder {
	if pg.params.Order != "" {
		return pg.params.Order
	}
	return paginationOrderDefault
}

func (pg *Pagination) List(d []any) *Pagination {
	pg.list = d
	return pg
}
func (pg *Pagination) Total(total int) *Pagination {
	pg.total = total

	totalPages := pg.total / pg.params.Size

	if pg.total%pg.params.Size != 0 {
		totalPages++
	}
	pg.totalPages = totalPages
	return pg
}
func (pg *Pagination) ListMapResult(list any) any {
	// m:= make(map[string]any)
	// res :=

	return map[string]any{
		"page":       pg.params.Page,
		"total":      pg.total,
		"list":       list,
		"totalPages": pg.totalPages,
		"order":      pg.params.Order,
		"size":       pg.Size(),
	}
}
func (pg *Pagination) Result() *PaginationResult {

	res := &PaginationResult{
		Page:      pg.params.Page,
		Total:     pg.total,
		List:      pg.list,
		TotalPage: pg.totalPages,
		Size:      pg.Size(),
	}

	return res
}
func (pg *Pagination) Response() *ResponseResult {
	return OkResult(pg.Result())
}

func NewPagination() *Pagination {
	return &Pagination{params: &PaginationParams{}}
}
