/*
    fileName: badger
    author: diogoxiang@qq.com
    date: 2019/7/7
*/
package main

import (
	// "bytes"
	"flag"
	"fmt"
	"github.com/dgraph-io/badger"
	"math/rand"
	"strconv"
	"time"
)

var (
	number    = *flag.Int("number", 100, "package")
	N         = *flag.Int("N", 10000, "for limit N")
	value     = RandStringBytesMaskImprSrc(512)
	valueByte = []byte(value)
	src       = rand.NewSource(time.Now().UnixNano())
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main()  {
		badgerDb()
	// benchmarkBadgerSet()
	// benchmarkBadgerGet()
}

/// RandStringBytesMaskImprSrc
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func badgerDb() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions("")
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("10"), []byte("100"))
	})

	db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte("10"))
		fmt.Println(item)
		return nil
	})
}


func benchmarkBadgerSet() {
	opts := badger.DefaultOptions("")
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	now := time.Now()
	for i := 0; i < N; i++ {
		db.Update(func(txn *badger.Txn) error {
			for i := 0; i < number; i++ {
				txn.Set([]byte(fmt.Sprintf("{}{}{}{}", rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())), valueByte)
			}
			return nil
		})
	}
	fmt.Println("benchmarkBadgerSet cast time: ", time.Now().Sub(now))
}


func benchmarkBadgerGet() {
	opts := badger.DefaultOptions("")
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < number; i++ {
			txn.Set([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	now := time.Now()
	for i := 0; i < N; i++ {
		db.View(func(txn *badger.Txn) error {
			item, _err := txn.Get([]byte(strconv.Itoa(rand.Intn(number))))
			if _err != nil {
				panic("BenchmarkBadgerGet error")
			} else {

				fmt.Println(item)
				// if val, err := item.Value(); err == nil {
				// 	if !bytes.Equal(valueByte, val) {
				// 		panic("BenchmarkBadgerGet error")
				// 	}
				// } else {
				// 	panic("BenchmarkBadgerGet error")
				// }
			}
			return nil
		})
	}
	fmt.Println("benchmarkBadgerGet cast time: ", time.Now().Sub(now))
}