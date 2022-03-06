package types

type WinKind int

const (
	WinTop10 WinKind = iota
	WinTop5
	WinTop3
	WinTop1
)
