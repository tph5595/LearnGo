//inspiration taken from https://learnxinyminutes.com/docs/go/

package main

import "fmt"

func main() {
	i := tuple{1, 2}
	fmt.Println(i)
}

//so instead of having classes go seperates the data and methods structures
//it does this by having structs(like c) that can only store data and
//interfaces which define methods (like java, name, inputs, outputs)
//to implement an interface into a struct you do the following for
//all of the methods defined in the interface
//
//  example:
//    func (varName structName) Method(inputs) outputs{
//      method.definition()
//    }

//defines a Stringer interface with a single method
type Stringer interface {
	String() string
}

//define a new type tuple that holds 2 values
type tuple struct {
	x, y int
}

//give our new value a string method(it will then have implemented the
//above interface)
//also important to note that go uses a String() how java uses it's
//toString()
func (t tuple) String() string {
	//Sprintf() creates a string using printf syntax
	return fmt.Sprintf("(%d,%d)", t.x, t.y)
}
