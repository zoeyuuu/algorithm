package main

/*
小美定义一个字符串是平衡串，当且仅当该字符串仅包含两种字符，且两种字符出现的次数相等。例如"ababba"是平衡串。
现在小美拿到了一个仅由小写字母组成的字符串，她想知道该字符串有多少子序列是平衡串？
定义子序列为从左到右取若干个字符（可以不连续）组成的字符串。例如，"aca"是"arcaea"的子序列。
输入描述：
第一行输入一个正整数n，代表字符串长度。
第二行输入一个长度为n的、仅由小写字母组成的字符串。
1\leq n \leq 200000
输出描述：
输出一个整数表示答案，由于答案可能很大，请输出答案对 10^9 + 7 取模的结果。
输入示例：
5
ababc
输出示例：
9
说明：
长度为 2 的子序列，共有 8 个。
长度为 4 的子序列，共有 1 个。
*/
