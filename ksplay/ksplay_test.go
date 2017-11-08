package ksplay

import (
	"testing"
//	"fmt"
)

var (
    splayTree = MakeEmpty()
)

func makeStrfromChan(ch chan string) string{
    targetStr := ""
    for{
	if len(ch) != 0{
	    targetStr += <-ch
	}else{
	    return targetStr
	}
    }
}

func Test_SplayInsert_only(t *testing.T) {
	sucStr := "12 11 10 6 2 1 nil nil 5 4 3 nil nil nil nil nil nil nil 13 nil nil "
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
	targetStr := makeStrfromChan(ch)
	if sucStr == targetStr {
		t.Log("插入测试通过")
	} else {
		t.Error("插入测试不通过")
	}
}

func Test_SplayFind_only(t *testing.T) {
     sucStr1 := "1 nil 12 10 2 nil 6 5 4 3 nil nil nil nil nil 11 nil nil 13 nil nil "
     sucStr2 := "6 1 nil 2 nil 5 4 3 nil nil nil nil 12 10 nil 11 nil nil 13 nil nil "
     ch := make(chan string, 100)
     _, splayTree := splayTree.Find(1)
     PrintTreePre(splayTree,ch)
     targetStr := makeStrfromChan(ch)
     if targetStr != sucStr1{
	t.Error("搜索测试不通过")
     }
     _, splayTree = splayTree.Find(6)
     PrintTreePre(splayTree,ch)
     targetStr = makeStrfromChan(ch)
     if targetStr != sucStr2{
	t.Error("搜索测试不通过")
     }
    t.Log("搜索测试通过")
}
