# kadt

个人数据结构库
具体示例请看对应的单元测试

----
## 二叉搜索树
### **kavl -- AVL树**

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

### **ksplay -- splay树**

节点结构

```
type splayNode struct {
        Element int
        Left    *splayNode
        Right   *splayNode
}
```

 - MakeEmpty
 - Find
 - Insert
 - Delete
 - PrintTreePre
