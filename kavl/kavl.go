package kavl

import (
	"fmt"
)

type avlNode struct {
	Element int
	Height  int
	Left    *avlNode
	Right   *avlNode
}

func MakeEmpty() *avlNode {
	return nil
}

func hMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(n *avlNode) int {
	if n == nil {
		return -1
	} else {
		return n.Height
	}
}

func (n *avlNode) Find(element int) *avlNode {
	for {
		if n == nil {
			return n
		} else if element == n.Element {
			return n
		} else if element < n.Element {
			n = n.Left
		} else {
			n = n.Right
		}
	}
}

func (k2 *avlNode) SingleRotateWithLeft() *avlNode {
	k1 := k2.Left
	k2.Left = k1.Right
	k1.Right = k2
	k2.Height = hMax(getHeight(k2.Left), getHeight(k2.Right)) + 1
	k1.Height = hMax(getHeight(k1.Left), k2.Height) + 1
	return k1
}

func (k2 *avlNode) SingleRotateWithRight() *avlNode {
	k1 := k2.Right
	k2.Right = k1.Left
	k1.Left = k2
	k2.Height = hMax(getHeight(k2.Left), getHeight(k2.Right)) + 1
	k1.Height = hMax(k2.Height, getHeight(k1.Right)) + 1
	return k1
}

func (k3 *avlNode) DoubleRotateWithLeft() *avlNode {
	k1 := k3.Left
	k2 := k1.Right
	k1.Right = k2.Left
	k3.Left = k2.Right
	k2.Left = k1
	k2.Right = k3
	k1.Height = hMax(getHeight(k1.Left), getHeight(k1.Right)) + 1
	k3.Height = hMax(getHeight(k3.Left), getHeight(k3.Right)) + 1
	k2.Height = hMax(k1.Height, k3.Height) + 1
	return k2
}

func (k3 *avlNode) DoubleRotateWithRight() *avlNode {
	k1 := k3.Right
	k2 := k1.Left
	k1.Left = k2.Right
	k3.Right = k2.Left
	k2.Left = k3
	k2.Right = k1
	k1.Height = hMax(getHeight(k1.Left), getHeight(k1.Right)) + 1
	k3.Height = hMax(getHeight(k3.Left), getHeight(k3.Right)) + 1
	k2.Height = hMax(k3.Height, k1.Height) + 1
	return k2
}

func (n *avlNode) Insert(element int) *avlNode {
	if n == nil {
		t := avlNode{
			Element: element,
			Height:  0,
		}
		return &t
	} else if element > n.Element {
		n.Right = n.Right.Insert(element)
		if getHeight(n.Right)-getHeight(n.Left) == 2 {
			if element > n.Right.Element {
				return n.SingleRotateWithRight()
			} else {
				return n.DoubleRotateWithRight()
			}
		}
	} else if element < n.Element {
		n.Left = n.Left.Insert(element)
		if getHeight(n.Left)-getHeight(n.Right) == 2 {
			if element < n.Left.Element {
				return n.SingleRotateWithLeft()
			} else {
				return n.DoubleRotateWithLeft()
			}
		}
	} else {
		return n
	}
	n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
	return n
}

func PrintTree(n *avlNode) {
	defer fmt.Println("")
	queue := make(chan *avlNode, 500)
	queue <- n
	for {
		if len(queue) == 0 {
			return
		}
		t := <-queue
		fmt.Printf("Element: %d\tHeight: %d\n", t.Element, t.Height)
		if t.Left != nil {
			queue <- t.Left
		}
		if t.Right != nil {
			queue <- t.Right
		}
	}
}

