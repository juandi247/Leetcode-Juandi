package main

import "fmt"

/*
Trie or TrieTrees they are just Automcpomlete things

They are specially done for searching words, or autocompletions by using a tree node base

Example:

             (root)
            /      \
           c        d
           |        |
           a        o
          / \        \
         t   r        g
        *    *        *

*/

/*
📌 Exercise – Implement Trie Insert & Search

You are asked to implement a Trie (prefix tree) data structure that supports two operations:

Insert(word) → insert a word into the trie.
Search(word) → return true if the word exists in the trie, otherwise false.

Example Trie after inserting: "car", "card", "cat"
(root)
  └── c
       └── a
            ├── r [isWord = true]
            │     └── d [isWord = true]
            └── t [isWord = true]


Search("car") → true

Search("card") → true

Search("cat") → true

Search("ca") → false (not a full word, only a prefix)

Search("dog") → false

Task

👉 Implement Insert(word string) and Search(word string) methods.

Constraints:

Only lowercase English letters a–z.

Word length ≤ 20.

Up to 10⁵ insert/search operations.

*/

type Trie struct {
	root *trieNode
}

type trieNode struct {
	isWord   bool
	children map[rune]*trieNode
}

func ExcerciseTrie() {

	rootNode := &trieNode{}
	myTrie := &Trie{
		root: rootNode,
	}

	myTrie.insertOnTrie("juandi")
	exists := myTrie.searchOnTrie("juandi")
	fmt.Println("the word juandi on  the trie is", exists)

}

func (trie *Trie) insertOnTrie(word string) {
	currentNode := trie.root

	// 1. go through each character of the word to add it, or if it exists already not add it
	for _, character := range word {

		_, exists := currentNode.children[character]
		// 2. If it doesnt exists, we create a new node with the character
		// as soon as we hit this, this is going to be executed always, because the new node doesnt have any data inside
		// so it will conitnue filling the gap
		if !exists {
			currentNode.children[character] = &trieNode{}
		}

		// update the node, so we can check now on the node+1 on the trie
		currentNode = currentNode.children[character]
	}

	// mark last node as Isword because we need to always mark it
	currentNode.isWord = true
}

func (trie *Trie) searchOnTrie(word string) bool {

	currentNode := trie.root

	for _, character := range word {

		_, exists := currentNode.children[character]
		if !exists {
			return false
		}

		currentNode = currentNode.children[character]
	}
	return currentNode.isWord
}

/*
	Cases to delete a word:

	1. The word is alone, it hasnt any shared nodes. Just buble up and delete.

	2. The word has 1 shared node, we want to eliminate CARD and we have CAT too, so we should only eliminate CA

	3. The word is a prefix, we have GONE and want to delete GO, we just delete the mark of the last node and DONT DELET ANYTINH

	4. The word has a prefix, we have ME, and want to delete MESSI, so we delete nodes, until we reach a marked endWord



	CONDITIONS:
	We could combine the 2 and 4.
	We could have some condition that. IF we reach a node, that is
	shared || has a prefix. then stop right there

	We need to first make a DFS to search the word and get the last marked node.
	If the LASTnode has children and ITS marked ? -> delete the mark and FINISH (acomplish case 2)
	If the node

*/

func (trie *Trie) deleteWord(elRoot *trieNode, word string) (bool, error) {
	// solo por validar nada mas
	exists := trie.searchOnTrie(word)
	if !exists {
		return false, fmt.Errorf("the word doesnt even exists on the trie")
	}

	// base case for the word itself
	if len(word) == 0 {
		return true, nil
	}
	// recursive case, until we reach the end
	nextNode, exists := elRoot.children[rune(word[0])]
	needToDelete, _ := trie.deleteWord(nextNode, word[1:])
	
	// base case if he has children (means that it could be case 3 that is a prefix, or its a shared node)
	// this works for case 2 and 3
	if len(elRoot.children) > 0 {
		elRoot.isWord = false
		return false, nil
	}

	if len(word) == 1 {
		elRoot.isWord = false
	}

	// base case if we reach a prefix, deleting GONE, and we find GO. we stop on the O
	if elRoot.isWord {
		return false, nil
	}



	return true, nil
}

func (trie *Trie) doeleteWord(elRoot *trieNode, word string) (bool, error) {
	// Caso base: si ya no hay más letras que procesar
	if len(word) == 0 {
		return false, nil
	}

	// Obtener la letra actual
	letter := rune(word[0])
	nextNode, _ := elRoot.children[letter]
	// Recursión para procesar el resto de la palabra
	needToDelete, _ := trie.deleteWord(nextNode, word[1:])

	// Si el hijo puede eliminarse, lo borramos del map
	if needToDelete {
		delete(elRoot.children, letter)
	}

	// Si llegamos al último nodo de la palabra, desmarcamos isWord
	if len(word) == 1 {
		nextNode.isWord = false
	}

	// Decidir si el nodo actual se puede eliminar:
	// 1. No tiene hijos
	// 2. No está marcado como palabra
	if len(elRoot.children) == 0 && !elRoot.isWord {
		return true, nil
	}

	return false, nil
}
