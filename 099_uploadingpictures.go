package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("99file.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	// process form submission
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]         //gives extenstion of the file
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext   //using hash to create file name with same old extension

		// create new file
		wd, err := os.Getwd()            //get the current working directory
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)          //path where the file needs to be created
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// copy
		mf.Seek(0, 0)            //go to start of the multipart file
		io.Copy(nf, mf)

		// add filename to this user's cookie
		c = appendValue(w, c, fname)

	}
	xs := strings.Split(c.Value, "|")
	t.ExecuteTemplate(w, "99file.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// takes in a file name now also
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}