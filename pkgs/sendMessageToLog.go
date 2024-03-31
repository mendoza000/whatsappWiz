package pkgs

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func SendMessageToLog(client *whatsmeow.Client, groupJID string, name string, message string) {
	_, err := client.SendMessage(context.Background(), types.JID{
		User:   groupJID,
		Server: types.GroupServer,
	}, &waProto.Message{
		Conversation: proto.String(message),
	})

	if err != nil {
		fmt.Println("Error al enviar el mensaje en el grupo " + name + " con el JID " + groupJID)
		return
	}
}
