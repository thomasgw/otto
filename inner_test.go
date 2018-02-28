package otto

import (
	"testing"
)

func TestInner(t *testing.T){
	ot := New()
	//ot.Run(`
	//var BankContract = function (){}
	//BankContract.prototype = function(){
	//	f1:function(){retunn 10}
	//}
	//`)
	//ot.Object("BankContract")
	// return a function reference, should be an error.
	//ot.Run(`
	//var result = function(){return 10}
	//function a(){
	//	return result
	//}
	//`)
	//ot.Run(`a()`)
	//// return a function, should be an error.
	//ot.Run(`
	//function m(){
	//	return function(){
	//		return 1
	//	}
	//}
	//`)
	//ot.Eval(`m()`)
	//// inner function define global variable, should be na error.
	//ot.Run(`
	//function m(){
	//	var x = 3
	//	add = function(){
	//		var m = 0
	//		return m
	//	}
	//}
	//`)
	//ot.Eval(`m()`)
	ot.Run(``)
}