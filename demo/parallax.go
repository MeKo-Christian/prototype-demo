package demo

import (
	"math"
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

var (
	bgStartTime = time.Now()
)

type star struct {
	x, y      int
	baseColor draw.Color
	phase     float64
	speed     float64
}

var stars []star
var starsInitialized bool

func updateParallax(w draw.Window) {
	if !starsInitialized {
		initStars(w)
		starsInitialized = true
	}

	t := time.Since(bgStartTime).Seconds()
	_, height := w.Size()

	// Draw static star layer
	drawStarfield(w)

	// Define and draw scrolling layers
	layers := []struct {
		path  string
		speed float64
		y     int
	}{
		{"assets/bg_layer1.png", 0, 0},
		{"assets/bg_layer2.png", 20, 0},
		{"assets/bg_layer3.png", 40, height - 500},
		{"assets/bg_layer4.png", 60, height - 512},
	}

	for idx, layer := range layers {
		if idx == 1 {
			drawStarfield(w)
		}
		drawParallaxLayer(w, layer.path, t, layer.speed, layer.y)
	}
}

func drawParallaxLayer(w draw.Window, path string, t float64, speed float64, y int) {
	imgW, imgH, err := w.ImageSize(path)
	if err != nil || imgH == 0 {
		return
	}

	width, _ := w.Size()
	offset := int(t*speed) % imgW

	// Draw enough copies to fill screen horizontally
	for x := -offset; x < width; x += imgW {
		_ = w.DrawImageFile(path, x, y)
	}
}

func initStars(w draw.Window) {
	width, _ := w.Size()
	yMin := 0
	yMax := 200

	for i := 0; i < 20; i++ {
		x := rand.Intn(width)
		y := yMin + rand.Intn(yMax-yMin)
		gray := 0.5 + rand.Float32()*0.5

		stars = append(stars, star{
			x: x,
			y: y,
			baseColor: draw.Color{
				R: gray, G: gray, B: gray, A: 1,
			},
			phase: rand.Float64() * 2 * math.Pi,
			speed: 1 + rand.Float64()*2, // between 1–3 Hz
		})
	}
}

func drawStarfield(w draw.Window) {
	t := time.Since(bgStartTime).Seconds()

	for _, s := range stars {
		brightness := 0.7 + 0.3*math.Sin(t*s.speed+s.phase)
		c := s.baseColor
		c.R *= float32(brightness)
		c.G *= float32(brightness)
		c.B *= float32(0.2 + 0.8*brightness)
		w.FillRect(s.x, s.y, 3, 3, c)
	}
}
