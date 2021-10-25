package ldb

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/spf13/cast"
	"github.com/syndtr/goleveldb/leveldb"
	leveldb_errors "github.com/syndtr/goleveldb/leveldb/errors"
)

var ldb_conn *DB

type Config struct {
	Datadir         string
	ErrorOnNotFound bool
}

type DB struct {
	conn   *leveldb.DB
	config Config
}

func NewLdb(config *Config) (*DB, error) {
	if ldb_conn != nil {
		return ldb_conn, nil
	}

	conn, err := leveldb.OpenFile(config.Datadir, nil)

	db := &DB{
		conn: conn,
	}

	ldb_conn = db

	return db, err
}

func (this *DB) Close() {
	if this.conn != nil {
		if err := this.conn.Close(); err != nil {
			log.Println(err)
		}
		this.conn = nil
	}
}

func (this *DB) GetObj(key string, value interface{}) error {
	v, err := this.conn.Get([]byte(key), nil)

	if err != nil {
		if err == leveldb_errors.ErrNotFound && !this.config.ErrorOnNotFound {
			return nil
		}

		return err
	}

	if v != nil {
		d := gob.NewDecoder(bytes.NewReader(v))
		return d.Decode(value)
	}

	return nil
}

func (this *DB) SetObj(key string, value interface{}) error {

	if value == nil {
		return leveldb_errors.ErrNotFound
	}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}
	return this.conn.Put([]byte(key), buf.Bytes(), nil)
}

func (this *DB) Get(key string) (string, error) {
	value, err := this.conn.Get([]byte(key), nil)

	if err != nil {
		if err == leveldb_errors.ErrNotFound && !this.config.ErrorOnNotFound {
			return "", nil
		}
		return "", err
	}

	if value != nil {
		return string(value), nil
	}

	return "", nil
}

func (this *DB) GetInt(key string) (int, error) {
	value, err := this.Get(key)
	if err != nil {
		return 0, err
	}

	return cast.ToInt(value), nil
}

func (this *DB) GetInt64(key string) (int64, error) {
	value, err := this.Get(key)
	if err != nil {
		return 0, err
	}

	return cast.ToInt64(value), nil
}

func (this *DB) Set(key string, value string) error {
	err := this.conn.Put([]byte(key), []byte(value), nil)
	return err
}

func (this *DB) SetInt64(key string, value int64) error {
	return this.Set(key, cast.ToString(value))
}

func (this *DB) SetInt(key string, value int) error {
	return this.Set(key, cast.ToString(value))
}
