package Array

type Any = interface {}
type Array []Any

func (this Array) Slice(start int, end int) Array {
	return this[start:end]
}

func (this *Array) Splice(start int, count int) Array {
	arr := *this

	spliced := append(Array{}, arr[start:start + count]...)
	newThis := append(arr[:start], arr[start + count:]...)

	*this = newThis

	return spliced
}

func (this Array) Every(callback func(item Any, index int, link Array) bool) bool {
	tmp := this.Copy()

	for idx, item := range tmp {
		if !callback(item, idx, tmp) {
			return false
		}
	}

	return true
}

func (this Array) Some(callback func(item Any, index int, link Array) bool) bool {
	tmp := this.Copy()

	for idx, item := range tmp {
		if callback(item, idx, tmp) {
			return true
		}
	}

	return false
}


func (this Array) Map(callback func(item Any, index int, link Array) Any) Array {
	tmp := this.Copy()

	for idx, item := range tmp {
		tmp[idx] = callback(item, idx, tmp)
	}

	return tmp
}

func (this *Array) Push(item ...Any) Any {
	arr := *this

	*this = append(arr, item...)
	return len(*this)
}

func (this *Array) Pop() Any {
	arr := *this

	length := len(arr) -1
	*this = arr[:length]
	return length
}

func (this Array) Copy() Array {
	return append(Array{}, this...)
}

func (this Array) Reduce(callback func(prev Any, cur Any, idx int, link Array) Any, init Any) Any {
	tmp := this.Copy()

	result := init
	prev := init

	for idx, item := range tmp {
		result = callback(prev, item, idx, tmp)
		prev = result
	}

	return result
}

func (this Array) Reverse() Array {
	tmp := this.Copy()
	length := len(tmp)

	offset := 0

	for ; offset < length / 2 ; offset++ {
		l := offset
		r := length - 1 - offset

		tmp[l], tmp[r] = tmp[r], tmp[l]
	}

	return tmp
}

func (this Array) ReduceRight(callback func(prev Any, cur Any, idx int, link Array) Any, init Any) Any {
	tmp := this.Reverse()

	result := init
	prev := init

	for idx, item := range tmp {
		result = callback(prev, item, idx, tmp)
		prev = result
	}

	return result
}

func (this Array) Filter(callback func(item Any, index int, link Array) bool) Array {
	tmp := this.Copy()

	for idx, item := range tmp {
		expr := callback(item, idx, tmp)
		if !expr {
			tmp.Splice(idx, 1)
		}
	}

	return tmp
}

func (this Array) ForEach(callback func(item Any, index int, link Array)) {
	tmp := this.Copy()
	for idx, item := range tmp {
		callback(item, idx, tmp)
	}
}
