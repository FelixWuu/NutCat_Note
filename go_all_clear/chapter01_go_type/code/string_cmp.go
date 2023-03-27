package code

import (
	"fmt"
	"time"
)

func GetStringComparisonSpeed() {
	// 生成3个字符串，实际上它们的值相等。
	// str1 和 str2 都是对 byteStr 的 deep-copy，实际上它们的底层自己序列不同
	// str3 与 str2 共享同一个底层字节序列
	byteStr := make([]byte, 1<<26)
	str1 := string(byteStr)
	str2 := string(byteStr)
	str3 := str1

	// 这里比较 str1 和 str2. 由于底层序列不同，对比需要逐个比对字节，复杂度为 O(n)
	startTime := time.Now()
	_ = str1 == str2
	duration := time.Now().Sub(startTime)
	fmt.Println("duration for (str1 == str2):", duration)

	// 这里比较 str2 和 str3. 由于共享底层序列，只需要比较底层引用着字符串切片的指针是否相等，复杂度为 O(1)
	startTime = time.Now()
	_ = str1 == str3
	duration = time.Now().Sub(startTime)
	fmt.Println("duration for (str1 == str3):", duration)
}
