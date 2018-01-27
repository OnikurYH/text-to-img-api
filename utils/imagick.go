package utils

import (
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

// ImagickDrawSetFont ...
func ImagickDrawSetFont(mw *imagick.MagickWand, dw *imagick.DrawingWand, font string, size float64, colour string) {
	dw.SetFont(font)

	pw := imagick.NewPixelWand()
	pw.SetColor(colour)
	dw.SetFillColor(pw)
	pw.Destroy()

	dw.SetFontSize(size)
}

// ImagickDrawMetrics ...
func ImagickDrawMetrics(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, dy *float64, sx float64, text string) {
	mw.AnnotateImage(dw, *dx, *dy, 0, text)
	mw.DrawImage(dw)
	fm := mw.QueryFontMetrics(dw, text)
	*dx += fm.TextWidth + sx
}

// ImagickGetHeightByIndex ...
func ImagickGetHeightByIndex(i int, fm *imagick.FontMetrics) float64 {
	h := fm.TextHeight
	if i == 0 {
		h = fm.CharacterHeight
	}
	return h + fm.Descender
}

// ImagickGetImageWidthHeightByText ...
func ImagickGetImageWidthHeightByText(mw *imagick.MagickWand, dw *imagick.DrawingWand, text string) (uint, uint) {
	texts := strings.Split(text, "\n")
	w := 0.0
	h := 0.0

	for i, t := range texts {
		fm := mw.QueryFontMetrics(dw, t)
		if fm.TextWidth > w {
			w = fm.TextWidth
		}
		h += ImagickGetHeightByIndex(i+1, fm)
	}

	return uint(w), uint(h)
}

// ImagickDrawText ...
func ImagickDrawText(mw *imagick.MagickWand, dw *imagick.DrawingWand, text string) {
	y := 0.0

	texts := strings.Split(text, "\n")
	for i, t := range texts {
		fm := mw.QueryFontMetrics(dw, t)
		y += ImagickGetHeightByIndex(i, fm)
		mw.AnnotateImage(dw, 0, y, 0, t)
		mw.DrawImage(dw)
	}
}
