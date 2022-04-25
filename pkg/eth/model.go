package eth

import (
	"database/sql/driver"
	"encoding/json"
)

type Block struct {
	BlockNum     uint64         `json:"block_num" gorm:"primaryKey,index"`
	BlockHash    string         `json:"block_hash" gorm:"index"`
	BlockTime    uint64         `json:"block_time"`
	ParentHash   string         `json:"parent_hash"`
	Transactions *BlockTxsHashs `json:"transactions,omitempty" gorm:"type:text"`
}

type BlockTxsHashs []string

func (bh *BlockTxsHashs) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), bh)
}

func (bh *BlockTxsHashs) Value() (driver.Value, error) {
	val, err := json.Marshal(bh)
	return string(val), err
}

type Transaction struct {
	TxHash string `json:"tx_hash" gorm:"primaryKey"`
	From   string `json:"from" gorm:"column:from_acc"`
	To     string `json:"to" gorm:"column:to_acc"`
	Nonce  uint64 `json:"nonce"`
	Data   string `json:"data"`
	Value  int64  `json:"value"`
	Logs   *Logs  `json:"logs" gorm:"column:logs;type:text"`
}

type Logs []*Log

func (l *Logs) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), l)
}

func (l *Logs) Value() (driver.Value, error) {
	val, err := json.Marshal(l)
	return string(val), err
}

type Log struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
