# 278. 第一个错误的版本
# https://leetcode-cn.com/problems/first-bad-version
# Topics: 二分查找

# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
failed = 3


def isBadVersion(version):
    if version >= failed:
        return True
    return False


class Solution:
    def firstBadVersion(self, n):
        l, r = 1, n
        while r > l:
            c = int((l + r) / 2)
            t = isBadVersion(c)
            if t:
                r = c
            else:
                l = c + 1
        return l


if __name__ == '__main__':
    s = Solution()
    failed = 4
    print(s.firstBadVersion(5))
