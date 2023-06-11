package pubsub

import (
	"errors"
	"strings"
)

var ErrEOF = errors.New("end of File")

func GetMockStream() *Stream {
	return &Stream{0, allTweets}
}

type Tweet struct {
	Username string
	text     string
}
type Stream struct {
	pos    int64 // position
	tweets []Tweet
}

var allTweets = []Tweet{
	{
		Username: "user1",
		text:     "random text1-topic1",
	},
	{
		Username: "user2",
		text:     "random text2-topic1",
	},
	{
		Username: "user3",
		text:     "random text3-topic2",
	},
}

func (t *Tweet) IsTopic1() bool {
	isTopic1 := strings.Contains(t.text, "topic1")
	return isTopic1
}
func (stream *Stream) Next() (*Tweet, error) {
	n := int64(len(allTweets))
	var twt *Tweet
	if stream.pos < n {
		twt = &stream.tweets[stream.pos]
	} else {
		return &Tweet{}, ErrEOF
	}
	stream.pos += 1
	return twt, nil
}
