package networking

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// CHANNEL_SIZE  is the number of incoming messages to buffer for each topic.
const CHANNEL_SIZE = 128

// Subscription represents a subscription to a single Subscription topic. Messages
// can be published to the topic with Subscription.Publish, and received
// messages are pushed to the Messages channel.
type Subscription struct {
	// Messages is a channel of messages received from other peers in the chat room
	Messages chan *publishMessage

	ctx    context.Context
	pubSub *pubsub.PubSub
	topic  *pubsub.Topic
	sub    *pubsub.Subscription

	TopicName string
	self      peer.ID
}

// publishMessage gets converted to/from JSON and sent in the body of pubsub messages.
type publishMessage struct {
	Message  string
	SenderID string
}

// subscribeToTopic tries to subscribe to the Subscription topic for the room name, returning
// a Subscription on success.
func subscribeToTopic(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, topicName TopicType) (*Subscription, error) {
	// join the pubsub topic
	topic, err := ps.Join(string(topicName))
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	cr := &Subscription{
		ctx:       ctx,
		pubSub:    ps,
		topic:     topic,
		sub:       sub,
		self:      selfID,
		TopicName: string(topicName),
		Messages:  make(chan *publishMessage),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr, nil
}

// Publish sends a message to the pubsub topic.
func (subscription *Subscription) Publish(message string) error {
	publishMessage := publishMessage{
		Message:  message,
		SenderID: subscription.self.Pretty(),
	}
	msgBytes, err := json.Marshal(publishMessage)
	if err != nil {
		return err
	}
	return subscription.topic.Publish(subscription.ctx, msgBytes)
}

func (subscription *Subscription) ListPeers() []peer.ID {
	return subscription.pubSub.ListPeers(subscription.TopicName)
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (subscription *Subscription) readLoop() {
	for {
		msg, err := subscription.sub.Next(subscription.ctx)
		if err != nil {
			close(subscription.Messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom == subscription.self {
			continue
		}

		message := new(publishMessage)
		err = json.Unmarshal(msg.Data, message)
		if err != nil {
			continue
		}

		// send valid messages onto the Messages channel
		subscription.Messages <- message
	}
}
