package main

import (
	"fmt"
)

/*
Understanding Structs
Structs allow us to create custom datatypes and also add custom methods to manipulate the data.

*/

// type (Name) struct { body }

type k8s struct {
	apiVersion string
	kind       string
}

func (k *k8s) updateKind() {
	k.kind = "Changed by Reference"
	fmt.Print(k)
}

func (k k8s) updatekind() {
	k.kind = "Changed by Value"
	fmt.Println(k)
}
func changeAPI2(k8sobj *k8s) {
	k8sobj.apiVersion = "v1alpha1"
	fmt.Println(k8sobj)
}
func changeAPI(k8sobj k8s) {
	k8sobj.apiVersion = "v1alpha1"
	fmt.Println(k8sobj)
}

/*
func main() {
	var empty k8s = k8s{}
	fmt.Println(empty)
	// Creating obj for struct
	var k8sobj k8s = k8s{
		apiVersion: "v1",
		kind:       "Pod",
	}

	//  Access values
	var apiVersion = k8sobj.apiVersion
	var kind = k8sobj.kind
	fmt.Println(apiVersion, kind)
	// Updating
	k8sobj.kind = "Deployment"
	fmt.Println(k8sobj)

	//This should update the apiversion to v1alpha1
	changeAPI(k8sobj)
	fmt.Println(k8sobj) // This shouldn't update the struct because we are passing it by value not by reference so it updates its copy
	// var k8sobj1 = &k8s{apiVersion: "v1", kind: "Pod"}
	changeAPI2(&k8sobj)

	var i interface{} = k8sobj
	fmt.Print("After Passing by ref\n", k8sobj)
	fmt.Printf("\n%T", i)
}
*/

func main() {
	var kobj k8s = k8s{
		kind: "V1",
	}
	// fmt.Println("Original Struct", kobj)

	k8s.updatekind(kobj)

	// fmt.Println("After Sending by Value\n", kobj)

	// ref := &kobj

	var val *int

	fmt.Println(val)

	val1 := 1
	val = &val1

	*val = 2

	fmt.Println(val1, *val)

	fmt.Println(val)

	var val3 *int

	val3 = new(int)

	fmt.Println(val3)

	*val3 = 10

	fmt.Println(*val3)
	// k8s.updateKind(ref)

	// fmt.Println("After Changing By Reference", kobj)

}
