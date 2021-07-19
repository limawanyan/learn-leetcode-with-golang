# 题目
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]

输出：[[-1,-1,2],[-1,0,1]]

示例 2：

输入：nums = []

输出：[]

示例 3：

输入：nums = [0]

输出：[]


提示：

0 <= nums.length <= 3000

-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# 解题思路
### 1. 暴力破解
    三重循环，时间复杂度O(n^3)

### 2. 双重循环 + map
    用 map 提前计算好任意 2 个数字之和，保存起来，可以将时间复杂度降到 O(n^2)
    两层循环 a 和 b
    由于 c = -a-b    
    所以直接去set集合查找是否存在即可
    时间复杂度 O(n^2) + O(1)

### 3. 排序 + 双指针
    先将数组元素进行排序(便于跳过重复元素，如果当前元素和前一个元素相同则跳过)
    从第一个元素进行循环
    设置两个指针（左右指针向内收缩）：
        左指针 = 循环元素后第一个元素
        右指针 = 数组最后一个元素
    如果循环元素 + 左指针元素 + 右指针元素：
        大于0:
            右指针向左移动一位(因为结果大于0，而数组已经是有序的，所以要减小大的值)
        小于0：
            左指针向右移动一位
        两指针相遇：
            结束，不存在
    优化：如果循环的元素大于0，则排序后的数组另外两个数肯定也大于0，sum不可能等于0，所以直接break    

    https://leetcode-cn.com/problems/3sum/solution/zhi-zhen-yi-dong-guo-cheng-zhong-tiao-guo-zhong-fu/