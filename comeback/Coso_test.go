package main

import ("testing"
 "fmt")

func TestReverse(t *testing.T) {
    // 1 -> 2 -> 3 -> 4 -> 5 -> nil
    head := &Node{Val: 1}
    head.Next = &Node{Val: 2}
    head.Next.Next = &Node{Val: 3}
    head.Next.Next.Next = &Node{Val: 4}
    head.Next.Next.Next.Next = &Node{Val: 5}

    reversed := reverseLinkedList(head)

    expected := []int{5, 4, 3, 2, 1}
    curr := reversed

fmt.Println("reversed::::")
    printList(reversed)
    for i, v := range expected {
        if curr == nil {
            t.Fatalf("expected value %d at position %d but got nil", v, i)
        }
        if curr.Val != v {
            t.Fatalf("expected %d at position %d, got %d", v, i, curr.Val)
        }
        curr = curr.Next
    }

    if curr != nil {
        t.Fatalf("list has extra elements")
    }
}


func printList(head *Node) {
    curr := head
    for curr != nil {
        fmt.Printf("%d -> ", curr.Val)
        curr = curr.Next
    }
    fmt.Println("nil")
}




func TestDetectCycle(t *testing.T) {
	t.Run("no cycle", func(t *testing.T) {
		// 1 -> 2 -> 3 -> 4 -> nil
		head := &Node{Val: 1}
		head.Next = &Node{Val: 2}
		head.Next.Next = &Node{Val: 3}
		head.Next.Next.Next = &Node{Val: 4}

		if detectcycle(head) {
			t.Fatalf("expected no cycle, got cycle")
		}
	})

	t.Run("has cycle", func(t *testing.T) {
		// 1 -> 2 -> 3 -> 4
		//      ^         |
		//      |_________|
		head := &Node{Val: 1}
		n2 := &Node{Val: 2}
		n3 := &Node{Val: 3}
		n4 := &Node{Val: 4}

		head.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n2 // cycle

		if !detectcycle(head) {
			t.Fatalf("expected cycle, got no cycle")
		}
	})

	t.Run("single node no cycle", func(t *testing.T) {
		head := &Node{Val: 1}

		if detectcycle(head) {
			t.Fatalf("expected no cycle, got cycle")
		}
	})

	t.Run("single node with cycle", func(t *testing.T) {
		head := &Node{Val: 1}
		head.Next = head

		if !detectcycle(head) {
			t.Fatalf("expected cycle, got no cycle")
		}
	})
}





func TestLRUCache(t *testing.T) {

	t.Run("basic operations", func(t *testing.T) {
		cache := Constructor(2)

		cache.put(1, 1)
		cache.put(2, 2)

		if v := cache.Get(1); v != 1 {
			t.Fatalf("expected 1, got %d", v)
		}

		cache.put(3, 3)
        fmt.Println(cache.Mapita)
		if v := cache.Get(2); v != 0 {
			t.Fatalf("expected 0 (evicted), got %d", v)
		}

		cache.put(4, 4)

		if v := cache.Get(1); v != 0 {
			t.Fatalf("expected 0 (evicted), got %d", v)
		}

		if v := cache.Get(3); v != 3 {
			t.Fatalf("expected 3, got %d", v)
		}

		if v := cache.Get(4); v != 4 {
			t.Fatalf("expected 4, got %d", v)
		}
	})

	t.Run("capacity one", func(t *testing.T) {
		cache := Constructor(1)

		cache.put(1, 1)

		if v := cache.Get(1); v != 1 {
			t.Fatalf("expected 1, got %d", v)
		}

		cache.put(2, 2)

		if v := cache.Get(1); v != 0 {
			t.Fatalf("expected 0 (evicted), got %d", v)
		}

		if v := cache.Get(2); v != 2 {
			t.Fatalf("expected 2, got %d", v)
		}
	})

	t.Run("get missing key", func(t *testing.T) {
		cache := Constructor(2)

		if v := cache.Get(42); v != 0 {
			t.Fatalf("expected 0 for missing key, got %d", v)
		}
	})

	t.Run("put duplicate key", func(t *testing.T) {
		cache := Constructor(2)

		cache.put(1, 1)
		cache.put(1, 10)

		if v := cache.Get(1); v != 1 {
			t.Fatalf("expected 1, got %d", v)
		}
	})
}