package example2

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	downloadURL := "https://httpbin.org/get"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", downloadURL, statusCode)
		{
			resp, err := http.Get(downloadURL)
			if err != nil {
				t.Fatalf("Should be able to make Get call: %v", err)
			}
			defer resp.Body.Close()

			t.Logf("Should be able to make the Get call.")

			if statusCode == resp.StatusCode {
				t.Logf("Should receive a %d status code.", statusCode)
			} else {
				t.Errorf("Should receive a %d status code : %d", statusCode, resp.StatusCode)
			}
		}
	}
}

/* output

=== RUN   TestDownload
--- PASS: TestDownload (1.91s)
	download_test.go:17: Given the need to test downloading content.
	download_test.go:19: 	Test 0:	When checking "https://httpbin.org/get" for status code 200
	download_test.go:27: 	✓	Should be able to make the Get call.
	download_test.go:30: 	✓	Should receive a 200 status code.
*/
