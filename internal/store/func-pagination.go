package store

import (
	"net/http"
	"strconv"
	"time"
)

type PaginatedFeedQuery struct {
	Limit      int    `json:"limit" validate:"gte=1,lte=20"`
	Offset     int    `json:"offset" validate:"gte=0"`
	Category   string `json:"category" validate:"max=100"`
	Sort       string `json:"sort" validate:"oneof=asc desc"`
	Search     string `json:"search" validate:"max=100"`
	Since      string `json:"since"`
	Until      string `json:"until"`
}

func (fq PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return fq, nil
		}

		fq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return fq, nil
		}

		fq.Offset = l
	}

	category := qs.Get("category")
	if category != "" {
		fq.Category = category
	}

	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := qs.Get("since")
	if since != "" {
		fq.Since = parseTime(since)
	}

	until := qs.Get("until")
	if until != "" {
		fq.Until = parseTime(until)
	}

	return fq, nil
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
