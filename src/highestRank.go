package main

func HighestRank(nums []int) int {
  numMap := make(map[int]int, 0)
  result := 0

  for _, value := range nums {
    if _, ok := numMap[value]; !ok {
      numMap[value] = 1
    } else {
      numMap[value] += 1
    }
  }

  numSlice := make([]int, 0, 0)
  mapVal := 0
  for _, val := range numMap {
    if val > mapVal {
      mapVal = val
    }
  }

  for k, v := range numMap {
    if v == mapVal {
      numSlice = append(numSlice, k)
    }
  }

  for _, res := range numSlice {
    if res > result {
      result = res
    }
  }

  return result
}
