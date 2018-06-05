package utils

import (
	"github.com/astaxie/beego"
	"errors"
)

var inputs chan inputData
var outputs chan outputData
var criticalChan chan int

type inputData struct {
	x int
	y int
}

type outputData struct {
	result int
	error bool
}

func consumeResults(){
	outputs <- outputData{result: -1, error: false}
}

func generateInputs() {
	inputs <- inputData{x: 1, y: 2}
}

func internalCalculationFunc(x, y int) (result int, err error) {
	beego.Debug("calculating z. x:", x, " y:", y)
	z := y
	switch {
	case x == 3:
		beego.Debug("x == 3")
		panic("Failure.")
	case y == 1:
		beego.Debug("y == 1")
		return 0, errors.New("Error!")
	case y == 2:
		beego.Debug("y == 2")
		z = x
	default:
		beego.Debug("default")
		z += x
	}
	retVal := z - 3
	beego.Debug("Returning ", retVal)

	return retVal, nil
}

func processInput(input inputData) {
	defer func() {
		if r := recover(); r != nil {
			beego.Error("Unexpected error occurred: ", r)
			outputs <- outputData{result: 0, error: true}
		}
	}()
	beego.Informational("Received input signal. x:", input.x, " y:", input.y)

	res, err := internalCalculationFunc(input.x, input.y)
	if err != nil {
		beego.Warning("Error in calculation:", err.Error())
	}

	beego.Informational("Returning result: ", res, " error: ", err)
	outputs <- outputData{result: res, error: err != nil}
}

func main() {
	// 设置日志输出级别
	beego.SetLevel(beego.LevelInformational)
	// 设置是否显示文件和行号
	beego.SetLogFuncCall(true)
	inputs = make(chan inputData)
	outputs = make(chan outputData)
	criticalChan = make(chan int)

	beego.Informational("Starting...")
	go consumeResults()
	beego.Informational("Starting receving results.")
	go generateInputs()
	beego.Informational("Starting sending data.")

	for{
		select {
		case input := <-inputs:
			processInput(input)
		case <-criticalChan:
			beego.Critical("Caught value from criticalChan: Go shut down.")
			panic("Shut down due to critical fault.")
		}
	}
}