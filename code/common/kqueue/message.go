package kqueue

// 下单信息消息
type CreateOrderMessage struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Fid    int64 `json:"fid"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}
