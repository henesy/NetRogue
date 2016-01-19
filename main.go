package main

import (
    "fmt"
tb  "github.com/nsf/termbox-go"
//    "flag"
    "time"
    "encode/json"
)

//go:generate go get -u github.com/nsf/termbox-go
//go:generate go build -i -o encMaps encMaps/encoder.go
//go:generate go build -i -o encPlayer encPlayer/encoder.go
//go:generate go build -i -o NetRogue

/* index the type of sprite is on the map */
type sType int
const (
    PLAYER sType = iota
    STATIC
    ITEM
    CREEP
    NOTHING
)

/* current position in (x, y, z) */
type position struct {
	x int
	y int
    z int
}

/* stats and characteristics of a sprite or the player */
type statistics struct {
	hlth int
	atk  int
	dfs  int
	prof rune
}

/* bg/fg colors for use in sprites */
type colors struct {
    fg  tb.Attribute
    bg  tb.Attribute
}

/* basic sprite meta-struct */
type sprite struct {
	pos    position
	s      statistics
    desc   string
    t      sType //type of sprite
    c      colors
}


/* loads json-encoded maps from the maps folder hierarchy into a world array */
func genWorld() {


}

/* prints to pos x, y */
func tbPrint(x, y int, fg, bg tb.Attribute, msg string) {
    for _, c := range msg {
            tb.SetCell(x, y, c, fg, bg)
            x++
    }
}

/* draws a box given two positions */
func tbDrawBox(tl, br position, fg, bg tb.Attribute, t string) {

}

/* draws the screen (continual loop) */
func draw(w, h int) {
    defer tb.Flush()
    for {


        tb.Flush()
        time.Sleep(20 * time.Millisecond)
    }
}

/* check checks the error err for an error and crashes the program if != nil */
func check(err error) {
    if err != nil {
        termbox.Close()
        fmt.Printf("%v\n", err)
        os.Exit(1)
    }
}


/* A multiplayer rogue-like in the spirit of TOAG, but implemented in termbox-go */
func main() {


    //flag.Parse()
    /* start network handling and channels */
    max_dimensions := 80*24
    screen := make([]sprite, max_dimensions)
    world := genWorld()

    /* initialize termbox interface */
    err := tb.Init()
    check(err)

    tb.SetInputMode(tb.InputAlt)
    w, h := tb.Size()

    tb.Clear(tb.ColorBlack, tb.ColorBlack)
    tb.Flush()
    go draw(w, h)

    for run := true;run == true; {
        switch ev := tb.PollEvent(); ev.Type {
            case tb.EventKey:
                //key := ev.Ch
                if ev.Key == tb.KeyCtrlQ {
                    tb.Flush()
                    run = false
                }
            default:
        }
    }
}
