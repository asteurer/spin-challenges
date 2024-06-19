package main

import (
	"bytes"
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

type Response struct {
	RequestBody   string `json:"requestBody"`
	DecryptedBody string `json:"decryptedBody"`
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		// files, err := os.ReadDir("/")
		// for _, file := range files {
		// 	fmt.Println(file.Name(), file.IsDir())
		// }

		// messageBytes, err := os.ReadFile("message.txt")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		messageBytes := []byte("TESTING")

		req, err := http.NewRequest("GET", "http://localhost:3000/crypto", bytes.NewBuffer(messageBytes))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header.Set("x-action", "encrypt")

		resp, err := spinhttp.Send(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.Body.Close()

		w.Header().Set("content-type", "application/json")
		w.Write(bodyBytes)
	})
}

func main() {}
