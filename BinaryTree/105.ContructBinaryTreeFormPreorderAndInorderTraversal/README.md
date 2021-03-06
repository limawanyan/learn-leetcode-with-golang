# 题目
给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。

 

    示例 1:
    
    
    Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
    Output: [3,9,20,null,null,15,7]
    示例 2:
    
    Input: preorder = [-1], inorder = [-1]
    Output: [-1]


    提示:
    
    1 <= preorder.length <= 3000
    inorder.length == preorder.length
    -3000 <= preorder[i], inorder[i] <= 3000
    preorder 和 inorder 均无重复元素
    inorder 均出现在 preorder
    preorder 保证为二叉树的前序遍历序列
    inorder 保证为二叉树的中序遍历序列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# 解题方法
    前序：根 ->  左 -> 右
    中序：左 ->  根 -> 右
    根据前序和中序可以有如下推导：
    前序：[3,9,20,15,7]
    中序：[9,3,15,20,7]
    推导二叉树步骤：
        首先我们根据前序遍历第一个元素3，可以得出二叉树的根为3,
        然后从中序中找到元素3,
        3的左边元素9,即为左子树,
        3的右边元素[15,20,7],即为右子树,
        所以我们可以得出左子树的前序遍历为[9],中序遍历为[9],
        右子树前序遍历为[20,15,7],中序数组遍历为[15,20,7],
        左子树只有一个元素[9],我们可以得出跟节点为3,根左节点就为9
    我们继续分解右子树：
        根据前序遍历数组[20,15,7],可以得出根为20
        从中序遍历[15,20,7]中找到20
        20左边[15]为左子树,右边[7]为右子树
    最后得出二叉树：   3
                    /  \
                   9   20
                      /  \
                    15    7

### 递归
    分析上面的推导过程，我们可以将问题进行拆解。
    根节点就是前序遍历的第一个元素
    定义一个索引变量 root = 中序遍历中根节点的索引
    递归遍历左右子树：
        左子树：前序遍历数组 = 前序遍历数组[1:root+1]
              中序遍历数组 = 中序遍历数组[0:root]
        右子树:前序遍历数组 = 前序遍历数组[root+1:]
              中序遍历数组 = 中序遍历数组[root+1:]
    
    时间复杂度O(n) n为树节点个数
    空间复杂度O(n)
