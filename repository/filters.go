package repository

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Filter interface {
	Apply(ctx *gin.Context) []Options
}

type APIFilter struct {
	ApplyFunc		func(ctx *gin.Context) []Options
}

func (a *APIFilter) Apply(ctx *gin.Context) []Options {
	return a.ApplyFunc(ctx)
}

func LimitFilter() Filter {
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		var options []Options
		if limitStr := ctx.Query("limit"); limitStr != "" {
			limit, err := strconv.ParseInt(limitStr, 10, 64)
			if err != nil {
				ctx.Status(400)
				return options
			}
			if limit > 20 {
				limit = 20
			}
			options = append(options, WithLimit(int(limit)))
		}
		return options
	}}
}

func LimitAndPageFilter() Filter {
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		var options []Options
		var err error
		var limit int64
		if limitStr := ctx.Query("limit"); limitStr != "" {
			limit, err = strconv.ParseInt(limitStr, 10, 64)
			if err != nil {
				return options
			}
			if limit > 20 {
				limit = 20
			}
			options = append(options, WithLimit(int(limit)))
		}
		if pageStr := ctx.Query("page"); pageStr != "" && limit != 0 {
			page, err := strconv.ParseInt(pageStr, 10, 64)
			if err != nil {
				return options
			}
			var offset = (limit * page) - limit
			if offset >= 0 {
				options = append(options, WithOffset(int(offset)))
			}
		}
		return options
	}}
}
