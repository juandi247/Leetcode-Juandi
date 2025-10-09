package main

type Trie struct {
	rootNode *Node
}

type Node struct {
	lettersMap map[rune]*Node
	isEndWord  bool
}

func constructor() Trie {

	rootNode := &Node{
		lettersMap: make(map[rune]*Node),
		isEndWord:  false,
	}

	return Trie{
		rootNode: rootNode}
}

//insert
func (t *Trie) insert(word string) {

	currentNode := t.rootNode
	for _, v := range word {
		//buscar la primera letra en el currentNode

		nextNode, exists := currentNode.lettersMap[v]

		if exists {
			currentNode = nextNode
			continue
		}

		newNode := &Node{
			lettersMap: make(map[rune]*Node),
			isEndWord:  false,
		}

		currentNode.lettersMap[v] = newNode
		currentNode = newNode
	}

	currentNode.isEndWord = true
}

//search
func (t *Trie) search(word string) bool {

	currentNode := t.rootNode

	for _, v := range word {
		nextNode, exists := currentNode.lettersMap[v]

		if !exists {
			return false
		}

		currentNode = nextNode
	}

	if !currentNode.isEndWord {
		return false
	}

	return true
}

//starts with prexis
func (t *Trie) search(word string) bool {

	currentNode := t.rootNode

	for _, v := range word {
		nextNode, exists := currentNode.lettersMap[v]

		if !exists {
			return false
		}

		currentNode = nextNode
	}

	return true
}
