package render

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func loadAssets(renderer *sdl.Renderer) (*sdl.Texture, error) {
	image, err := img.Load("assets/spritesheet_complete.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return nil, err
	}
	defer image.Free()

	spritesheet, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return nil, err
	}
	return spritesheet, nil
}
