package unionfind

func NewCollection() Collection {
	return make(map[int]*ElementInfo)
}

type ElementInfo struct {
	Parent int
	Size   uint // Anzahl Elemente im Baum
}

type Collection map[int]*ElementInfo

func (k Collection) MakeSet(x int) {
	k[x] = &ElementInfo{Parent: x, Size: 1}
}

func (k Collection) Union(e, f int) {
	if k[e].Size < k[f].Size {
		e, f = f, e
	}

	k[f].Parent = e
	k[e].Size += k[f].Size
}

func (k Collection) Find(x int) int {
	y := x
	for ; k[y].Parent != y; y = k[y].Parent {
	}

	return y
}