//DeleteMin find mininum number in the tree, n can't be a nil.
func (n *avlNode) DeleteMin() (*avlNode, int) {
	if n.Left == nil {
		if n.Right == nil {
			return nil, n.Element
		} else {
			return n.Right, n.Element
		}
	}
	var backStack []*avlNode
	fNode := n
	cNode := n.Left
	for {
		if cNode.Left == nil {
			if cNode.Right == nil {
				fNode.Left = nil
				//fNode.Height = getHeight(fNode.Right) + 1 // fNode.Left already nil
				for item := len(backStack) - 1; item >= 0; item-- {
					if getHeight(fNode.Right)-getHeight(fNode.Left) == 2 {
						if getHeight(fNode.Right.Right) >= getHeight(fNode.Right.Left) {
							backStack[item].Left = fNode.SingleRotateWithRight()
						} else {
							backStack[item].Left = fNode.DoubleRotateWithRight()
						}
					}
					backStack[item].Height = hMax(getHeight(backStack[item].Left), getHeight(backStack[item].Right)) + 1
					fNode = backStack[item]
				}
				if getHeight(n.Right)-getHeight(n.Left) == 2 {
					if getHeight(n.Right.Right) >= getHeight(n.Right.Left) {
						n = n.SingleRotateWithRight()
					} else {
						n = n.DoubleRotateWithRight()
					}
				}
				n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
				return n, cNode.Element
			} else {
				fNode.Left = cNode.Right
				//fNode.Height = hMax(getHeight(fNode.Left), getHeight(fNode.Right)) + 1
				for item := len(backStack) - 1; item >= 0; item-- {
					if getHeight(fNode.Right)-getHeight(fNode.Left) == 2 {
						if getHeight(fNode.Right.Right) >= getHeight(fNode.Right.Left) {
							backStack[item].Left = fNode.SingleRotateWithRight()
						} else {
							backStack[item].Left = fNode.DoubleRotateWithRight()
						}
					}
					backStack[item].Height = hMax(getHeight(backStack[item].Left), getHeight(backStack[item].Right)) + 1
					fNode = backStack[item]
				}
				if getHeight(n.Right)-getHeight(n.Left) == 2 {
					if getHeight(n.Right.Right) >= getHeight(n.Right.Left) {
						n = n.SingleRotateWithRight()
					} else {
						n = n.DoubleRotateWithRight()
					}
				}
				n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
				return n, cNode.Element
			}
		} else {
			backStack = append(backStack, fNode)
			fNode = cNode
			cNode = cNode.Left
		}
	}
}

func (n *avlNode) FindMin() int {
	if n.Left == nil {
		return n.Element
	}
	m := n.Left.FindMin()
	if m < n.Element {
		return m
	} else {
		return n.Element
	}
}

func (n *avlNode) Delete(element int) *avlNode {
	if n == nil {
		return nil
	} else if element < n.Element {
		n.Left = n.Left.Delete(element)
	} else if element > n.Element {
		n.Right = n.Right.Delete(element)
	} else if n.Left != nil && n.Right != nil {
		/*
			n.Right, n.Element = n.Right.DeleteMin()
			n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
			if n.Left.Height-getHeight(n.Right) == 2 {
				if getHeight(n.Left.Left) >= getHeight(n.Left.Right) {
					n = n.SingleRotateWithLeft()
				} else {
					n = n.DoubleRotateWithLeft()
				}
			}*/
		n.Element = n.Right.FindMin()
		n.Right = n.Right.Delete(n.Element)
		if n.Left.Height-getHeight(n.Right) == 2 {
			if getHeight(n.Left.Left) >= getHeight(n.Left.Right) {
				n = n.SingleRotateWithLeft()
			} else {
				n = n.DoubleRotateWithLeft()
			}
		}
		n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
		return n
	} else {
		if n.Left != nil {
			n = n.Left
		} else {
			n = n.Right
		}
		return n
	}
	diff := getHeight(n.Left) - getHeight(n.Right)
	switch diff {
	case 2:
		if getHeight(n.Left.Left) >= getHeight(n.Left.Right) {
			n = n.SingleRotateWithLeft()
		} else {
			n = n.DoubleRotateWithLeft()
		}
		return n
	case -2:
		if getHeight(n.Right.Right) >= getHeight(n.Right.Left) {
			n = n.SingleRotateWithRight()
		} else {
			n = n.DoubleRotateWithRight()
		}
		return n
	}
	n.Height = hMax(getHeight(n.Left), getHeight(n.Right)) + 1
	return n
}

func calHeight(root *avlNode) int {
	if root == nil {
		return -1
	} else {
		return hMax(calHeight(root.Left), calHeight(root.Right)) + 1
	}
}

func IsBalanced(root *avlNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	l := calHeight(root.Left)
	if root.Left != nil && root.Left.Height != l {
		fmt.Printf("%d %d %d\n", root.Left.Element, l, root.Left.Height)
		//  return false
	}
	r := calHeight(root.Right)
	if root.Right != nil && root.Right.Height != r {
		fmt.Printf("%d %d %d\n", root.Right.Element, r, root.Right.Height)
		//   return false
	}
	if l > r {
		if l-r > 1 {
			fmt.Printf("-- %d %d %d\n", root.Element, l-r, root.Left.Element)
			return false
		}
	} else {
		if r-l > 1 {
			fmt.Printf("%d %d %d  %d--\n", root.Element, l, root.Left.Element, root.Right.Element)
			return false
		}
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}
