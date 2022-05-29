package idgen

var ID int32

func Init() {
	ID = 0
}

func GetID() int32 {
	ID++
	return ID
}
