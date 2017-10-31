# kadt

个人数据结构库
具体示例请看对应的单元测试

----

##**1.kval**

节点结构

```
type avlNode struct {
        Element int
        Height  int
        Left    *avlNode
        Right   *avlNode
}

```

 - MakeEmpty
 - Find
 - Insert
 - Delete
 - PrintTree
 - IsBalanced
