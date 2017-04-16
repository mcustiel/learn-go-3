package render

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func loadAssets() (*sdl.Texture, error) {
	image, err := img.Load("assets/spritesheet_complete.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return err
	}
	defer image.Free()

	spritesheet, err =: display.renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return err
	}
	return spritesheet, nil
}
