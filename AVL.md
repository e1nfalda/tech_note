# AVL

<!-- :algorithm: -->

## 定义

1. 左右子树的高度差不大于1.
2. 每一个子树都是平衡二叉树

## 实现

### 平衡因子(Balance Factor)

**Height**: `max(root.left.height, root.right.height)`
`BalanceFactor = root.left.height - root.right.height `有正负之分.

### 平衡操作(tree rebalance)

`左右旋转操作`
`LL`,`RR`: 做一次左旋或者右旋即可.
`LR`(`RL`反过来): 添加节点到左子树的右边. 先进行左旋,然后再进行一次右旋.

### 增加

1. 添加节点.

2. 调用rebalance方法
   
   ### 删除

3. 删除节点.

4. 调用rebalance方法.

---

[AVL树介绍及算法pdf](../articles/algorithm/AVL 树.pdf)
noteable 里笔记.
