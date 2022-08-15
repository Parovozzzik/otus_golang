package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("simple case with not empty remainder of the division", func(t *testing.T) {
		from := "./testdata/input.txt"
		to := "./testdata/input_2.txt"
		limit := int64(0)
		offset := int64(0)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("simple case with empty remainder of the division", func(t *testing.T) {
		from := "./testdata/input_6600.txt"
		to := "./testdata/input_6600_2.txt"
		limit := int64(0)
		offset := int64(0)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("limit 1000", func(t *testing.T) {
		from := "./testdata/out_offset0_limit10.txt"
		to := "./testdata/out_offset0_limit10_2.txt"
		limit := int64(10)
		offset := int64(0)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("limit 1000", func(t *testing.T) {
		from := "./testdata/out_offset0_limit1000.txt"
		to := "./testdata/out_offset0_limit1000_2.txt"
		limit := int64(1000)
		offset := int64(0)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("limit 10000", func(t *testing.T) {
		from := "./testdata/out_offset0_limit10000.txt"
		to := "./testdata/out_offset0_limit10000_2.txt"
		limit := int64(10000)
		offset := int64(0)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("offset 100 limit 1000", func(t *testing.T) {
		from := "./testdata/out_offset100_limit1000.txt"
		to := "./testdata/out_offset100_limit1000_2.txt"
		limit := int64(1000)
		offset := int64(100)
		err := Copy(from, to, offset, limit)

		defer os.Remove(to)

		require.Equal(t, err, nil)
	})

	t.Run("offset 6000 limit 1000", func(t *testing.T) {
		from := "./testdata/out_offset6000_limit1000.txt"
		to := "./testdata/out_offset6000_limit1000_2.txt"
		limit := int64(1000)
		offset := int64(6000)
		err := Copy(from, to, offset, limit)
		if err != nil {
			defer os.Remove(to)
		}

		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("urandom", func(t *testing.T) {
		from := "/dev/urandom"
		to := "./testdata/urandom"
		limit := int64(0)
		offset := int64(0)
		err := Copy(from, to, offset, limit)
		if err != nil {
			defer os.Remove(to)
		}

		require.Equal(t, err, ErrUnsupportedFile)
	})
}
