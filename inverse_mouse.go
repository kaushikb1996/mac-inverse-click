package main

import (
	"fmt"
	"os"
	"os/exec"
)

func executeProgram() {
	fmt.Println("\nInversing Mouse...")

	//Getting the current path to run the applescript, you can hard code your path here if you would like to move the mouse.applescript elsewhere.
	//dir:= "Your Path"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s", err)
	}

	//Executing the applescript
	out, errs := exec.Command("osascript", dir+"/mouse.applescript").Output()
	if errs != nil {
		fmt.Printf("%s", errs)
		return
	}
	output := string(out[:])
	fmt.Println("Done" + output)

	return
}

func main() {
	executeProgram()
}
