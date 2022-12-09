package main

import "fmt"

/*
Understanding Structs
Structs allow us to create custom datatypes and also add custom methods to manipulate the data.

*/

// type (Name) struct { body }

type k8s struct {
	apiVersion string
	kind       string
}

func changeAPI2(k8sobj *k8s) {
	k8sobj.apiVersion = "v1alpha1"
	fmt.Println(k8sobj)
}
func changeAPI(k8sobj k8s) {
	k8sobj.apiVersion = "v1alpha1"
	fmt.Println(k8sobj)
}
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
	fmt.Print("After Passing by ref\n", k8sobj)
}
