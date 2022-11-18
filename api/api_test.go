package api

import (
	"testing"
)

func BenchmarkGetDecodeYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDecodeYAML("https://raw.githubusercontent.com/firo-18/pnc-db/main/data/dolls/Evelyn.yaml", &struct{}{})
	}
}

func BenchmarkGetDecodeJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDecodeJSON("https://raw.githubusercontent.com/firo-18/pnc-db/2ddf05fb08deae1c37c6bbdac00a036bfdc25318/data/dolls/Evelyn.json", &struct{}{})
	}
}
