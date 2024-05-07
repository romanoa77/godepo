package fbufstat

import "encoding/json"

type Bufstat struct {
	N_itm     uint   `json:"n_itm"`
	Buff_size uint32 `json:"buff_size"`
}

func New(nf uint, bsize uint32) Bufstat {

	Ftable := Bufstat{nf, bsize}

	return Ftable

}

func NewObj(buf []byte) Bufstat {

	var Ftable Bufstat

	json.Unmarshal(buf, &Ftable)

	return Ftable

}

func (Class Bufstat) GetObj() Bufstat {

	return Class

}

func (Class *Bufstat) UpdateCnt() {

	Class.N_itm += 1

}

func (Class Bufstat) GetJSONObj() []byte {

	bt_arr, _ := json.Marshal(Class)

	return bt_arr

}

func (Class *Bufstat) UpdateSize(buf uint32) {

	Class.Buff_size += buf

}
