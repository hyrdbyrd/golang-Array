package Array

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
	for idx, item := range this {
		if !callback(item, idx, this) {
			return false
		}
	}

	return true
}

func (this Array) Some(callback func(item Any, index int, link Array) bool) bool {
	for idx, item := range this {
		if callback(item, idx, this) {
			return true
		}
	}

	return false
}


func (this Array) Map(callback func(item Any, index int, link Array) Any) Array {
	for idx, item := range this {
		this[idx] = callback(item, idx, this)
	}

	return this
}

func (this *Array) Push(item ...Any) Any {
	arr := *this

	*this = append(arr, item...)
	return len(*this)
}

func (this *Array) Pop() Any {
	arr := *this

	length := len(arr) - 1
	*this = arr[:length]
	return length
}

func (this *Array) Shift() Any {
	arr := *this

	*this = append(Array{}, arr[1:]...)
	return len(*this)
}

func (this *Array) Unshift(item ...Any) Any {
	item = append(item, *this...)

	*this = append(Array{}, item...)
	return len(*this)
}

func (this Array) Reduce(callback func(prev Any, cur Any, idx int, link Array) Any, init Any) Any {
	result := init
	prev := init

	for idx, item := range this {
		result = callback(prev, item, idx, this)
		prev = result
	}

	return result
}

func (this Array) Reverse() Array {
	length := len(this)

	offset := 0

	for ; offset < length / 2 ; offset++ {
		l := offset
		r := length - 1 - offset

		this[l], this[r] = this[r], this[l]
	}

	return this
}

func (this Array) ReduceRight(callback func(prev Any, cur Any, idx int, link Array) Any, init Any) Any {
	result := init
	prev := init

	for idx, item := range this.Reverse() {
		result = callback(prev, item, idx, this)
		prev = result
	}

	return result
}

func (this Array) Filter(callback func(item Any, index int, link Array) bool) Array {
	for idx, item := range this {
		if !callback(item, idx, this) {
			this.Splice(idx, 1)
		}
	}

	return this
}

func (this Array) ForEach(callback func(item Any, index int, link Array)) {
	for idx, item := range this {
		callback(item, idx, this)
	}
}
