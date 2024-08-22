package main

import (
	"encoding/json"
	"log"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message reqBody as an loosely-typed map.
		var reqBody map[string]any
		if err := json.Unmarshal(msg.Body, &reqBody); err != nil {
			return err
		}

		body := make(map[string]any)
		body["type"] = "echo_ok"
		body["echo"] = reqBody["echo"]

		return n.Reply(msg, body)
	})

	n.Handle("generate", func(msg maelstrom.Message) error {
		var reqBody map[string]any
		if err := json.Unmarshal(msg.Body, &reqBody); err != nil {
			return err
		}

		body := make(map[string]any)
		body["type"] = "generate_ok"
		body["id"] = "sdfs"
		if msgId, ok := reqBody["msg_id"].(float64); !ok {
			log.Fatal("couldn't convert msg_id")
		} else {
			// could've used https://github.com/oklog/ulid as well
			body["id"] = n.ID() + "-" + strconv.FormatInt(int64(msgId), 10)
		}

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
