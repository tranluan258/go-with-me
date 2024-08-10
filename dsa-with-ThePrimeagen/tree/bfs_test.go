package tree

import "testing"

func TestBFS(t *testing.T) {
	tree := &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 2,
			left: &BinaryNode{
				val: 5,
				left: &BinaryNode{
					val:   7,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   10,
					left:  nil,
					right: nil,
				},
			},
			right: &BinaryNode{
				val: 23,
				left: &BinaryNode{
					val:   17,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   13,
					left:  nil,
					right: nil,
				},
			},
		},
		right: &BinaryNode{
			val: 9,
			left: &BinaryNode{
				val: 11,
				left: &BinaryNode{
					val:   35,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   21,
					left:  nil,
					right: nil,
				},
			},
			right: &BinaryNode{
				val: 18,
				left: &BinaryNode{
					val:   15,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   33,
					left:  nil,
					right: nil,
				},
			},
		},
	}

	res := BFS(tree, 9)
	if !res {
		t.Fatalf(`Expected %t rec %t`, true, res)
	}

	notExist := BFS(tree, 50)
	if notExist {
		t.Fatalf(`Expected %t rec %t`, false, notExist)
	}
}
