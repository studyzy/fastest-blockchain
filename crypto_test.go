package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var msg = "HelloWorld"
var THREAD = runtime.NumCPU()
var COUNT = 100000

func TestSignVerify(t *testing.T) {

}

func TestP256(t *testing.T) {
	h := sha256.Sum256([]byte(msg))
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.Nil(t, err)

	sign, err := priv.Sign(rand.Reader, h[:], nil)
	require.Nil(t, err)
	require.NotEqual(t, nil, sign)

	pub := &priv.PublicKey

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(THREAD)

	for cpu := 0; cpu < THREAD; cpu++ {
		go func(c int) {
			defer wg.Done()
			for i := 0; i < COUNT; i++ {
				hash := sha256.Sum256([]byte(msg))
				b := ecdsa.VerifyASN1(pub, hash[:], sign)
				//require.Nil(t, err)
				require.True(t, b)
			}
			t.Logf("CPU[%d] verify count=%d, cost:%v", c, COUNT, time.Since(start))
		}(cpu)
	}
	wg.Wait()
	t.Logf("total verify count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
}

func TestEd25519Verification(t *testing.T) {
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
			t.Logf("CPU[%d] verify count=%d, cost:%v", c, COUNT, time.Since(start))
		}(cpu)
	}
	wg.Wait()
	t.Logf("total verify count=%d, cost:%v, TPS:%v", COUNT*THREAD, time.Since(start),
		float64(COUNT*THREAD)/time.Since(start).Seconds())
}
