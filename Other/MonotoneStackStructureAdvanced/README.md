## 描述
给定一个可能含有重复值的数组 arr，找到每一个 i 位置左边和右边离 i 位置最近且值比 arr[i] 小的位置。返回所有位置相应的信息。
## 输入描述：
第一行输入一个数字 n，表示数组 arr 的长度。
以下一行输入 n 个数字，表示数组的值
## 输出描述：
输出n行，每行两个数字 L 和 R，如果不存在，则值为 -1，下标从 0 开始。

    示例1
    输入：
    7
    3 4 1 5 6 2 7

    输出：
    -1 2
    0 2
    -1 -1
    2 5
    3 5
    2 -1
    5 -1

    备注：
    1 <= n <= 1000000
    -1000000 <= arri <= 1000000

## 学习视频
    https://www.bilibili.com/video/BV1dP4y1t7gj?from=search&seid=7570905550042333830&spm_id_from=333.337.0.0