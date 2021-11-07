package iterator

///////////////////////////////////////////
// 迭代器模式（Iterator）：
// 就是可以遍历元素，类似于C#的 ICollection 接口和 foreach 循环；
///////////////////////////////////////////

///////////////////////////////////////////

type collection interface {
	getNext()
	hasNext()
}

///////////////////////////////////////////

///////////////////////////////////////////
type Station struct {
	name      string
	frequency float32
}

type RadioStation struct {
	stations   []*Station
	writeIndex int
	readIndex  int
}

func (rt *RadioStation) addStation(s *Station) {
	rt.stations[rt.writeIndex] = s
	rt.writeIndex++
}

func (rt *RadioStation) getNext() *Station {
	currentValue := rt.stations[rt.readIndex]
	rt.readIndex++
	return currentValue
}

func (rt *RadioStation) hasNext() bool {
	return rt.readIndex < rt.writeIndex
}

///////////////////////////////////////////
