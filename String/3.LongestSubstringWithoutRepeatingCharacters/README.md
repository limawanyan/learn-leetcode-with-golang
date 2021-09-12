# 题目
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

 

    示例 1:
    
    输入: s = "abcabcbb"
    输出: 3
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
    示例 2:
    
    输入: s = "bbbbb"
    输出: 1
    解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
    示例 3:
    
    输入: s = "pwwkew"
    输出: 3
    解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
    示例 4:
    
    输入: s = ""
    输出: 0
    
    
    提示：
    
    0 <= s.length <= 5 * 104
    s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

# 解题思路
### 暴力破解

### 滑动窗口
    1. 只增不减滑动窗口
    两个指针start和end表示窗口大小，遍历一次字符串，窗口在遍历过程中滑动或增大
    窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
    窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符的后一位
    遍历结束，返回end-start，窗口大小
    
    注解： 
    例：“wabcdwabf” 最长子串为 6
        初始化两指针为0
        前面“wabcd”无重复,直到 end = 5
        索引为5的w在滑动窗口索引0处重复
        start = 0 + 0 + 1 = 1
        end = 5 + 0 + 1 = 6
        继续循环,索引为6的“a”元素
        a元素在“abcdw”窗口的0索引处重复
        start = 1 + 0 + 1 = 2
        end = 6 + 0 + 1 = 7
        继续循环,索引为7的“b”元素
        b元素在“bcdwa”窗口的0索引处重复
        start = 2 + 0 + 1 = 3
        end = 7 + 0 + 1 = 8
        继续循环,索引为8的“f”元素
        f元素在“cdwab”窗口未出现重复
        判断循环索引+1是否大于end，大于则说明遍历到窗口外，窗口需要增大
        所以 end++ = 9
        最终返回结果 end - start = 9 - 3 = 6

题解来源：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/zhi-zeng-da-bu-jian-xiao-de-hua-dong-chuang-kou-10/

### 动态规划
    