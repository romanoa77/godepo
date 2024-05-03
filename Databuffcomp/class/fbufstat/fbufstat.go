package fbufstat

import "encoding/json"

type bufStat struct {
	N_itm     uint   `json:"n_itm"`
	Buff_size uint32 `json:"buff_size"`
}

func New(nf uint, bsize uint32) bufStat {

	Ftable := bufStat{nf, bsize}

	return Ftable

}

func (Class *bufStat) UpdateCnt() {

	Class.N_itm += 1

}

func (Class bufStat) GetJSONObj() []byte {

	bt_arr, _ := json.Marshal(Class)

	return bt_arr

}

func (Class *bufStat) UpdateSize(buf uint32) {

	Class.Buff_size += buf

}
