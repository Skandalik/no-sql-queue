package queue

import jsoniter "github.com/json-iterator/go"


type Message struct {
	Text string
}

// Marshal returns byte-encoded JSON and error if occurs
func (m Message) Marshal() ([]byte, error) {
	return jsoniter.Marshal(m)
}