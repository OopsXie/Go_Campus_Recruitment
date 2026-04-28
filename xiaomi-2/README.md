# Xiaomi-2: 二叉树平衡性检查（AVL Tree Balance Check）

## 题目描述

给定一个二叉树的结构定义（通过左右子节点数组给出），判断这个二叉树是否为平衡二叉树。

**平衡二叉树定义**：
- 一棵空树是平衡的
- 对于任意节点，其左右子树的高度差不超过1
- 且左右子树都是平衡的

**输入格式**：
- 节点通过下标索引
- `left[i]` 表示节点i的左子树下标（-1表示无左子树）
- `right[i]` 表示节点i的右子树下标（-1表示无右子树）
- 从根节点（下标0）开始检查

## 解题思路

### 核心思想
使用递归进行深度优先搜索（DFS），自下而上地检查每个节点的平衡性。

### 算法设计

**递归函数**：`calculate(root, left, right, n) → (isBalance, height)`

返回值：
- `isBalance`: 当前子树是否平衡
- `height`: 当前子树的高度

**递归过程**：

1. **基础情况**：
   ```
   if root == -1:
       return (true, 0)  // 空树是平衡的，高度为0
   ```

2. **递归求解**：
   ```
   leftBalance, leftHeight = calculate(left[root], ...)
   rightBalance, rightHeight = calculate(right[root], ...)
   ```

3. **检查平衡性**：
   - 左子树是否平衡
   - 右子树是否平衡
   - 左右高度差是否 ≤ 1：`|leftHeight - rightHeight| ≤ 1`

4. **返回结果**：
   ```
   if 所有条件满足:
       return (true, max(leftHeight, rightHeight) + 1)
   else:
       return (false, 0)
   ```

### 时间复杂度
- **时间**: $O(n)$ - 每个节点访问一次
- **空间**: $O(h)$ - h是树的高度，递归栈深度

## 代码实现

```go
func calulate(root int, left []int, right []int, n int) (bool, int) {
    // 1. 处理空节点
    if root == -1 {
        return true, 0
    }
    
    // 2. 递归检查左右子树
    leftBalance, leftHeight := calulate(left[root], left, right, n)
    if !leftBalance {
        return false, 0
    }
    
    rightBalance, rightHeight := calulate(right[root], left, right, n)
    if !rightBalance {
        return false, 0
    }
    
    // 3. 检查高度差
    if math.Abs(float64(leftHeight-rightHeight)) > 1 {
        return false, 0
    }
    
    // 4. 返回当前树的高度
    return true, max(leftHeight, rightHeight) + 1
}
```

## 关键要点

1. **自下而上检查**：先检查子树再检查父树
2. **早期终止**：若发现不平衡立即返回false
3. **高度定义**：空树高度为0，叶节点高度为1
4. **平衡条件**：高度差 ≤ 1，且左右子树都平衡

## 学习价值

- 理解AVL树的平衡性定义
- 学习树的递归遍历和检查
- 掌握如何通过递归返回多个值
- 理解自下而上地解决树的问题
- 优化：避免遍历整个树（若不平衡可提前返回）

## 扩展思考

1. 如何检查是否是完全二叉树？
2. 如何计算树的直径（最长路径）？
3. 如何检查是否是二叉搜索树（BST）？
