package pagination

import "github.com/vapankov/yaca/internal/types"

func PaginateSlice[T any](s []T, p *types.Pagination) []T {
	if p == nil {
		return s
	}

	var (
		sLen   = uint(len(s))
		offset = min(p.PageNumber*p.PageSize, sLen)
		limit  = min(offset+p.PageSize, sLen)
	)

	result := s[offset:limit]

	return result
}
