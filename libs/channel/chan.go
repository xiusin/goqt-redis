package channel

const ADD = 0

const SUB = 1

type UpdateKeyChStruct struct {
	ServeId  int
	ServeIdx int
	DbIdx    int
	Action   int
	Key      string
}

var UpdateKeyCh = make(chan UpdateKeyChStruct)
