
## 给定一个整数n，求由“0”字符和“1”字符组成的长度为n的所有字符串中，满足“0”字符的左边必有“1”字符的字符串的数量。
    链接：https://www.nowcoder.com/questionTerminal/5a04a774b9f54d9eb783a4d583e7a60a
    来源：牛客网

    输入描述:
    输入一行，包含一个整数n（1 \leq n \leq2*10^7）（1≤n≤2∗10
    7
    ）。

    输出描述:
    输出一个整数，表示返回的答案，由于字符串的数量巨大，可能会溢出，请输出对2^{29}2
    29
    取模后的答案。
    示例1
    输入
    1
    输出
    1
    说明
    只有“1”满足

    示例2
    输入
    2
    输出
    2
    说明
    只有“10”和“11”满足

    示例3
    输入
    3
    输出
    3
    说明
    只有“101”，“110”，“111”满足
    
    备注:
    时间复杂度O（log_2n）O（log2n）。额外空间复杂度O（1）O（1）。