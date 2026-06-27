package pagination_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vapankov/yaca/internal/libs/pagination"
	"github.com/vapankov/yaca/internal/types"
)

func TestPaginateSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil pagination", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{1, 2, 3}
			expectedSlice = []int{1, 2, 3}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				nil,
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})

	t.Run("zero pagination", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{1, 2, 3}
			expectedSlice = []int{}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				&types.Pagination{},
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{}
			expectedSlice = []int{}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				&types.Pagination{
					PageNumber: 1,
					PageSize:   10,
				},
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})

	t.Run("page fits", func(t *testing.T) {
		t.Parallel()

		t.Run("first page", func(t *testing.T) {
			var (
				givenSlice    = []int{1, 2, 3}
				expectedSlice = []int{1, 2}
				actualSlice   = pagination.PaginateSlice(
					givenSlice,
					&types.Pagination{
						PageNumber: 0,
						PageSize:   2,
					},
				)
			)

			require.Equal(
				t,
				expectedSlice,
				actualSlice,
			)
		})

		t.Run("in between", func(t *testing.T) {
			t.Parallel()

			var (
				givenSlice    = []int{1, 2, 3}
				expectedSlice = []int{2}
				actualSlice   = pagination.PaginateSlice(
					givenSlice,
					&types.Pagination{
						PageNumber: 1,
						PageSize:   1,
					},
				)
			)

			require.Equal(
				t,
				expectedSlice,
				actualSlice,
			)
		})

		t.Run("last page", func(t *testing.T) {
			t.Parallel()

			var (
				givenSlice    = []int{1, 2, 3}
				expectedSlice = []int{3}
				actualSlice   = pagination.PaginateSlice(
					givenSlice,
					&types.Pagination{
						PageNumber: 2,
						PageSize:   1,
					},
				)
			)

			require.Equal(
				t,
				expectedSlice,
				actualSlice,
			)
		})
	})

	t.Run("page overlaps with slice end", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{1, 2, 3}
			expectedSlice = []int{3}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				&types.Pagination{
					PageNumber: 1,
					PageSize:   2,
				},
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})

	t.Run("page after slice end", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{1, 2, 3}
			expectedSlice = []int{}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				&types.Pagination{
					PageNumber: 3,
					PageSize:   1,
				},
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})

	t.Run("page covers whole slice", func(t *testing.T) {
		t.Parallel()

		var (
			givenSlice    = []int{1, 2, 3}
			expectedSlice = []int{1, 2, 3}
			actualSlice   = pagination.PaginateSlice(
				givenSlice,
				&types.Pagination{
					PageNumber: 0,
					PageSize:   10,
				},
			)
		)

		require.Equal(
			t,
			expectedSlice,
			actualSlice,
		)
	})
}
