//+build AspectG

package aspect

import "github.com/alimy/aspectg/aspect"

// Generate start to use aspect.Build(...) generate AOP code
//go:generate aspectg ./...
func Build() {
	aspect.Build(&Audience{}, &TrackCounter{})
}
