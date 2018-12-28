package concert

type Performance interface {
	Perform()
}

// Start just begin a sample instance
func Start() {
	cnTom := &CnTom{Name: "AOP"}
	cnTom.Perform()
	cnTom.PlayTrack(8013)
}
