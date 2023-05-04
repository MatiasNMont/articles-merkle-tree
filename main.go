package main

import (
	"fmt"
	merkel "golang-proyect/data"
)

func main() {
	data := [][]byte{
		[]byte("Hello"),
		[]byte("World"),
	}

	fmt.Printf("Original data: %v\n", data)

	tree := merkel.BuildMerkleTree(data)
	fmt.Printf("New tree1 root hash: %x\n", tree.Hash)
	newData := []byte("New data")
	fmt.Printf("New data: %s\n", newData)

	newTree := merkel.AddDataToMerkleTree(tree, newData)

	fmt.Printf("New tree2 root hash: %x\n", newTree.Hash)

	newTransaction := []byte("Deploy a1")
	newTree = merkel.AddDataToMerkleTree(newTree, newTransaction)
	fmt.Printf("New tree3 root hash: %x\n", newTree.Hash)

}
