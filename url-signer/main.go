package main

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var KEY string

const URL_DOMAIN = "cdn.alananetwork.private"
const TTL = 3600 //seconds

func getUrlSigned(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST: /generate received from:", r.Header.Get("X-FORWARDED-FOR"))

	fileRequested := r.Header.Get("X-File")
	if len(fileRequested) == 0 {
		fmt.Println(fmt.Errorf("ERROR: Bad X-File header"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("File requested:", fileRequested)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("%s", generateSignedUrl(fileRequested)))
}

func generateSignedUrl(filename string) string {
	urlPath := "/file/" + filename
	var expiration int = int(time.Now().UnixMilli() + TTL)
	var link string = KEY + "-" + urlPath + "-" + strconv.Itoa(expiration)

	hash := md5.Sum([]byte(link))
	hashBase64 := base64.RawURLEncoding.EncodeToString(hash[:])
	link = fmt.Sprintf("https://%s%s?md5=%s&expires=%s", URL_DOMAIN, urlPath, hashBase64, strconv.Itoa(expiration))
	return fmt.Sprintf("%s", link)
}

func main() {
	http.HandleFunc("/generate", getUrlSigned)
	fmt.Println("Started Server...")

	err := http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("ERROR: Failed starting server: %s\n", err)
		os.Exit(1)
	}
}
