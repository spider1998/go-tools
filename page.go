package tools

import (
	"github.com/go-ozzo/ozzo-routing"
	"strconv"
)

const (
	DefaultPageSize int = 100
	MaxPageSize     int = 1000
)

type pager struct {
	page       int
	pageSize   int
	totalCount int
}

func (p *pager) offset() int {
	return (p.page - 1) * p.pageSize
}

func (p *pager) limit() int {
	return p.pageSize
}

func newPager(page, perPage, total int) *pager {
	if perPage < 1 {
		perPage = 100
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &pager{
		page:       page,
		pageSize:   perPage,
		totalCount: total,
	}
}

func getPaginatedListFromRequest(c *routing.Context, count int) *pager {
	page := parseInt(c.Query("page"), 1)
	pageSize := parseInt(c.Query("page_size"), DefaultPageSize)
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	p := newPager(page, pageSize, count)

	c.Response.Header().Set("X-Page-Total", strconv.Itoa(p.totalCount))
	c.Response.Header().Set("X-Page", strconv.Itoa(p.page))
	c.Response.Header().Set("X-Page-Size", strconv.Itoa(p.pageSize))

	return p
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}


/*eg:
	pager := getPaginatedListFromRequest(c, len(lists))
		if pager.offset()+pager.limit() <= pager.totalCount {
			return c.Write(lists[pager.offset() : pager.offset()+pager.limit()])
		} else {
			return c.Write(lists[pager.offset():pager.totalCount])
			}*/
