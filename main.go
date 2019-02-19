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

	img1Rect := sdl.Rect{0, 0, 60, 60}
	img2Rect := sdl.Rect{60, 0, 60, 60}
	newImgRect := sdl.Rect{0, 0, 120, 120}

	newImg, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, 120, 120)
	if err != nil {
		panic(err)
	}
	defer newImg.Destroy()

	renderer.SetRenderTarget(newImg)
	renderer.Copy(img1, nil, &img1Rect)
	renderer.Copy(img2, nil, &img2Rect)

	renderer.SetRenderTarget(nil)

	for running := true; running != false; {
		event := sdl.PollEvent()
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}

		renderer.Copy(newImg, nil, &newImgRect)
		renderer.Present()
		renderer.Clear()
	}
}
