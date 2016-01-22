package netrogue

import (
//	"fmt"
//    "encoding/json"
tb  "github.com/nsf/termbox-go"
)

/* Sprites and functions and methods and all that go here, mass find -> replace
goes into the [main/encoder]*.go files, fun. */

/* index the type of Sprite is on the map */
type Stype int
const (
    PLAYER Stype = iota
    STATIC
    ITEM
    CREEP
    NOTHING
)

/* current Position in (x, y, z) */
type Position struct {
	X int
	Y int
    Z int
}

/* stats and characteristics of a Sprite or the player */
type Statistics struct {
	Hlth int
	Atk  int
	Dfs  int
	Prof rune
}

/* bg/fg colors for use in Sprites */
type Colors struct {
    Fg  tb.Attribute
    Fg  tb.Attribute
}

/* basic Sprite meta-struct */
type Sprite struct {
	Pos    Position
	S      statistics
    Desc   string
    T      sType //type of Sprite
    C      colors
}

//var World [10][10][10][23][80]Sprite


/* prints to pos x, y */
func Print(x, y int, fg, bg tb.Attribute, msg string) {
    for _, c := range msg {
            tb.SetCell(x, y, c, fg, bg)
            x++
    }
}

/* draws a box given two nr.Positions */
func DrawBox(tl, br nr.Position, fg, bg tb.Attribute, t string) {
	var ul, bl, ur, br, side, topbot string
	if t == "thick" {
		ul, bl, ur, br, side, topbot = "╔", "╚", "╗", "╝", "║", "═"
	} else if t == "thin" {
		ul, bl, ur, br, side, topbot = "┌", "", "┐", "", "│", "─"
	} else {
		ul, bl, ur, br, side, topbot = "*", "*", "*", "*", "*", "*"
	}
	//sides
	for i := tl.Y+1;i < br.Y-1;i++ {
		tbPrint(tl.X, i, tb.ColorWhite, tb.ColorBlack, side)
		tbPrint(br.X, i, tb.ColorWhite, tb.ColorBlack, side)
	}
	//bottom/top
	for i := tl.X+1;i < br.X-1;i++ {
		tbPrint(i, tl.Y, tb.ColorWhite, tb.ColorBlack, topbot)
		tbPrint(i, br.Y, tb.ColorWhite, tb.ColorBlack, topbot)
	}
	//corner
	tbPrint(tl.X, tl.Y, tb.ColorWhite, tb.ColorBlack, ul)
	tbPrint(tl.X, br.Y, tb.ColorWhite, tb.ColorBlack, bl)
	tbPrint(br.X, tl.Y, tb.ColorWhite, tb.ColorBlack, ur)
	tbPrint(br.X, br.Y, tb.ColorWhite, tb.ColorBlack, br)
}
