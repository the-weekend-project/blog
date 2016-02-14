package api

import (
	"net/http"

	"github.com/the-weekend-project/blogApi"
)

func init() {
	http.Handle("/", blogApi.GetRouter())
}
