package myslice

type MySlice struct {
	S []int
}

func CreateNewSlice(ss []int) MySlice {
	return MySlice{S: ss}
}
