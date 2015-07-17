package numbers


import (
	"testing"
	"github.com/robertkrimen/otto"
	"fmt"
)


func TestSimple(t *testing.T) {
	vm := otto.New()
	_,err := vm.Run(`
	var array  = [1, 2, 3];
	var array1 = [0, 1, 2];
	var array2 = [3, 4, 5];
	console.log(numbers.statistic.mean(array));
	console.log(numbers.statistic.median(array));
	console.log(numbers.statistic.mode(array));
	console.log(numbers.statistic.standardDev(array));
//	console.log(numbers.statistic.randomSample(lower, upper, n));
	console.log(numbers.statistic.correlation(array1, array2));
	`)
	if err != nil {
		fmt.Println(err)
	}
}