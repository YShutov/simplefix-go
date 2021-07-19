package session

import "errors"

var (
	ErrNotEnoughMessages = errors.New("not enough messages in storage")
	ErrInvalidBoundaries = errors.New("invalid boundaries")
	ErrInvalidSequence   = errors.New("unexpected sequence index")
)

type MessageStorage interface {
	Save(msg []byte, msgSeqNum int) error
	Messages(msgSeqNumFrom, msgSeqNumTo int) ([][]byte, error)
}
