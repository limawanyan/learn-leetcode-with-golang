# 前缀和后缀搜索

设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。

实现 WordFilter 类：

WordFilter(string[] words) 使用词典中的单词 words 初始化对象。

f(string pref, string suff) 返回词典中具有前缀 prefix 和后缀 suff 的单词的下标。如果存在不止一个满足要求的下标，返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。


    示例：
    
    输入
    ["WordFilter", "f"]
    [[["apple"]], ["a", "e"]]
    输出
    [null, 0]
    解释
    WordFilter wordFilter = new WordFilter(["apple"]);
    wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词：前缀 prefix = "a" 且 后缀 suff = "e" 。

    提示：
    
    1 <= words.length <= 104
    1 <= words[i].length <= 7
    1 <= pref.length, suff.length <= 7
    words[i]、pref 和 suff 仅由小写英文字母组成
    最多对函数 f 执行 104 次调用

# 解题思路
### 哈希Map
    外部循环每个单词，内部对每个单词再进行双重循环，计算前缀后缀所有的组合可能性
    前缀后缀通过一个特殊字符进行拼接生成Key存储到Map中，将单词索引作为值存储。
    题中条件：在不止满足一个要求下标时返回其中最大的下标，循环后面的单词如果存在相同的则会覆盖之前的索引。

### 字典树