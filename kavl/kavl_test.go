package kavl

import (
	"math/rand"
	"testing"
)

const (
	Length    = 30000
	LengthMin = 3000
)

var (
	avlTree = MakeEmpty()
)

func Test_FastTest(t *testing.T) {
	var i int
	for i = 1; i < LengthMin; i++ {
		if rand.Intn(2) != 1 {
			avlTree = avlTree.Insert(rand.Intn(LengthMin))
		} else {
			avlTree = avlTree.Delete(rand.Intn(LengthMin))
		}
		if !IsBalanced(avlTree) {
			break
		}
	}
	if i != LengthMin {
		t.Error("混合测试不通过，树不平衡")
	} else {
		t.Log("混合测试通过")
	}
}

func Test_Insert_only(t *testing.T) {
	var i int
	avlTree = MakeEmpty()
	for i = 1; i < Length; i++ {
		avlTree = avlTree.Insert(rand.Intn(Length))
		if !IsBalanced(avlTree) {
			break
		}
	}
	if i != Length {
		t.Error("插入测试不通过，树不平衡")
	} else {
		t.Log("插入测试通过")
	}
}

func Test_Delete_only(t *testing.T) {
	var i int
	for i = 1; i < Length; i++ {
		avlTree = avlTree.Delete(rand.Intn(Length))
		if !IsBalanced(avlTree) {
			break
		}
	}
	if i != Length {
		t.Error("删除测试不通过，树不平衡")
	} else {
		t.Log("删除测试通过")
	}
}

func Test_Mixture(t *testing.T) {
	var i int
	avlTree = MakeEmpty()
	for i = 1; i < Length; i++ {
		if rand.Intn(2) != 1 {
			avlTree = avlTree.Insert(rand.Intn(Length))
		} else {
			avlTree = avlTree.Delete(rand.Intn(Length))
		}
		if !IsBalanced(avlTree) {
			break
		}
	}
	if i != Length {
		t.Error("混合测试不通过，树不平衡")
	} else {
		t.Log("混合测试通过")
	}
}
