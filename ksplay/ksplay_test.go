package ksplay

import (
	"testing"
)

func Test_SplayInsert_only(t *testing.T) {
	sucStr := "12 11 10 6 2 1 nil nil 5 4 3 nil nil nil nil nil nil nil 13 nil nil "
	splayTree := MakeEmpty()
	splayTree = splayTree.Insert(3)
	splayTree = splayTree.Insert(1)
	splayTree = splayTree.Insert(5)
	splayTree = splayTree.Insert(4)
	splayTree = splayTree.Insert(2)
	splayTree = splayTree.Insert(6)
	splayTree = splayTree.Insert(13)
	splayTree = splayTree.Insert(10)
	splayTree = splayTree.Insert(11)
	splayTree = splayTree.Insert(12)
	ch := make(chan string, 100)
	PrintTreePre(splayTree, ch)
	targetStr := ""
	for {
		if len(ch) != 0 {
			targetStr += <-ch
		} else {
			break
		}
	}
	if sucStr == targetStr {
		t.Log("插入测试通过")
	} else {
		t.Error("插入测试不通过")
	}
}
