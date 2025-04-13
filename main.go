package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/meko-christian/prototype-demo/demo"
)

func main() {
	draw.RunWindow("protptype demo", 800, 600, demo.Update)
}
