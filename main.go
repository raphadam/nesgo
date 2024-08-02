package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/raphadam/gelly"
	"github.com/raphadam/nesgo/nes"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	_ "embed"
)

//go:embed ui.png
var UI_BYTES []byte
var UI_IMG *ebiten.Image

func main() {
	gelly.Run(gelly.Scene{
		Name:   "Nes game",
		Layers: []gelly.Layer{&Game{}},
	})
}

type Game struct {
	console      nes.Console
	tex          *ebiten.Image
	face         font.Face
	smallFace    font.Face
	instructions [300]nes.Instruction
}

func (g *Game) Init(c *gelly.Client) {
	ebiten.SetRunnableOnUnfocused(false)
	ebiten.SetWindowSize(1280, 720)
	c.SetLayoutSize(1280, 720)

	// ui img
	img, _, err := image.Decode(bytes.NewBuffer(UI_BYTES))
	if err != nil {
		log.Fatal("unable to load image ui")
	}
	UI_IMG = ebiten.NewImageFromImage(img)

	// ttf
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal("unable to parse font", err)
	}

	g.face = truetype.NewFace(ttfFont, &truetype.Options{Size: 32})
	g.smallFace = truetype.NewFace(ttfFont, &truetype.Options{Size: 24})
	g.tex = ebiten.NewImage(32, 32)

	g.Reset()
}

func (g *Game) Reset() {
	// init the cpu
	cpu := &nes.Cpu{}
	cpu.StackPointer = 0xFF
	cpu.ProgramCounter = 0x0600
	cpu.Status |= (nes.Break | nes.Break2)

	g.console = nes.Console{
		Cpu:       cpu,
		Cartridge: nes.Cartridge{},
	}

	copy(g.console.Cpu.Ram[0x0600:], gameCode[:])
	g.console.Cpu.ProgramCounter = 0x0600
}

func (g *Game) Message(c *gelly.Client, msg gelly.Message) bool {
	return false
}

func (g *Game) Update(c *gelly.Client, dt time.Duration) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		c.Close()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.console.Write(0xff, 0x77)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.console.Write(0xff, 0x73)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.console.Write(0xff, 0x61)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.console.Write(0xff, 0x64)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
	}

	g.console.Write(0xfe, uint8(rand.Intn(15)+1))

	// update
	for i := 0; i < 150; i++ {
		g.instructions[i] = g.console.Cpu.ExecuteInstruction(g.console)
	}
}

