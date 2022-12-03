package crawl

import (
	"testing"
)

func TestCrawl(t *testing.T) {
	testcases := []struct {
		nURLs, nWorkers, sPerURL, cCap int
	}{
		{10, 5, 3, 5},
		{1, 1, 1, 1},
		{100, 5, 1, 5},
	}
	for _, tc := range testcases {
		err := Crawl(tc.nURLs, tc.nWorkers, tc.sPerURL, tc.cCap)
		if err != nil {
			t.Errorf("err: %v", err)
		}
	}
}
