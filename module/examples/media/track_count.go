package media

var (
	playCount = make(map[int]uint)
)
func TrackCount(trackNum int) {
	playCount[trackNum]++
}
