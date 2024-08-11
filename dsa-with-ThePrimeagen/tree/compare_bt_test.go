package tree

import "testing"

func TestCompareBT(t *testing.T) {
	treeA := &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 2,
			left: &BinaryNode{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &BinaryNode{
				val:   23,
				left:  nil,
				right: nil,
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

	treeB := &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 2,
			left: &BinaryNode{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &BinaryNode{
				val:   23,
				left:  nil,
				right: nil,
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

	treeC := &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 2,
			left: &BinaryNode{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &BinaryNode{
				val:   23,
				left:  nil,
				right: nil,
			},
		},
		right: &BinaryNode{
			val: 9,
			left: &BinaryNode{
				val:   11,
				left:  nil,
				right: nil,
			},
			right: &BinaryNode{
				val:   18,
				left:  nil,
				right: nil,
			},
		},
	}

	if !compare(treeA, treeB) {
		t.Fatalf("Expected %t rec %t", true, compare(treeA, treeB))
	}

	if compare(treeA, treeC) {
		t.Fatalf("Expected %t rec %t", false, compare(treeA, treeC))
	}
}
