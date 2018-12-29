//+build AspectG

package aspect

import (
	. "github.com/alimy/aspectg/aspect"
	"github.com/alimy/aspectg/module/examples/media"
)

type TrackCounter struct {
	CountTrack Advisor `aspect:"execution(.../media.CompactDisc(trackNumber int)) && args(trackNumber),Before"`
}

func (t *TrackCounter) countTrack(jp JoinPoint) {
	args := jp.Args()
	if trackNumber, ok := args["trackNumber"]; ok {
		number, _ := trackNumber.(int)
		media.TrackCount(number)
	}
}

// AspectOf make a aspect rules instance
func (t *TrackCounter) AspectOf() interface{} {
	t.CountTrack = AdvisorFunc(t.countTrack)
	return t
}
