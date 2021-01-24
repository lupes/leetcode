package question_1411_1420

import (
	"sort"
	"strconv"
)

// 1418. 点菜展示表
// https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant/
// Topics: 哈希表

func displayTable(orders [][]string) [][]string {
	var foodMap = make(map[string]int)
	var tables = make([]int, 0)
	var tableMap = make(map[string]map[string]int)
	for _, order := range orders {
		foodMap[order[2]] = -1
		table, ok := tableMap[order[1]]
		if !ok {
			table = make(map[string]int)
			number, _ := strconv.Atoi(order[1])
			tables = append(tables, number)
		}
		table[order[2]]++
		tableMap[order[1]] = table
	}
	sort.Ints(tables)

	var foods = make([]string, 0, len(foodMap))
	for food, _ := range foodMap {
		foods = append(foods, food)
	}
	sort.Strings(foods)

	var res = [][]string{{"Table"}}
	for i, food := range foods {
		res[0] = append(res[0], food)
		foodMap[food] = i + 1
	}

	for _, number := range tables {
		table := make([]string, len(foods)+1)
		table[0] = strconv.Itoa(number)
		for i := 1; i <= len(foods); i++ {
			table[i] = "0"
		}
		for food, num := range tableMap[table[0]] {
			table[foodMap[food]] = strconv.Itoa(num)
		}
		res = append(res, table)
	}

	return res
}
