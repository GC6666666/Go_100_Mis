package No_76

import (
	"fmt"
	"testing"
)

//func consumer(ch <-chan Event) {
//	timeduration := 1 * time.Hour
//	timer := time.NewTimer(timeduration)
//	for {
//		timer.Reset(timeduration)
//		select {
//		case: event := <-ch:
//			handle(event)
//		case <- ctx.Done():
//			log.Println("warning: no message received")
//		}
//	}
//}

//	func TestConsumer(t *testing.T) {
//		name1 := "gongchao"
//		name2 := "xiaoliu"
//		name3 := "nihao"
//		accounts := []account{
//			{
//				balance: 1,
//				user:    &name1,
//			},
//			{
//				balance: 2,
//				user:    &name2,
//			},
//		}
//		for i := range accounts {
//			fmt.Printf("%s ", *accounts[i].user)
//		}
//		for _, a := range accounts {
//			a.balance += 100
//			a.user = &name3
//		}
//		for i := range accounts {
//			fmt.Printf("%s ", *accounts[i].user)
//		}
//	}
//
//	type account struct {
//		balance int
//		user    *string
//	}
func TestRangeSlice(t *testing.T) {
	// 测试1：验证range复制行为
	s := []int{0, 1, 2}
	fmt.Printf("原始切片地址:%p 长度:%d 容量:%d\n", &s[0], len(s), cap(s))

	rangeCount := 0
	for range s {
		rangeCount++
		s = append(s, 10)
		fmt.Printf("第%d次迭代 切片地址:%p 长度:%d 容量:%d\n",
			rangeCount, &s[0], len(s), cap(s))
	}

	// 测试2：验证底层数组共享
	s2 := []int{0, 1, 2}
	originalAddr := fmt.Sprintf("%p", &s2[0])

	// 获取range使用的临时切片
	rangeSlice := s2
	rangeAddr := fmt.Sprintf("%p", &rangeSlice[0])

	fmt.Printf("原始切片底层数组地址:%s\n", originalAddr)
	fmt.Printf("range临时切片底层数组地址:%s\n", rangeAddr)

	// 测试3：验证append后的地址变化
	s2 = append(s2, 10)
	newAddr := fmt.Sprintf("%p", &s2[0])
	fmt.Printf("append后切片底层数组地址:%s\n", newAddr)
}
