package dispatcher


/*type Element struct {
	Value int
}*/

type FIFOQueue struct {
	list [] *Element //Change type for more performance
	size int
	head int
	tail int
	count int
}

func NewFIFOQueue(size int) *FIFOQueue{
	
	return &FIFOQueue{
		list: make([] *Element, size),
		size: size,
		head: 0,
		tail: 0,
		count: 0,
	}
}

func (q *FIFOQueue) Push( element *Element) {
	q.list[q.tail] = element
	q.tail = (q.tail+1) % q.size
	q.count++
}

func (q *FIFOQueue) Pop () *Element{

	if q.count == 0 {
		return nil
	}
	
	element := q.list[q.head]
	q.head = (q.head + 1) % q.size
	q.count--
	return element
}

func (q *FIFOQueue) Size() int {
	return q.size
}

func (q *FIFOQueue) NumElements() int {
	return q.count
}


/*
func main() {
	queue := NewFIFOQueue(10)

	for i:=0; i<10; i++ {
		queue.Push(&Element{i})
	}

	for i:=0; i<11; i++ {
		println(queue.Pop())
	}
}
*/
/* Output
&{0}
&{1}
&{2}
&{3}
&{4}
&{5}
&{6}
&{7}
&{8}
&{9}
<nil>
*/
