package mto

type Chat struct{
	user string
	writable bool
	time int
}

type Message struct{
	user string
	msg string
}