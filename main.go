package main

import (
	"fmt"
	"sync"
)

func main() {
	// 3. Hafta Örneği
	resultsChan := make(chan int)
	totalChan := make(chan int, 1)

	// toplayıcıyı baştan çalıştırıp gelen değerler `resultsChan` ve
	// çıktı değerini döneceği `totalChan` kanallarını tanımladık.
	go calculateTotal(resultsChan, totalChan)

	// waitgroup ile tüm değerlerin geldiğini teyit ediyoruz ve
	// bekliyoruz.
	wgCalculation := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wgCalculation.Add(2)
		go calculateSquare(i, resultsChan, wgCalculation)
		go calculateCube(i, resultsChan, wgCalculation)
	}
	wgCalculation.Wait()

	// resultChan'ı kapatara calculateTotal'e bu kanalın kapatıldığını
	// haber veriyoruz
	close(resultsChan)

	// totalChan beklenildiği için programı bitirmiyor
	// hiç data gelmemiş bir channelda çıktı vermeye çalışırsak orası
	// bekleme olarak çalışır
	fmt.Printf("Total Value : %d", <-totalChan)
}

func calculateTotal(resultsChan chan int, totalChan chan int) {
	total := 0
	for {
		select {
		// gelen değerleri kontrol et ve kanalın kapalı açık olmasını kontrol et
		// eğer ok(kanal açık) değilse(!ok) o zaman toplam değerini totalChan'a
		// çıktı kanalına dön ve işlemi bitir.
		case val, ok := <-resultsChan:
			if !ok {
				totalChan <- total
				return
			}
			total += val
		}
	}
}

func calculateSquare(n int, rChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rChan <- n * n
}

func calculateCube(n int, rChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rChan <- n * n * n
}
