package otto

import "testing"

func TestInner(t *testing.T){
	ot := New()
	// return a function reference, should be an error.
	ot.Run(`
	var result = function(){return 10}
	function a(){
		return result
	}
	`)
	ot.Eval(`a()`)
	// return a function, should be an error.
	ot.Run(`
	function m(){
		return function(){
			return 1
		}
	}
	`)
	ot.Eval(`m()`)
	// inner function define global variable, should be na error.
	ot.Run(`
	function m(){
		var x = 3
		add = function(){
			var m = 0
			return m
		}
	}
	`)
	ot.Eval(`m()`)
}