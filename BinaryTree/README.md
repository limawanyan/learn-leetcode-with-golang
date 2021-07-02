# 二叉树
### 前序遍历
    先访问根节点，再前序遍历左子树，再前序遍历右子树
### 中序遍历
    先中序左子树，在访问根节点，再中序遍历右子树
### 后序遍历
    先后序遍历左子树，再后序遍历右子树，再访问跟节点
### 层次遍历(DFS、BFS)
    按层次遍历

# 分治法应用
* 快速排序
* 归并排序
* 二叉树相关问题
  
      先处理局部，再合并结果 
      分治法模板：
        1、递归返回条件
        2、分段处理
        3、合并结果

      func traversal(root *TreeNode) ResultType  {
          // nil or leaf
          if root == nil {
          // do something and return
          }
        
          // Divide
          ResultType left = traversal(root.Left)
          ResultType right = traversal(root.Right)
        
          // Conquer
          ResultType result = Merge from left and right
        
          return result
      }

## 归并排序（MergeSort）
    