package main

import (
    "fmt"
tb  "github.com/nsf/termbox-go"
//    "flag"
    "time"
    "encode/json"
nr  "netrogue"
)

//go:generate go get -u github.com/nsf/termbox-go
//go:generate cd net-rogue && ./install_netrogue_lib.sh
//go:generate go build -i -o encMaps encMaps/encoder.go
//go:generate go build -i -o encPlayer encPlayer/encoder.go
//go:generate go build -i -o NetRogue

/* loads json-encoded maps from the maps folder hierarchy into a world array */
func loadWorld() {


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
    screen := make([]nr.Sprite, max_dimensions)
    world := loadWorld()

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
