package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth, winHeight int32  = 800, 600
	img1Filename        string = "C:/Go/src/github.com/MJPB/Tests/grass-tiles-2-small.png"
	img2Filename        string = "C:/Go/src/github.com/MJPB/Tests/VW62v2X.jpg"
	tileSize            int32  = 64
)

func createBackground(renderer *sdl.Renderer) (*sdl.Texture, sdl.Rect) {

	img1, err := img.LoadTexture(renderer, img1Filename)
	if err != nil {
		panic(err)
	}
	defer img1.Destroy()

	img2, err := img.LoadTexture(renderer, img2Filename)
	if err != nil {
		panic(err)
	}
	defer img2.Destroy()

	newImg, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 120, 120)
	if err != nil {
		panic(err)
	}

	img1Rect := sdl.Rect{0, 0, tileSize, tileSize}
	img2Rect := sdl.Rect{tileSize, 0, tileSize, tileSize}
	newTexRect := sdl.Rect{0, 0, winWidth, winHeight}

	renderer.SetRenderTarget(newImg)

	renderer.Copy(img1, &img1Rect, &img1Rect)
	renderer.Copy(img2, nil, &img2Rect)

	renderer.SetRenderTarget(nil)

	return newImg, newTexRect
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("My Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	background, backgroundRect := createBackground(renderer)
	defer background.Destroy()

	for running := true; running != false; {
		event := sdl.PollEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}

		renderer.Copy(background, nil, &backgroundRect)
		renderer.Present()
		renderer.Clear()
	}
}
