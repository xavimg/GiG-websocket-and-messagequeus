package nats

import (
	"testing"
	"time"

	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"

	natsserver "github.com/nats-io/nats-server/test"
)

const TEST_PORT = 8369

func RunServerOnPort(port int) *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	return RunServerWithOptions(&opts)
}

func RunServerWithOptions(opts *server.Options) *server.Server {
	return natsserver.RunServer(opts)
}

func TestNATSConnection(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8369")
	if err != nil {
		t.Errorf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()
}

func TestWrongNATSConnection(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8368")
	if err == nil {
		t.Errorf("Error connecting to NATS server")
	}
	defer nc.Close()
}

func TestPublishToNATS(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8369")
	if err != nil {
		t.Errorf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()

	topic := "GiG"
	nc.Publish(topic, []byte("something"))
}

func TestPublishToNATSNoTopic(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8369")
	if err != nil {
		t.Errorf("Error connecting to NATS server: %v", err)
	}
	defer nc.Close()

	nc.Publish("", []byte("something"))
}

func TestPublishTopicExist(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8369")
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	topic := "test"

	data := []byte("hello")

	sub, err := nc.SubscribeSync("test")
	if err != nil {
		t.Fatal(err)
	}

	err = nc.Publish("test", data)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := sub.NextMsg(time.Second)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, topic, msg.Subject)
}

func TestPublisTopicNotExists(t *testing.T) {
	go RunServerOnPort(TEST_PORT)

	nc, err := nats.Connect("nats://localhost:8369")
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	sub, err := nc.SubscribeSync("test")
	if err != nil {
		t.Fatal(err)
	}

	expectedTopic := "test"
	topicSubscribed := "golang"
	err = nc.Publish(expectedTopic, []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	msg, err := sub.NextMsg(time.Second)
	if err != nil {
		t.Fatal(err)
	}

	require.NotEqual(t, topicSubscribed, msg.Subject)

}
