package example3

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// response
var feed = `
<?xml version="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>Go workshop</title>
			<description>Golang</description>
		<link>http://www.gotest.net</link>
	</channel>
</rss> 
`

// mockServer returns a pointer to a server to handle the mock get call
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}
