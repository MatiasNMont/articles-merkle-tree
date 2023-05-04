package data

import (
	"crypto/sha256"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		node.Hash = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

func BuildMerkleTree(data [][]byte) *MerkleNode {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	return &nodes[0]
}

func AddDataToMerkleTree(tree *MerkleNode, data []byte) *MerkleNode {
	newNode := NewMerkleNode(nil, nil, data)
	var newData [][]byte
	newData = append(newData, newNode.Hash)

	currentNode := tree
	for currentNode.Left != nil && currentNode.Right != nil {
		if currentNode.Left != nil {
			newData = append(newData, currentNode.Left.Hash)
			currentNode = currentNode.Left
		} else {
			newData = append(newData, currentNode.Right.Hash)
			currentNode = currentNode.Right
		}
	}

	newTree := BuildMerkleTree(newData)

	return newTree
}
