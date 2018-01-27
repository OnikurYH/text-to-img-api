package api

import (
	"net/http"
	"strconv"
	"strings"

	"onikur.com/text-to-img-api/utils"

	"gopkg.in/gographics/imagick.v3/imagick"
)

// Controller ...
type Controller struct {
	Mux *http.ServeMux
}

// Options ...
type Options struct {
	FontSize    int
	LineMaxChar int
}

// NewOptions ...
func NewOptions() Options {
	opts := Options{
		FontSize:    12,
		LineMaxChar: 30,
	}
	return opts
}

func (ctrl *Controller) index(res http.ResponseWriter, req *http.Request) {
	text := strings.TrimPrefix(req.URL.Path, "/-/")

	// opts := NewOptions()

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	mw.ReadImage("xc:transparent")

	utils.ImagickDrawSetFont(mw, dw, "Arial", 48, "#333333")
	sw, sh := utils.ImagickGetImageWidthHeightByText(mw, dw, text)

	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_TRANSPARENT)
	mw.SetSize(sw, sh)
	mw.ReadImage("xc:transparent")

	utils.ImagickDrawText(mw, dw, text)
	mw.SetImageFormat("png")
	b := mw.GetImageBlob()

	res.Header().Set("Content-Length", strconv.Itoa(len(b)))
	res.Header().Set("Content-Type", "image/png")
	res.Write(b)
}

// Init ...
func (ctrl *Controller) Init() {
	ctrl.Mux.HandleFunc("/-/", ctrl.index)
}
