package models

//首字母大写表示公有权限
//type User struct {
//	Uid int32 `json:"uid"` //转成json格式时，Uid会变成uid
//}

type Envelope struct {
	EnvelopeId int64 `json:"envelope_id"` // 红包id
	Uid        int64 `json:"uid"`         //拥有者
	Value      int32 `json:"value"`       //红包金额（分）
	MaxCount   int32 `json:"max_count"`   // 最多抢几次
	CurMax     int32 `json:"cur_count"`   // 当前为第几次抢
	Opened     bool  `json:"opened"`      //是否打开
	SnatchTime int32 `json:"snatch_time"` //抢到时间
	OpenedTime int32 `json:"opened_time"` //打开时间
}

//type Response struct {
//	Code int32  `json:"code"` //状态码
//	Msg  string `json:"msg"`  //信息
//	Data gin.H  `json:"data"` //数据内容
//}
