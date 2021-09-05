## 题目
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

    push(x) —— 将元素 x 推入栈中。
    pop() —— 删除栈顶的元素。
    top() —— 获取栈顶元素。
    getMin() —— 检索栈中的最小元素。
    
    
    示例:
    
    输入：
    ["MinStack","push","push","push","getMin","pop","top","getMin"]
    [[],[-2],[0],[-3],[],[],[],[]]
    
    输出：
    [null,null,null,null,-3,null,0,-2]
    
    解释：
    MinStack minStack = new MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    minStack.getMin();   --> 返回 -3.
    minStack.pop();
    minStack.top();      --> 返回 0.
    minStack.getMin();   --> 返回 -2.


提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 解题方法
### 辅助栈
    时间复杂度:O(1)
    空间复杂度:O(n)
    
    定义一个最小值辅助栈
    在push的时候同时push最小值辅助栈,
    最小值辅助栈当前元素与上一个元素比较,取小的为最新的辅助栈元素。

### 一个栈
    时间复杂度:O(1)
    空间复杂度:O(1)
    
    栈中存储自定义数据结构，结构体包含当前值和最小值,
    每次push,当前节点记录当前值和最小值,最小值拿当前值和上一个节点存储最小值比较。