package model

import "strconv"

type PaginationOrder = string

const (
	PaginationOrderAsc     PaginationOrder = "ASC"
	PaginationOrderDesc    PaginationOrder = "DESC"
	PaginationOrderDefault PaginationOrder = PaginationOrderDesc
)

const (
	PaginationSizeDefault = 10
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

func AdornPaginationParams(p *PaginationParams) *PaginationParams {
	if p.Size == 0 {
		p.Size = PaginationSizeDefault
	}
	if p.Order == "" {
		p.Order = PaginationOrderDefault
	}

	if p.Page == 0 {
		p.Page = 1
	}
	return p
}

type PaginationData struct {
	*PaginationParams
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}
type PaginationResult struct {
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

type Pagination struct {
	params *PaginationParams
	data   []any
}

func (pg *Pagination) ParamsAny(p any) *PaginationParams {

	// np := &PaginationParams{}
	// raw, e := utils.StructToMap(p, "json")

	// if e == nil {
	// 	np.Order = raw["user"].(int)
	// }

	return pg.params
}
func (pg *Pagination) Query(q PaginationQuery) *PaginationParams {
	p := PaginationParams{}

	page, e := strconv.Atoi(q.Page)

	if e == nil {
		p.Page = page
	}

	size, e := strconv.Atoi(q.Size)
	if e == nil {
		p.Size = size
	}

	p.Order = q.Order

	pg.params = AdornPaginationParams(&p)
	return pg.params
}
func (pg *Pagination) Params(p PaginationParams) *PaginationParams {
	pg.params = AdornPaginationParams(&p)
	return pg.params
}
func (pg *Pagination) GetParams() *PaginationParams {
	return pg.params
}
func (pg *Pagination) Start() int {
	p, s := pg.params.Page, pg.params.Size
	return p*s - s
}

func (pg *Pagination) Data(d []any) *Pagination {
	pg.data = d
	return pg
}
func (pg *Pagination) Result(res any) PaginationData {
	pd := PaginationData{}
	pd.Size = pg.params.Size
	pd.Order = pg.params.Order
	pd.Page = pg.params.Page

	return pd
}
