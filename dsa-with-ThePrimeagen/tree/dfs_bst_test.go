package tree

import "testing"

func TestDFS(t *testing.T) {
	tree := &BinaryNode{
		val: 19,
		left: &BinaryNode{
			val: 15,
			left: &BinaryNode{
				val: 9,
				left: &BinaryNode{
					val:   7,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   11,
					left:  nil,
					right: nil,
				},
			},
			right: &BinaryNode{
				val: 17,
				left: &BinaryNode{
					val:   16,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   18,
					left:  nil,
					right: nil,
				},
			},
		},
		right: &BinaryNode{
			val: 30,
			left: &BinaryNode{
				val: 24,
				left: &BinaryNode{
					val:   23,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   25,
					left:  nil,
					right: nil,
				},
			},
			right: &BinaryNode{
				val: 35,
				left: &BinaryNode{
					val:   32,
					left:  nil,
					right: nil,
				},
				right: &BinaryNode{
					val:   36,
					left:  nil,
					right: nil,
				},
			},
		},
	}

	res := dsf(tree, 17)
	if !res {
		t.Fatalf(`Expected %t rec %t`, true, res)
	}

	notExist := dsf(tree, 50)
	if notExist {
		t.Fatalf(`Expected %t rec %t`, false, notExist)
	}
}
