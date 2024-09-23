package main

import (
	"bufio"
	"fmt"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxRadiusDigits = 4
const ballFlightStepsNumber = 50

func main() {
	printIntro()
	countCircleAreaForUsersRadius()
	doBallFlight()
}

func printIntro() {
	fmt.Println()
	fmt.Println("[][][][][][][][][][]")
	fmt.Println("This program cunts area of a circle. And then shows you how a ball flies :) ")
}

func countCircleAreaForUsersRadius() {
	area := math.Pi * math.Pow(getRadiusFromUser(), 2)
	fmt.Println("Circle's area is: ", area)
}

func getRadiusFromUser() float64 {
	fmt.Println("Please give me radius.", maxRadiusDigits, "digits max.")
	fmt.Print("Input: ")
	return readInputToFloat(maxRadiusDigits)
}

func readInputToFloat(limit int) float64 {
	reader := bufio.NewReader(os.Stdin)

	inputString, _ := reader.ReadString('\n')
	if len(inputString) > limit+1 {
		fmt.Println("Too many symbols...", limit, "is max")
		os.Exit(1)
	}

	input, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return input
}

func doBallFlight() {
	flyABall(askUserHowToFly())
}

func askUserHowToFly() func(arg float64) float64 {
	fmt.Println("Choose: 1 - for drop, 2 - for throw")
	fmt.Print("Input: ")
	switch readInputToFloat(1) {
	case 1:
		fmt.Println("Dropping a ball")
		return dropFunction
	case 2:
		fmt.Println("Throwing a ball")
		return throwFunction
	default:
		fmt.Println("Can't do unknown path. Dropping then.")
		return dropFunction
	}
}

func dropFunction(arg float64) float64 {
	return math.Pow(arg, 2)/1200 + 70
}

func throwFunction(arg float64) float64 {
	return math.Pow(arg-1200, 2)/1200 + 100
}

func flyABall(flyingFunction func(float64) float64) {
	//create screen
	window, cnvs, err := sdlcanvas.CreateWindow(1280, 720, "TheScreen")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer window.Destroy()

	//set initial params for drawing cycle
	canvasHeight := float64(cnvs.Height())
	ballRadius := canvasHeight * 0.04
	currentXPoint := 0.0 + ballRadius
	step := canvasHeight / ballFlightStepsNumber

	//start drawing cycle
	window.MainLoop(func() {
		currentYPoint := flyingFunction(currentXPoint)
		if currentYPoint > canvasHeight {
			fmt.Println("Done.")
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}

		drawGreenBall(cnvs, currentXPoint, currentYPoint, ballRadius)
		currentXPoint += step
		time.Sleep(50 * time.Millisecond)
	})
}

func drawGreenBall(canvas *canvas.Canvas, currentXPoint float64, currentYPoint float64, radius float64) {
	//clean and set style
	prepareCanvas(canvas)

	canvas.BeginPath()
	canvas.Arc(currentXPoint, currentYPoint, radius, 0, math.Pi*2, false)
	canvas.ClosePath()
	canvas.Stroke()
	canvas.Fill()
}

func prepareCanvas(canvas *canvas.Canvas) {
	//clean it with black paint "#000"
	canvas.SetFillStyle("#000")
	canvas.FillRect(0, 0, float64(canvas.Width()), float64(canvas.Height()))

	//set the line and filler style for a figure
	canvas.SetFillStyle(56, 142, 60)
	canvas.SetStrokeStyle("#FFF")
	canvas.SetLineWidth(10)
}
