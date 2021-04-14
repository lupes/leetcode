package common

type Bitmap struct {
	size uint64
	flag []byte
}

func NewBitmap(size uint64) *Bitmap {
	return &Bitmap{
		size: size,
		flag: make([]byte, size>>3+1),
	}
}

func (self *Bitmap) index(id uint64) (uint64, uint64) {
	//fmt.Printf("id:%d index:%d position:%d\n", id, id>>3, id&7)
	return id >> 3, id & 7
}

func (self *Bitmap) Add(id uint64) bool {
	if id > self.size {
		return false
	}
	index, position := self.index(id)
	self.flag[index] |= 1 << position
	return true
}

func (self *Bitmap) Del(id uint64) bool {
	if id > self.size {
		return false
	}
	index, position := self.index(id)
	if self.flag[index]&(1<<position) > 0 {
		self.flag[index] ^= 1 << position
	}

	return true
}

func (self *Bitmap) Contain(id uint64) bool {
	if id > self.size {
		return false
	}
	index, position := self.index(id)
	return self.flag[index]&(1<<position) > 0
}
