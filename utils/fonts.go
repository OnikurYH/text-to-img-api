package utils

import (
	"sync"

	"onikur.com/text-to-img-api/conf"

	"gopkg.in/gographics/imagick.v3/imagick"
)

var once sync.Once

// FontsObject ...
type FontsObject struct {
	Available []string
	Disabled  []string
}

// CacheFonts ...
func (fo *FontsObject) CacheFonts() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	fonts := mw.QueryFonts("*")
	fo.Available = make([]string, 0)
	fo.Disabled = make([]string, 0)
	for _, font := range fonts {
		if enabled, ok := conf.Get().Font.Include[font]; ok && enabled {
			fo.Available = append(fo.Available, font)
		} else {
			fo.Disabled = append(fo.Disabled, font)
		}
	}
}

var fontsInstance *FontsObject

// Fonts ...
func Fonts() *FontsObject {
	once.Do(func() {
		fontsInstance = &FontsObject{}
	})
	return fontsInstance
}
