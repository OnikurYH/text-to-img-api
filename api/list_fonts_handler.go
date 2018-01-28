package api

import (
	"encoding/json"
	"net/http"

	"onikur.com/text-to-img-api/utils"
)

// ListFontsHandler ...
type ListFontsHandler struct{}

func (h *ListFontsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var a []string
	if status := req.URL.Query().Get("status"); status == "disabled" {
		a = utils.Fonts().Disabled
	} else {
		a = utils.Fonts().Available
	}

	js, err := json.Marshal(a)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}
