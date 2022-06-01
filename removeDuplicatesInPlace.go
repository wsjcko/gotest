package main

import "sort"

func RemoveDuplicatesInPlace(userIds []int64) []int64 {
	if len(userIds) < 2 {
		return userIds
	}

	//排序
	sort.SliceStable(userIds, func(i, j int) bool { return userIds[i] < userIds[j] })
	uniqPointer := 0
	for i := 1; i < len(userIds); i++ {
		//比较元素，不相同，写入唯一指针的右侧
		if userIds[uniqPointer] != userIds[i] {
			uniqPointer++
			userIds[uniqPointer] = userIds[i]
		}
	}
	return userIds[:uniqPointer+1]
}
