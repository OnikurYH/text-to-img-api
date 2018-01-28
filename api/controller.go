package api

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"

	"onikur.com/text-to-img-api/conf"

	"onikur.com/text-to-img-api/utils"

	"gopkg.in/gographics/imagick.v3/imagick"
)

var extansion = "png"

// Controller ...
type Controller struct {
	Mux *http.ServeMux
}

// Options ...
type Options struct {
	Font        string
	FontSize    float64
	FontColor   string
	LineMaxChar int
}

// NewOptions ...
func NewOptions() Options {
	opts := Options{
		Font:        conf.Get().Font.Defaults.Font,
		FontSize:    conf.Get().Font.Defaults.FontSize,
		FontColor:   conf.Get().Font.Defaults.FontColor,
		LineMaxChar: 0,
	}
	return opts
}

func (ctrl *Controller) index(res http.ResponseWriter, req *http.Request) {
	if !strings.HasSuffix(req.URL.Path, "."+extansion) {
		http.Redirect(res, req, req.URL.Path+"."+extansion+req.URL.Query().Encode(), 301)
		return
	}

	text := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(req.URL.Path, "/-/"), "/api/text/"), extansion)

	// Options
	qs := req.URL.Query()
	opts := NewOptions()
	if v, err := strconv.ParseFloat(qs.Get("fsize"), 64); err == nil {
		opts.FontSize = math.Max(conf.Get().Font.Defaults.MinFontSize, math.Min(v, conf.Get().Font.Defaults.MaxFontSize))
	}
	if qs.Get("fcolor") != "" {
		opts.FontColor = qs.Get("fcolor")
	}
	if qs.Get("f") != "" {
		if e, ok := conf.Get().Font.Include[qs.Get("f")]; e && ok {
			opts.Font = qs.Get("f")
		}
	}

	// Split to multi-line
	if opts.LineMaxChar > 0 {
		text = utils.StringsInsertRuneStep(text, opts.LineMaxChar, "\n")
	}

	// Start drawing
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	mw.ReadImage("xc:transparent")
	utils.ImagickDrawSetFont(mw, dw, opts.Font, opts.FontSize, "#"+opts.FontColor)

	sw, sh := utils.ImagickGetImageWidthHeightByText(mw, dw, text)

	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_TRANSPARENT)
	mw.SetSize(sw, sh)
	mw.ReadImage("xc:transparent")

	utils.ImagickDrawText(mw, dw, text)
	mw.SetImageFormat(extansion)
	b := mw.GetImageBlob()

	res.Header().Set("Content-Length", strconv.Itoa(len(b)))
	res.Header().Set("Content-Type", "image/"+extansion)
	res.Write(b)
}

func (ctrl *Controller) listAvailableFonts(res http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(utils.Fonts().Available)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func (ctrl *Controller) listFonts(res http.ResponseWriter, req *http.Request) {
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

// Init ...
func (ctrl *Controller) Init() {
	ctrl.Mux.HandleFunc("/-/", ctrl.index)
	ctrl.Mux.HandleFunc("/api/text/", ctrl.index)
	ctrl.Mux.HandleFunc("/api/fonts", ctrl.listFonts)
}
