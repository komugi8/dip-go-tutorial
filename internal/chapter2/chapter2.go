package chapter2

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("http://mock:80/users")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	q := u.Query()
	q.Set("age", "25")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("key", "dip")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Post(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("http://mock:80/users")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	form := &url.Values{}
	form.Set("name", "dip 次郎")
	form.Set("age", "24")

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", "dip")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
