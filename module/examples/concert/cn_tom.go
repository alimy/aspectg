package concert

import "fmt"

type CnTom struct {
	Name string
}

func (t *CnTom) Perform() {
	fmt.Printf("Hello %s\n", t.Name)
}

func (t *CnTom) PlayTrack(trackNum int) {
	fmt.Printf("play track:%d\n", trackNum)
}
