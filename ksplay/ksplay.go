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
	k2 = k1.Right
	k3 = k2.Right
	k2.Right = k3.Left
	k1.Right = k2.Left
	k2.Left = k1
	k3.Left = k2
	return k3
}

func (k1 *splayNode) SwingRotateWithLeft() *splayNode {
	k2 = k1.Left
	k3 = k2.Left
	k2.Left = k3.Right
	k1.Left = k2.Right
	k2.Right = k1
	k3.Right = k2
	return k3
}

func (leaf *splayNode) splay(Stack recordStack) {
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
						leaf = Stack[frontH].SwingRotateWithLeft()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SwingRotateWithLeft()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SwingRotateWithLeft()
					}
				} else { //Right Swing
					if frontH == 0 {
						leaf = Stack[frontH].SwingRotateWithLeft()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].SwingRotateWithRight()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].SwingRotateWithRight()
					}
				}
			} else { //Double Rotate
				if Stack[frontH].Direction { //Left DoubleRotate
					if frontH == 0 {
						leaf = Stack[frontH].DoubleRotateWithLeft()
						break
					}
					if Stack[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].DoubleRotateWithLeft()
					} else { //bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].DoubleRotateWithLeft()
					}
				} else { //Right DoubleRotate
					if FrontH == 0 {
						leaf = Stack[frontH].DoubleRotateWithLeft()
						break
					}
					if Stacl[frontH-1].Direction { //bound with grandpa Left
						Stack[frontH-1].SNode.Left = Stack[frontH].DoubleRotateWithRight()
					} else { // bound with grandpa Right
						Stack[frontH-1].SNode.Right = Stack[frontH].DoubleRotateWithRight()
					}
				}
			}
			currentHeight -= 2
		}
	}
	return leaf
}
