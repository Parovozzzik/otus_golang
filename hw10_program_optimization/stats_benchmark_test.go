package hw10programoptimization

import (
	"archive/zip"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, err := zip.OpenReader("testdata/users.dat.zip")
		require.NoError(b, err)
		defer r.Close()

		data, err := r.File[0].Open()
		require.NoError(b, err)
		stat, err := GetDomainStat(data, "biz")
		require.NoError(b, err)
		require.Equal(b, ExpectedBizStat, stat)
	}
}
