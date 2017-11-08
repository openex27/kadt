package ksplay

import (
	"testing"
	//"fmt"
)

var (
    splayTreeG = MakeEmpty()
)

func makeStrfromChan(splayTree *splayNode) string{
    ch := make(chan string, 100)
    PrintTreePre(splayTree,ch)
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
	targetStr := makeStrfromChan(splayTree)
	if sucStr == targetStr {
		t.Log("插入测试通过")
	} else {
		t.Error("插入测试不通过")
	}
	splayTreeG = splayTree
}

func Test_SplayFind_only(t *testing.T) {
     sucStr1 := "1 nil 12 10 2 nil 6 5 4 3 nil nil nil nil nil 11 nil nil 13 nil nil "
     sucStr2 := "6 1 nil 2 nil 5 4 3 nil nil nil nil 12 10 nil 11 nil nil 13 nil nil "
     
     empty := MakeEmpty()
     _,empty = empty.Find(1)

     _, splayTree := splayTreeG.Find(1)
     targetStr := makeStrfromChan(splayTree)
     if targetStr != sucStr1{
	t.Error("搜索测试不通过")
     }

     _, splayTree = splayTree.Find(6)
     targetStr = makeStrfromChan(splayTree)
     if targetStr != sucStr2{
	t.Error("搜索测试不通过")
     }

    t.Log("搜索测试通过")
    splayTreeG = splayTree
}

func Test_SplayDelete_only(t *testing.T) {
    sucStr1 := "3 2 1 nil nil nil 6 5 nil nil 12 10 nil 11 nil nil 13 nil nil "
    sucStr2 := "2 nil 3 nil 6 5 nil nil 12 10 nil 11 nil nil 13 nil nil "
    splayTree := splayTreeG.Delete(4)
    if makeStrfromChan(splayTree) != sucStr1{
	t.Error("删除测试不通过")
    }
    splayTree = splayTree.Delete(1)
    if makeStrfromChan(splayTree) != sucStr2{
	t.Error("删除测试不通过")
    }
    t.Log("删除测试通过")
    splayTreeG = splayTree
}
