package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var msg = "HelloWorld"
var THREAD = runtime.NumCPU()
var COUNT = 100000

type DataAndSign struct {
	Data []byte
	Sign []byte
}

func TestP256(t *testing.T) {

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.Nil(t, err)

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(THREAD)
	txBatch := []*DataAndSign{}
	lock := sync.Mutex{}
	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()
			for i := 0; i < COUNT; i++ {
				hash := sha256.Sum256(Uint32ToBytes(uint32(c*COUNT + i)))
				sign, _ := ecdsa.SignASN1(rand.Reader, priv, hash[:])
				lock.Lock()
				txBatch = append(txBatch, &DataAndSign{Data: Uint32ToBytes(uint32(c*COUNT + i)), Sign: sign})
				lock.Unlock()
			}
		}(cpu)
	}
	wg.Wait()
	t.Logf("ECDSA total generate and sign tx count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())

	pub := &priv.PublicKey

	wg = sync.WaitGroup{}
	start = time.Now()
	wg.Add(THREAD)
	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()

			for i := 0; i < COUNT; i++ {
				data := txBatch[c*COUNT+i]
				hash := sha256.Sum256([]byte(data.Data))
				b := ecdsa.VerifyASN1(pub, hash[:], data.Sign)
				if !b {
					t.Fatal(fmt.Sprintf("index=%d, data:%x,sign:%x", c*COUNT+i, data.Data, data.Sign))
				}
			}
		}(cpu)
	}
	wg.Wait()
	t.Logf("ECDSA total verify tx count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
}

func TestEd25519(t *testing.T) {
	message := []byte(msg)
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	signature := ed25519.Sign(priv, message)

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(THREAD)

	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()
			for i := 0; i < COUNT; i++ {
				b := ed25519.Verify(pub, message, signature)
				if !b {
					t.Fatal("Verify fail")
				}
			}
			//t.Logf("CPU[%d] verify count=%d, cost:%v", c, COUNT, time.Since(start))
		}(cpu)
	}
	wg.Wait()
	t.Logf("ED25519 total verify count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
}
