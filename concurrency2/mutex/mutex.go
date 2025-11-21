package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func Mutex() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	orders := generateOrders(20)
	go func() {
		defer wg.Done()
		processOrders(orders)
	}()
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatuses(order)
			}
		}()
	}
	wg.Wait()
	reportOrderStatus(orders)
	fmt.Printf("All operations completed. Updated %d orders. Exiting.\n", totalUpdates)
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}
	return orders
}

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func reportOrderStatus(orders []*Order) {
	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %s\n",
			order.ID, order.Status,
		)
	}
}

func updateOrderStatuses(order *Order) {
	order.mu.Lock()
	defer order.mu.Unlock()
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf("Updated order %d status: %s\n", order.ID, status)

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}
