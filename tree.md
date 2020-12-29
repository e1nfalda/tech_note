# 数据结构树 
:Tech:algorithm:

## 常用的树
* `无序树`：树中任意节点的子节点之间没有顺序关系，这种树称为无序树，也称为[自由树](https://zh.wikipedia.org/wiki/自由树)；
* `有序树`：树中任意节点的子节点之间有顺序关系，这种树称为有序树；
  * `二叉树`
    每个节点最多含有两个子树的树称为二叉树；
    * `完全二叉树`:对于一颗二叉树，假设其深度为d（d>1）。除了第d层外，其它各层的节点数目均已达最大值，且第d层所有节点从左向右连续地紧密排列，这样的二叉树被称为完全二叉树；
      *` 满二叉树`: 所有叶节点都在最底层的完全二叉树；
    * `平衡二叉树`: 当且仅当任何节点的两棵子树的高度差不大于1的二叉树；
    * `二叉查找树(BST)`：也称二叉搜索树、有序二叉树；
  * `N叉树(N-ary Tree)`
  * `霍夫曼树`：(带权路径)最短的二叉树称为哈夫曼树或最优二叉树；
  * `B树 `:一种对读写操作进行优化的自平衡的二叉查找树，能够保持数据有序，拥有多于两个子树。
  * `平衡树`
    * AVL树
    * 红黑树
    * 2-3树

## 重点
* 遍历方法及特点。
  * depth-first
    * pre-oder
    * in-order
    * post-order

----

[平衡树wiki]: https://zh.wikipedia.org/wiki/%E5%B9%B3%E8%A1%A1%E6%A0%91
[二叉搜索树wiki]: https://zh.wikipedia.org/wiki/排序二元樹
[满二叉树]: https://zh.wikipedia.org/wiki/满二叉树
[霍夫曼树]: https://zh.wikipedia.org/wiki/霍夫曼树
[B树]: https://zh.wikipedia.org/wiki/B树
[常用数据结构总结]: https://www.freecodecamp.org/news/the-top-data-structures-you-should-know-for-your-next-coding-interview-36af0831f5e3/
