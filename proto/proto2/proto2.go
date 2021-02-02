package proto2

const (
	_                    = iota
	C2SPlayerLoginProto2 //用户登录请求
	S2CPlayerLoginProto2 //用户登录响应

	C2SChooseRoomProto2 //选择房间请求
	S2CChooseRoomProto2 //选择房间响应
)

type Player struct {
	UID        int
	PlayerName string
	OpenID     string
}

type C2SPlayerLogin struct {
	Protocol int    `json:"protocol"`
	Protoco2 int    `json:"protocol2"`
	Code     string `json:"code"`
}

type S2CPlayerLogin struct {
	Protocol   int
	Protoco2   int
	PlayerData *Player
}
