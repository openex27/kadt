package ksplay

import (
	"fmt"
)

type splayNode struct {
	Element int
	Left    *splayNode
	Right   *splayNode
}

//recordPoint use for constitute recordStack,
//direction descirbe search direction,
//if search left subtree then it will be true ,otherwise false
type recordPoint struct {
	SNode     *splayNode
	Direction bool
}

type recordStack []recordPoint

func MakeEmpty() *splayNode {
	return nil
}

func (k1 *splayNode) DoubleRotateWithLeft() *splayNode {
	k2 := k1.Left
	k3 := k2.Right
	k2.Right = k3.Left
	k1.Left = k3.Right
	k3.Left = k2
	k3.Right = k1
	return k3
}

func (k1 *splayNode) DoubleRotateWithRight() *splayNode {
	k2 := k1.Right
	k3 := k2.Left
	k1.Right = k3.Left
	k2.Left = k3.Right
	k3.Left = k1
	k3.Right = k2
	return k3
}

func (k1 *splayNode) SwapWithLeft() *splayNode {
	k2 := k1.Left
	k1.Left = k2.Right
	k2.Right = k1
	return k2
}

func (k1 *splayNode) SwapWithRight() *splayNode {
	k2 := k1.Right
	k1.Right = k2.Left
	k2.Left = k1
	return k2
}

func (k1 *splayNode) SwingRotateWithRight() *splayNode {
	k2 := k1.Right
	k3 := k2.Right
	k2.Right = k3.Left
	k1.Right = k2.Left
	k2.Left = k1
	k3.Left = k2
	return k3
}

func (k1 *splayNode) SwingRotateWithLeft() *splayNode {
	k2 := k1.Left
	k3 := k2.Left
	k2.Left = k3.Right
	k1.Left = k2.Right
	k2.Right = k1
	k3.Right = k2
	return k3
}

func (leaf *splayNode) splay(Stack recordStack) *splayNode {
	currentHeight := len(Stack) - 1
	for {
		if currentHeight < 0 {
			break
		} else if currentHeight == 0 {
			if Stack[0].Direction {
				leaf = Stack[0].SNode.SwapWithLeft()
			} else {
				leaf = Stack[0].SNode.SwapWithRight()
			}
			break
		} else {
			frontH := currentHeight - 1
			if Stack[currentHeight].Direction == Stack[frontH].Direction { //Swing Rotate
				if Stack[frontH].Direction { //Left Swing
					if frontH == 0 {
						leaf = Stack[frontH].SNode.SwingRotateWithLeft()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SNode.SwingRotateWithLeft()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SNode.SwingRotateWithLeft()
					}
				} else { //Right Swing
					if frontH == 0 {
						leaf = Stack[frontH].SNode.SwingRotateWithRight()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SNode.SwingRotateWithRight()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SNode.SwingRotateWithRight()
					}
				}
			} else { //Double Rotate
				if Stack[frontH].Direction { //Left DoubleRotate
					if frontH == 0 {
						leaf = Stack[frontH].SNode.DoubleRotateWithLeft()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SNode.DoubleRotateWithLeft()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SNode.DoubleRotateWithLeft()
					}
				} else { //Right DoubleRotate
					if frontH == 0 {
						leaf = Stack[frontH].SNode.DoubleRotateWithRight()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SNode.DoubleRotateWithRight()
					} else { // bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SNode.DoubleRotateWithRight()
					}
				}
			}
			currentHeight -= 2
		}
	}
	return leaf
}


func (root *splayNode) Insert(value int) *splayNode {
	Stack := recordStack{}
	if root == nil {
		tempNode := splayNode{
			Element: value,
		}
		return &tempNode
	}
	for {
		if root.Element > value {
			t := recordPoint{
				SNode:     root,
				Direction: true,
			}
			Stack = append(Stack, t)
			if root.Left != nil {
				root = root.Left
			} else {
				tempNode := splayNode{
					Element: value,
				}
				root.Left = &tempNode
				return root.Left.splay(Stack)
			}
		} else if root.Element < value {
			t := recordPoint{
				SNode:     root,
				Direction: false,
			}
			Stack = append(Stack, t)
			if root.Right != nil {
				root = root.Right
			} else {
				tempNode := splayNode{
					Element: value,
				}
				root.Right = &tempNode
				return root.Right.splay(Stack)
			}
		} else if root.Element == value {
			return root.splay(Stack)
		}
	}
}


func PrintTreePre(root *splayNode, ch chan string) {
    if root == nil {
	if ch == nil {
	fmt.Printf("nil\n")
	}else{
	    ch <- fmt.Sprintf("nil ")
	}
	return
    }
    if ch == nil{
	fmt.Printf("%d\n",root.Element)
    }else{
	ch <- fmt.Sprintf("%d ",root.Element)
    }
    PrintTreePre(root.Left, ch)
    PrintTreePre(root.Right, ch)
}
