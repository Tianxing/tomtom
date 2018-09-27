package dao

import(
    "database/sql"
    "errors"
)

type Dao struct {
    DB *sql.DB
    Tx *sql.Tx
}

var ErrNoTx = errors.New("NOT BEGIN OR COMMITED TX")
var ErrTxBegan = errors.New("TX HAD BEGAN")

func (d *Dao)Begin() error {
    if d.Tx != nil {
        return ErrTxBegan
    }

    tx, err := d.DB.Begin()
    if err != nil {
        return err
    }

    d.Tx = tx
    return nil
}

func (d *Dao)Rollback() error {
    if d.Tx == nil {
        return ErrNoTx
    }

    err := d.Tx.Rollback()
    d.Tx = nil
    if err != nil {
        return err
    }
    return nil
}

func (d *Dao)Commit() error {
    if d.Tx == nil {
        return ErrNoTx
    }

    err := d.Tx.Commit()
    d.Tx = nil
    if err != nil {
        return err
    }
    return nil
}

func (d *Dao)Query(query string, args ...interface{}) (*sql.Rows, error) {

    if d.Tx != nil {
        return d.Tx.Query(query, args...)
    }
    
    return d.DB.Query(query, args...)
}

func (d *Dao)QueryRow(query string, args ...interface{}) (*sql.Row) {

    if d.Tx != nil {
        return d.Tx.QueryRow(query, args...)
    }
    
    return d.DB.QueryRow(query, args...)
}

func (d *Dao)Exec(query string, args ...interface{}) (sql.Result, error) {
    if (d.Tx != nil) {
        return d.Tx.Exec(query, args...)
    }
    return d.DB.Exec(query, args...)
}
