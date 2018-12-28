//+build AspectG

// The build tag makes sure the stub is not built in the final build.
package main

import "github.com/alimy/aspectg/module/examples/aspect"

func main() {
	aspect.Build()
}
