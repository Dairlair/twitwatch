package entity

type TweetStreamsInterface interface {
	GetTweet() TweetInterface
	GetStreams() []StreamInterface
}

type tweetStreams struct {
	tweet TweetInterface
	streams []StreamInterface
}

func NewTweetStreams (tweet TweetInterface, streams []StreamInterface) TweetStreamsInterface {
	return &tweetStreams{
		tweet: tweet,
		streams: streams,
	}
}

// GetID returns the Stream ID
func (ts *tweetStreams) GetTweet() TweetInterface {
	return ts.tweet
}

// GetTrack returns the stream's track
func (ts *tweetStreams) GetStreams() []StreamInterface {
	return ts.streams
}

func StreamsMapToSlice(streamsMap map[int64]StreamInterface) []StreamInterface {
	var streamsSlice []StreamInterface
	for _, value := range streamsMap {
		streamsSlice = append(streamsSlice, value)
	}
	return streamsSlice
}