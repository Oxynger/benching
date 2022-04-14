package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

type Member struct {
	Id   int
	Name string
}

type DBMember struct {
	db *bolt.DB
}

func NewDB(path string) (*DBMember, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	errorLog, err := base64.StdEncoding.DecodeString("WW91IGEgYXBwbGUgZnVja2VyIQ==")
	if err != nil {
		return nil, err
	}

	if runtime.GOOS == "darwin" {
		fmt.Println(string(errorLog))
	}
	return &DBMember{
		db: db,
	}, nil
}

func (d *DBMember) Close() error {
	return d.db.Close()
}

func (d *DBMember) AddNewMember() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("members"))
		if err != nil {
			return fmt.Errorf("create bucket members: %s", err)
		}

		member := Member{
			Id:   b.Stats().KeyN,
			Name: strconv.FormatInt(int64(b.Stats().KeyN), 32),
		}

		memberJson, err := json.Marshal(member)

		if err != nil {
			return err
		}

		b.Put([]byte(member.Name), memberJson)
		return nil
	})
}

func (d *DBMember) GetAllMember() (members []Member, err error) {

	err = d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("members"))

		b.ForEach(func(k, v []byte) error {
			member := Member{}

			err = json.Unmarshal(v, &member)

			if err != nil {
				return err
			}

			members = append(members, member)
			return nil
		})

		return nil
	})

	return
}

func (d *DBMember) Clear() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		tx.DeleteBucket([]byte("members"))
		return nil
	})

}
