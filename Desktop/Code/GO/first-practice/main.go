package main

import (
	"fmt"

	"practice-go-mod-url.ng/conditionals"
	data "practice-go-mod-url.ng/data-types" // this is called alias import
	"practice-go-mod-url.ng/functions"
)

var name = "sadiq"

/*
%f or %1f ... for interpolating number as decimal with the optional number after the % symbol and before f representing decimal places
%v defautt string interpolator
%s can be use in place of %v
%d number interpolator
%T returns typeof variable rather than value
Sprintf: returning the interpolated value as a varable instead of printing to standard out as Printf
*/
func main() {
	functions.SayHello(name)

	data.GoDataTypes()
	boolconditionResult, stringconditionResult := conditionals.Conditions(21) //Extract the 2 return data from "Conditions" func

	fmt.Printf("%v: %s\n\n", boolconditionResult, stringconditionResult)

}
