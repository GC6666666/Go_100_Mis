package No_70

import (
	"sync"
	"testing"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func NewCache() *Cache {
	return &Cache{
		balances: make(map[string]float64),
	}
}

// 原始的有数据竞争的方法
func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances // 这里会产生数据竞争，因为在解锁后仍在使用map
	c.mu.RUnlock()

	sum := 0.0
	for _, b := range balances { // 在锁释放后访问map，可能与写入操作发生竞争
		sum += b
	}
	return sum / float64(len(balances))
}

// 修复后的方法
func (c *Cache) AverageBalanceSafe() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock() // 确保在方法返回前保持锁定

	if len(c.balances) == 0 {
		return 0
	}

	sum := 0.0
	for _, b := range c.balances {
		sum += b
	}
	return sum / float64(len(c.balances))
}
func (c *Cache) AverageBalanceSafe1() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, b := range m {
		sum += b
	}
	return sum / float64(len(m))
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balances[id] = balance
}

func TestCacheRace(t *testing.T) {
	// 使用 go test -race 运行此测试可以检测到数据竞争
	t.Run("demonstrate race condition", func(t *testing.T) {
		cache := NewCache()

		// 添加初始数据
		cache.AddBalance("user1", 100)
		cache.AddBalance("user2", 200)

		// 创建通道用于同步goroutine
		done := make(chan bool)

		// 启动写入goroutine
		go func() {
			for i := 0; i < 100; i++ {
				cache.AddBalance("user3", float64(i))
				time.Sleep(time.Microsecond) // 增加竞争机会
			}
			done <- true
		}()

		// 启动读取goroutine
		go func() {
			for i := 0; i < 100; i++ {
				_ = cache.AverageBalance() // 使用有竞争的方法
				time.Sleep(time.Microsecond)
			}
			done <- true
		}()

		// 等待两个goroutine完成
		<-done
		<-done
	})

	// 展示修复后的版本
	t.Run("demonstrate fixed version", func(t *testing.T) {
		cache := NewCache()

		// 添加初始数据
		cache.AddBalance("user1", 100)
		cache.AddBalance("user2", 200)

		done := make(chan bool)

		// 启动写入goroutine
		go func() {
			for i := 0; i < 100; i++ {
				cache.AddBalance("user3", float64(i))
				time.Sleep(time.Microsecond)
			}
			done <- true
		}()

		// 启动读取goroutine，使用安全版本
		go func() {
			for i := 0; i < 100; i++ {
				_ = cache.AverageBalanceSafe() // 使用修复后的方法
				time.Sleep(time.Microsecond)
			}
			done <- true
		}()

		// 等待两个goroutine完成
		<-done
		<-done
	})
}
