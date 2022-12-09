package scope

//Global and Public accessable everywhere

var GlobalVar string = "GlobalLevel"

// Private Only accessible within the package
var packageLevel string = "Package Level"

// Private function only accessible to same package.
func test() {
	var local string = "LocalVar"
	local = local + "abc"

}

// //Type Meta
// apiVersion:
// kind:

// //Object Meta
// metadata:
//  name:
//  labels:
//  namespace;

// spec:
