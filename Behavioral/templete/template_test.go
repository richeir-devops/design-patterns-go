package templete

import "testing"

func Test01(t *testing.T) {
	androidBuilder := &AndroidBuilder{}
	androidBuilder.build()

	IosBuilder := &IosBuilder{}
	IosBuilder.build()
}
