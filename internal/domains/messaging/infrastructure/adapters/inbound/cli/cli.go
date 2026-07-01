package cli

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
	"github.com/vapankov/yaca/internal/domains/messaging/infrastructure/adapters/outbound/clock"
	idGen "github.com/vapankov/yaca/internal/domains/messaging/infrastructure/adapters/outbound/generators/identifiers"
	messageRepoFileStore "github.com/vapankov/yaca/internal/domains/messaging/infrastructure/adapters/outbound/repositories/message/filestore"
	fileStorage "github.com/vapankov/yaca/internal/domains/messaging/infrastructure/storage/file"
)

const (
	messagingCommandPost = "post"
	messagingCommandView = "view"
)

func Run() {
	defer func() {
		fmt.Println("QUIT")
	}()

	const (
		messageStoreFileDefault = "messages"
		messagingCommandDefault = messagingCommandView
	)

	var (
		messageStoreFile string
		messagingCommand string
		messagingData    string
	)

	flag.StringVar(&messageStoreFile, "file", messageStoreFileDefault, "Use storage file.")
	flag.StringVar(&messagingCommand, "cmd", messagingCommandDefault, "Command to run: \"post\" or \"view\".")
	flag.StringVar(&messagingData, "data", "", "Message to send.")

	flag.Parse()

	var (
		fileLineStore     = fileStorage.NewLineStore(messageStoreFile)
		messageRepository = messageRepoFileStore.New(fileLineStore)
		messsageIDGen     = idGen.NewMessage()
		clock             = clock.New()

		messagingUsecases = usecases.New(
			messageRepository,
			messsageIDGen,
			clock,
		)
	)

	switch messagingCommand {
	case messagingCommandPost:
		if len(messagingData) == 0 {
			fmt.Println("EMPTY MESSAGE")
			return
		}

		if _, err := messagingUsecases.PostMessage(context.TODO(), &usecases.PostMessageInput{
			MessageContents: values.MessageContents(messagingData),
		}); err != nil {
			fmt.Println("ERROR", err.Error())
			return
		}

		fmt.Println("MESSAGE SENT")

	case messagingCommandView:
		result, err := messagingUsecases.ViewMessages(context.TODO(), &usecases.ViewMessagesInput{})
		if err != nil {
			fmt.Println("ERROR", err.Error())
			return
		}

		if len(result.Items) == 0 {
			fmt.Println("NO MESSAGES")
			return
		}

		for i := range result.Items {
			printMessage(&result.Items[i])
		}
	}

}

func printMessage(msg *usecases.ViewMessagesResultItem) {
	fmt.Println(">>>")
	fmt.Println("ID\t", msg.MessageID)
	fmt.Println("DATA\t", msg.MessageContents)
	fmt.Println("TS\t", time.Time(msg.MessageCreatedAt).Local().Format(time.RFC3339))
	fmt.Println("<<<")
}
