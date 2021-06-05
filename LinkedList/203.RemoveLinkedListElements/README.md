# 题目
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

    示例 1：
    输入：head = [1,2,6,3,4,5,6], val = 6
    输出：[1,2,3,4,5]
    示例 2：
    输入：head = [], val = 1
    输出：[]
    示例 3：
    输入：head = [7,7,7,7], val = 7
    输出：[]

提示：

列表中的节点在范围 [0, 104] 内

1 <= Node.val <= 50

0 <= k <= 50

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-linked-list-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# 解题思路
### 递归
    递归终止条件:head == nil
    递归 head.next
    返回值:如果head == val 返回 head.next
          如果head != val 返回head

### 迭代
    由于第一个节点可能会被删除，所以需要创建一个哑结点，dummyHead.next = head
    初始化 temp = dummyHead
    迭代temp,如果下一个元素等于val,通过temp.next = temp.next.next删除节点
    如果下一个元素不等于val,保留下一个节点,移动temp指针到下一个元素