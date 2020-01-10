# 690. 员工的重要性
# https://leetcode-cn.com/problems/employee-importance
# Topics: 深度优先搜索 广度优先搜索 哈希表

# Employee info
class Employee:
    def __init__(self, id, importance, subordinates):
        # It's the unique id of each node.
        # unique id of this employee
        self.id = id
        # the importance value of this employee
        self.importance = importance
        # the id of direct subordinates
        self.subordinates = subordinates


class Solution:
    def getImportance(self, employees, id):
        """
        :type employees: Employee
        :type id: int
        :rtype: int
        """
        m = {}
        for employee in employees:
            m[employee.id] = employee

        now = [id]
        res = 0
        while len(now) > 0:
            next = []
            for i in now:
                employee = m[i]
                res += employee.importance
                next.extend(employee.subordinates)
            now = next
        return res


if __name__ == '__main__':
    employees = [
        Employee(1, 5, [2, 3]),
        Employee(2, 3, []),
        Employee(3, 3, []),
    ]
    s = Solution()
    print(s.getImportance(employees, 1))
    print(s.getImportance(employees, 2))
