package message

const (
	DQueryForAcc MessageType = "DQueryForAcc"
	DReplyToAcc  MessageType = "DReplyToAcc"
)

type QueryACC struct {
	Address string
}

type ReplyToAcc struct {
	Balance string
	Address string
}
