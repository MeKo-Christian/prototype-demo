package demo

import "github.com/gonutz/prototype/draw"

var (
	soundInitialized bool
	soundPath        = "assets/prototype-demo.wav"
)

func Update(w draw.Window) {
	if !soundInitialized {
		w.PlaySoundFile(soundPath)
		soundInitialized = true
	}
	updateParallax(w)
	updateScroller(w)
}
