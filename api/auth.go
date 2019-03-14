package api

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var token Token

func serverManager(srv *http.Server, keepAlive chan bool) {
	for {
		select {
		case <-keepAlive:
			ctx := context.Background()

			err := srv.Shutdown(ctx)

			if err != nil {
				fmt.Printf("Failed to shutdown temporary http server, this should have no effect on your ability to complete this process.  It's running on port 15298.")
			}

			return
		default:
		}
	}
}

func GetCode(id string) (c string) {
	m := GetAuthorizationURL(id)
	fmt.Printf("\nNavigate to the following URL to Authorize Spotistats:\n%s\n", m)
	keepAlive := make(chan bool)
	server := &http.Server{Addr: ":8000"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		keepAlive <- true
		fmt.Fprintf(w, "<h1>Almost done!</h1><p>Spotistats has been approved, just copy the following code back to the CLI:\n <span style=\"color: #FF0000\">%s</span></p>", code)
	})

	go serverManager(server, keepAlive)
	server.ListenAndServe()

	fmt.Print("\nEnter Code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c = strings.TrimSpace(scanner.Text())

	return c
}

func Authenticate(id, secret string) {
	code := GetCode(id)
	token = CodeAuthorize(id, secret, code)
	fmt.Println("\nSuccessful Authentication")
}
