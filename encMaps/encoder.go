package main

import (
    "fmt"
    "encoding/json"
//tb  "github.com/nsf/termbox-go"
nr  "netrogue"
    "os"
    "time"
)

/* all custom types and sprites must be synchronized/matched with
the primary main.go server/client */

type Coordinates struct {
    X   int
    Y   int
    Z   int
}

var world [10][10][10][23][80]nr.Sprite
var zTenCount int
var xTenCount int
var y, x, z int
var orderNum int

/* check checks the error err for an error and crashes the program if != nil */
func check(err error) {
    if err != nil {
        termbox.Close()
        fmt.Printf("%v\n", err)
        //os.Exit(1)
    }
}

/* calls parseMap() and handles each transition */
func mapHandler(mapChan chan int, mapDir os.FileInfo, num int) {
        for {
            if orderNum == num {
                f, err := os.Open(mapDir.Name())
                check(err)
                maps, err := f.Readdirnames()
                check(err)
                //z
                for k := 0;k < 10;k++ {
                    c := Coordinates{x, y, z}
                    go parseMap(mapChan, c)
                    z++
                }

                //advance the position of the coordinates (if we can...)
                z = 0
                if y+1 > 10 {
                    y = 0
                    if z+1 > 10 {
                        z = 0
                    } else {
                        z++
                    }
                } else {
                    y++
                }

                //this directory is done, proceed to next directory, advance order number
                orderNum++
                break
            } else {
                time.Sleep(1 * time.Millisecond)
            }
        }
}

/* parses a map into a nr.sprite to `co` coordinates in world[] array */
func parseMap(mapChan chan int, tMap os.FileInfo, co coordinates) {
    /*
        each map will be the "height" in the 3D array for world
        each map will be a 2D array of sprites

    */


    mapChan <- 1
}


/* Encodes map files into sprite arrays with colors && descriptions for exporting */
func main() {
        /* This program should only be run from the base NR directory
            as such, check if we are in the base NR directory */
        y, x, z = 0, 0, 0
        orderNum = 0

        curDir, err := os.Open("./")
        check(err)
        dirList, err := curDir.Readdirnames()
        dirsToCheck := map[string]bool{"colors": false, "net-rogue": false, "world": false, "encMaps": false}
        for _, d := range dirList {
            for k, _ := range dirsToCheck {
                if d == k {
                    dirsToCheck[k] = true
                }
            }
        }
        numFound := 0
        for _, v := range dirsToCheck {
            if v == true {
                numFound++
            }
        }
        if numFound >= 2 {
            fmt.Println("Confirmed we're in the base directory, continuing...")
        } else {
            fmt.Println("Warning: This program is intended to run from the base NR directory!")
            fmt.Print("Continue? [y/n]: ")
            rsps := ""
            fmt.Scanln(&rsps)
            if rsps != "y" || rsps != "Y" {
                os.Exit(1)
            }
        }
        /* END initial directory check */
        //move to opening world/ then this!
        worldD := os.Open("world")
        //world := make([]nr.Sprite, 1840000, 2000000)
        dirs, err:= worldD.Readdirnames()
        check(err)


        mapChan := make(chan int, 1)
        descend := func(newdLst []os.FileInfo) {
            numDirs := 0
            for _, v := range newLst {
                if v.IsDir() {
                    numDirs++
                    f, err := os.Open(v.Name())
                    check(err)
                    nnLst, err := f.Readdir()
                    check(err)
                    go descend(nnLst) //if there are more dirs
                } else {
                    //fork folder for parsing
                    go mapHandler(mapChan, newDLst, orderNum)
                    orderNum = 1
                    break
                }
            }

        }
        go descend(dirs)

        //wait for all maps to be registered
        rplyCnt := 0
        for rplyCnt < 1000 {
            select {
                r := <- mapChan:
                    rplyCnt++
                default:
                    time.Sleep(1 * time.Millisecond)
            }
        }

        //here is where we encode the world[] array...


}
