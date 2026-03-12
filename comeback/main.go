package main

// import "fmt"

// import "fmt"

func main() {

}

/*
reverse a linked list
*/

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func reverseLinkedList(head *Node) *Node {
	//return is the new head

	// X -> A -> B -> C

	/*
	   Traverse until the next value is null.

	   if the node.Next is null means we are on the last value


	   SWAPING

	   X -> should point now to previous

	   x.next=nil

	   <-X

	   then



	*/

	curr := head
	var prev *Node

	for curr != nil {

		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp

	}

	return prev

}

func detectcycle(head *Node) bool {

	/*
	   si tiene cycle es que no termina nunca.

	   si visita un nodo mas de 1 vez significa que hay ciclo

	   2 aproaches:

	   un map, va guardando y pone en true o false si ya fue visitado, y si ya entonces retorna

	   otra que poria evitar o(N) en memoria.
	   podria ser 2 pointers:
	   uno lento y uno rapido

	   uno avanza de a 2 y el otro de a 1. si los dos se encuntn es uqe ya hay ciclo, si el rapido llego a nil primero significa que ya no hay ciclo

	*/

	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast.Next != nil {

		fast = fast.Next.Next

		if fast == nil {
			return false
		}
		slow = slow.Next

		if fast == slow {
			return true
		}

	}

	//because the list ended, so there is no cycle
	return false

}

type LRUCache struct {
	capacity int
	size     int
	head     *Node
	Tail     *Node
	Mapita   map[int]*Node
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		capacity: capacity,
		size:     0,
		head:     nil,
		Tail:     nil,
		Mapita:   make(map[int]*Node, capacity),
	}
}

func (cache *LRUCache) Get(key int) int {

	nodo, exist := cache.Mapita[key]

	if exist == false {
		return 0
	}

	if cache.size<=1 || cache.head==nodo{
		return nodo.Val
	}

	
	beforeNodoPrev:= nodo.Prev
	beforeNodoPrev.Next = nodo.Next
	
	if cache.Tail == nodo && nodo.Prev!=nil{
		cache.Tail = nodo.Prev
	}


	if nodo.Prev!=nil{
		nodo.Next.Prev= beforeNodoPrev
	}
	

	nodo.Prev = nil
	cache.head.Prev= nodo
	nodo.Next = cache.head
	cache.head=nodo



	return nodo.Val
}

func (cache *LRUCache) put(key int, value int) {

	_, exist := cache.Mapita[key]

	if exist == true {
		// el nodo ya existia antes por lo tanto early return, esto es solo para manejar el caso, no evalues nada( podria agregarle un error o algo pero meh)
		return
	}

	newNode := &Node{
		Key:  key,
		Val:  value,
		Next: cache.head,
		Prev: nil,
	}

	if cache.size == 0 {
		cache.size++
		cache.head = newNode
		cache.Tail = newNode
		cache.Mapita[key] = newNode
		return
	}

	if cache.size < cache.capacity {
		cache.size++
		// sigifnica que solo lo agregamos y ya porque sigue habiedno size suficiente dentro del mapa
		cache.head.Prev = newNode
		cache.head = newNode

		cache.Mapita[key] = newNode
		return
	}

	// quitar el ultimooo, porque estamos usando fixed sizes por securityyy allcoation estaticaa
	delete(cache.Mapita, cache.Tail.Key)
	cache.Tail = cache.Tail.Prev
	cache.head.Prev = newNode
	cache.head = newNode
	cache.Mapita[key] = newNode

}