func (g *Game) Draw(r *ebiten.Image) {
	x := 0
	y := 0

	for i := 0x0200; i < 0x0600; i++ {
		p := g.console.Read(uint16(i))
		if x >= 32 {
			x = 0
			y++
		}
		switch p {
		case 0:
			g.tex.Set(x, y, color.RGBA{R: 10, G: 10, B: 10, A: 255})
		case 1:
			g.tex.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
		case 2, 9:
			g.tex.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
		case 3, 10:
			g.tex.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		case 4, 11:
			g.tex.Set(x, y, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		case 5, 12:
			g.tex.Set(x, y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
		case 6, 13:
			g.tex.Set(x, y, color.RGBA{R: 255, G: 0, B: 255, A: 255})
		case 7, 14:
			g.tex.Set(x, y, color.RGBA{R: 255, G: 255, B: 0, A: 255})
		default:
			g.tex.Set(x, y, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		}

		x++
	}

	// draw game
	op := ebiten.GeoM{}
	op.Scale(22, 22)
	// op.Translate(100, 100)
	r.DrawImage(g.tex, &ebiten.DrawImageOptions{GeoM: op})
	r.DrawImage(UI_IMG, &ebiten.DrawImageOptions{})

	// draw cpu state
	text.Draw(r, fmt.Sprintf("%03d", g.console.Cpu.Accumulator), g.face, 780, 415, gelly.ColorWhite)
	text.Draw(r, fmt.Sprintf("%03d", g.console.Cpu.IndirectX), g.face, 920, 415, gelly.ColorWhite)
	text.Draw(r, fmt.Sprintf("%03d", g.console.Cpu.IndirectY), g.face, 1060, 415, gelly.ColorWhite)
	text.Draw(r, fmt.Sprintf("%03d", g.console.Cpu.Status), g.face, 1200, 415, gelly.ColorWhite)

	for i := range 8 {
		text.Draw(r, g.instructions[i].Name, g.smallFace, 730, 490+(i*28), gelly.ColorWhite)
	}

	// text.Draw(r, fmt.Sprintf("PC: %d", g.console.Cpu.ProgramCounter), g.face, 800, 160, gelly.ColorGreen)
	// text.Draw(r, fmt.Sprintf("SP: %d", g.console.Cpu.StackPointer), g.face, 800, 180, gelly.ColorGreen)
	// text.Draw(r, "RAM:", g.face, 800, 220, gelly.ColorGreen)

	y = 0
	x = 0
	for i := 0; i < 48; i += 4 {
		if x%6 == 0 {
			y++
			x = 0
		}
		x++

		text.Draw(r, fmt.Sprintf(
			"%02x %02x %02x %02x",
			g.console.Cpu.Ram[i],
			g.console.Cpu.Ram[i+1],
			g.console.Cpu.Ram[i+2],
			g.console.Cpu.Ram[i+3],
		), g.face, 675+(y*200), 450+(x*40), gelly.ColorWhite)
	}
}

func (g *Game) Dispose(c *gelly.Client) {
}

var gameCode = [...]uint8{
	0x20, 0x06, 0x06, 0x20, 0x38, 0x06, 0x20, 0x0d, 0x06, 0x20, 0x2a, 0x06, 0x60, 0xa9, 0x02, 0x85,
	0x02, 0xa9, 0x04, 0x85, 0x03, 0xa9, 0x11, 0x85, 0x10, 0xa9, 0x10, 0x85, 0x12, 0xa9, 0x0f, 0x85,
	0x14, 0xa9, 0x04, 0x85, 0x11, 0x85, 0x13, 0x85, 0x15, 0x60, 0xa5, 0xfe, 0x85, 0x00, 0xa5, 0xfe,
	0x29, 0x03, 0x18, 0x69, 0x02, 0x85, 0x01, 0x60, 0x20, 0x4d, 0x06, 0x20, 0x8d, 0x06, 0x20, 0xc3,
	0x06, 0x20, 0x19, 0x07, 0x20, 0x20, 0x07, 0x20, 0x2d, 0x07, 0x4c, 0x38, 0x06, 0xa5, 0xff, 0xc9,
	0x77, 0xf0, 0x0d, 0xc9, 0x64, 0xf0, 0x14, 0xc9, 0x73, 0xf0, 0x1b, 0xc9, 0x61, 0xf0, 0x22, 0x60,
	0xa9, 0x04, 0x24, 0x02, 0xd0, 0x26, 0xa9, 0x01, 0x85, 0x02, 0x60, 0xa9, 0x08, 0x24, 0x02, 0xd0,
	0x1b, 0xa9, 0x02, 0x85, 0x02, 0x60, 0xa9, 0x01, 0x24, 0x02, 0xd0, 0x10, 0xa9, 0x04, 0x85, 0x02,
	0x60, 0xa9, 0x02, 0x24, 0x02, 0xd0, 0x05, 0xa9, 0x08, 0x85, 0x02, 0x60, 0x60, 0x20, 0x94, 0x06,
	0x20, 0xa8, 0x06, 0x60, 0xa5, 0x00, 0xc5, 0x10, 0xd0, 0x0d, 0xa5, 0x01, 0xc5, 0x11, 0xd0, 0x07,
	0xe6, 0x03, 0xe6, 0x03, 0x20, 0x2a, 0x06, 0x60, 0xa2, 0x02, 0xb5, 0x10, 0xc5, 0x10, 0xd0, 0x06,
	0xb5, 0x11, 0xc5, 0x11, 0xf0, 0x09, 0xe8, 0xe8, 0xe4, 0x03, 0xf0, 0x06, 0x4c, 0xaa, 0x06, 0x4c,
	0x35, 0x07, 0x60, 0xa6, 0x03, 0xca, 0x8a, 0xb5, 0x10, 0x95, 0x12, 0xca, 0x10, 0xf9, 0xa5, 0x02,
	0x4a, 0xb0, 0x09, 0x4a, 0xb0, 0x19, 0x4a, 0xb0, 0x1f, 0x4a, 0xb0, 0x2f, 0xa5, 0x10, 0x38, 0xe9,
	0x20, 0x85, 0x10, 0x90, 0x01, 0x60, 0xc6, 0x11, 0xa9, 0x01, 0xc5, 0x11, 0xf0, 0x28, 0x60, 0xe6,
	0x10, 0xa9, 0x1f, 0x24, 0x10, 0xf0, 0x1f, 0x60, 0xa5, 0x10, 0x18, 0x69, 0x20, 0x85, 0x10, 0xb0,
	0x01, 0x60, 0xe6, 0x11, 0xa9, 0x06, 0xc5, 0x11, 0xf0, 0x0c, 0x60, 0xc6, 0x10, 0xa5, 0x10, 0x29,
	0x1f, 0xc9, 0x1f, 0xf0, 0x01, 0x60, 0x4c, 0x35, 0x07, 0xa0, 0x00, 0xa5, 0xfe, 0x91, 0x00, 0x60,
	0xa6, 0x03, 0xa9, 0x00, 0x81, 0x10, 0xa2, 0x00, 0xa9, 0x01, 0x81, 0x10, 0x60, 0xa2, 0x00, 0xea,
	0xea, 0xca, 0xd0, 0xfb, 0x60,
}
