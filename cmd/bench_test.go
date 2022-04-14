package main

import (
	"benching/db"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	b.StopTimer()

	benchDB, err := db.NewDB("test.db")
	if err != nil {
		b.Fatal(err)
	}
	defer benchDB.Close()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err = benchDB.AddNewMember()
		if err != nil {
			b.Fatal(err)
		}
	}

	b.StopTimer()
	benchDB.Clear()
	b.StartTimer()

}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()

	benchDB, err := db.NewDB("test.db")
	if err != nil {
		b.Fatal(err)
	}
	defer benchDB.Close()

	for i := 0; i < 100; i++ {
		err = benchDB.AddNewMember()
		if err != nil {
			b.Fatal(err)
		}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := benchDB.GetAllMember()
		if err != nil {
			b.Fatal(err)
		}
	}

	b.StopTimer()
	benchDB.Clear()
	b.StartTimer()

}

func BenchmarkClear(b *testing.B) {
	b.StopTimer()

	benchDB, err := db.NewDB("test.db")
	if err != nil {
		b.Fatal(err)
	}
	defer benchDB.Close()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			err = benchDB.AddNewMember()
			if err != nil {
				b.Fatal(err)
			}
		}

		benchDB.Clear()
	}

}
