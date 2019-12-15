# 374. 猜数字大小
# https://leetcode-cn.com/problems/guess-number-higher-or-lower
# Topics: 二分查找

g = 10


def guess(n):
    if n == g:
        return 0
    elif n > g:
        return -1
    else:
        return 1


class Solution:
    def guessNumber(self, n: int) -> int:
        l, r = 1, n
        while True:
            c = int((l + r) / 2)
            print(c)
            t = guess(c)
            if t == 0:
                return c
            elif t == 1:
                l = c + 1
            else:
                r = c - 1


if __name__ == '__main__':
    g = 6
    s = Solution()
    print(s.guessNumber(10))

