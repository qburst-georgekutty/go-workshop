package example1

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	// test cases
	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://httpbin.org/get", http.StatusOK},
		{"https://qburst-georgekutty.github.io/go-workshop/new/", http.StatusNotFound},
	}
	t.Log("Given the need to test downloading different context.")
	{
		// run each test cases
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWhen checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				res, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("Should be able to make the Get call : %v", err)
				}
				defer res.Body.Close()

				t.Logf("Should be able to make the Get call.")

				if res.StatusCode == tt.statusCode {
					t.Logf("Should receive a %d status code.", tt.statusCode)
				} else {
					t.Errorf("Should receive a %d status code : %v", tt.statusCode, res.StatusCode)
				}
			}
		}
	}
}
