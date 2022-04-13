package main

import (
	"sync"
	"testing"
	"time"
)

func TestGenerateTxs(t *testing.T) {
	GenerateMemKey()
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(THREAD)
	txBatch := [][]*Transaction{}
	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()
			txBatch = append(txBatch, GenerateTxs(COUNT))
		}(cpu)
	}
	wg.Wait()
	t.Logf("total generate tx count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
	start = time.Now()
	wg.Add(THREAD)
	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()
			txs := txBatch[c]
			//txs[0].Payload = []byte{123} //wrong data
			for _, tx := range txs {
				if err := VerifyTx(tx); err != nil {
					t.Fatal(err.Error())
				}
			}
		}(cpu)
	}
	wg.Wait()
	t.Logf("total verify tx count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
}
