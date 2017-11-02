package ksplay

import (
    "fmt"
)

type splayNode struct {
    Element int
    Left *splayNode
    Right *splayNode
}

type recordPoint struct {
    sNode *splayNode
    direction bool
}

type recordStack []recordPoint

func MakeEmpty() *splayNode {
    return nil
}

func (k1 *splayNode) DoubleRotateWithLeft() *splayNode{
    k2 := k1.Left
    k3 := k2.Right
    k2.Right = k3.Left
    k1.Left = k3.Right
    k3.Left = k2
    k3.Right = k1
    return k3
}

func (k1 *splayNode) DoubleRotateWithRight() *splayNode{
    k2 := k1.Right
    k3:= k2.Left
    k1.Right = k3.Left
    k2.Left = k3.Right
    k3.Left = k1
    k3.Right = k2
    return k3
}

func (k1 *splayNode) SwapWithLeft() *splayNode{
    k2 := k1.Left
    k1.Left = k2.Right
    k2.Right = k1
    return k2
}

func (k1 *splayNode) SwapWithRight() *splayNode{
    k2 := k1.Right
    k1.Right = k2.Left
    k2.Left = k1
    return k2
}

func (k1 *splayNode) SwingRotateWithRight() *splayNode{

}

func (k1 *splayNode) SwingRotateWithLeft() *splayNode{

}

func (*splayNode) splay(k3 *splayNode, Stack recordStack) {
    
}
