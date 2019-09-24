/*
	Bot per il gioco di telegram Lumberjack,
	dipendenze (
		"github.com/go-vgo/robotgo"
		"github.com/vova616/screenshot"
	)
*/

package main

import (
	"image"
	"image/color"

	"github.com/go-vgo/robotgo"

	"github.com/vova616/screenshot"
)

/*
	I rettangoli sono presi da una finestra di chrome in finestra massimizzata, lo schermo ha una risoluzione di 1920x1080
	Più il rettangolo è grande più si compensa al ritardo della pressione del tasto virtuale
	Se il rettangolo è troppo grande si compensa troppo e si rischia di cambiare lato troppo presto
	Se il rettangolo è troppo piccolo si rischia di non prevedere in tempo il ramo
*/

//Test diminuire in base ai millisecondi
//Partenza 600 Y, 100 ms
//550 Y, 50 ms
//520 Y, 30 ms fallisce
//502 Y, 30 ms poco affidabile
//503 Y, 30 ms più affidabile
//500 Y, 25 ms fallisce
var left = image.Rectangle{image.Point{850, 503 /* Y */}, image.Point{900, 687}}
var right = image.Rectangle{image.Point{990, 503 /* Y */}, image.Point{1040, 687}}

//Colore del ramo
var branch = color.RGBA{161, 116, 56, 255}

func main() {
	orientation := left
	key := "left"
	wood := false

	for {
		//Catturo uno dei rettangoli del gioco a seconda del tasto premuto
		screen, _ := screenshot.CaptureRect(orientation)
		imgBounds := screen.Bounds()
		width := imgBounds.Max.X
		height := imgBounds.Max.Y
		//Estraggo tutti i colori dei pixel del rettangolo e controllo se corrisponde al colore del ramo
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				r, g, b, a := screen.At(x, y).RGBA()
				imageColor := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
				if imageColor == branch {
					wood = true
					break
				}
			}
		}
		//Se il ramo è stato trovato cambio posizione e il rettangolo di acquisizione
		if wood {
			//fmt.Println("WOOD")
			if orientation == left {
				//fmt.Println("Switch to right")
				key = "right"
				orientation = right
			} else {
				//fmt.Println("Switch to left")
				key = "left"
				orientation = left
			}
			wood = false
		}
		robotgo.KeyTap(key)
		robotgo.MilliSleep(30)
	}
}
