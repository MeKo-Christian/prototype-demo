package demo

import (
	"math"
	"time"

	"github.com/gonutz/prototype/draw"
)

var (
	scrollText = "   WELCOME TO THE PROTOTYPE DEMO SCROLLER! "
	startTime  = time.Now()
)

func updateScroller(w draw.Window) {
	t := time.Since(startTime).Seconds()
	screenW, screenH := w.Size()

	scale := 3.0
	speed := 100.0 // pixels per second
	amp := 80.0
	freq := 0.13

	charW := 12.0 * scale
	textW := int(float64(len(scrollText)) * charW)
	y := int(float64(screenH)/8 + amp/2)

	// Update scroll position
	offsetX := int(t * speed)
	x := screenW - offsetX

	// Draw each character in order
	for i, ch := range scrollText {
		charX := x + int(float64(i)*charW)
		if charX+int(charW) < 0 {
			continue
		}
		if charX > screenW {
			continue
		}
		yOffset := int(amp * math.Sin(float64(i)*freq+t*2))
		w.DrawScaledText(string(ch), charX, y+yOffset, float32(scale), draw.Cyan)
	}

	// Reset the scroll only when last char is offscreen
	if x+textW < 0 {
		startTime = time.Now()
	}
}
