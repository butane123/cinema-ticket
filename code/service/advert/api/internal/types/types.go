// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	IsCom   int64  `json:"isCom"`
	Status  int64  `json:"status"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type UpdateRequest struct {
	Id      int64  `json:"id"`
	Title   string `json:"title,optional"`
	Content string `json:"content,optional"`
	IsCom   int64  `json:"isCom,optional"`
	Status  int64  `json:"status,optional"`
}

type UpdateResponse struct {
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	IsCom   int64  `json:"isCom"`
	Status  int64  `json:"status"`
}

type CommerListResponse struct {
	List []*CommerAdvert `json:"list"`
}

type CommerAdvert struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NormalListResponse struct {
	List []*NormalAdvert `json:"list"`
}

type NormalAdvert struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
