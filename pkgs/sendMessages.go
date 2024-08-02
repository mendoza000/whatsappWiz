package pkgs

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"os"
	"whatsappWiz/data"
)

func SendMessage(client *whatsmeow.Client, groupJID string, name string) {

	// Leer la imagen del archivo
	imagePath := "data/cuentas-hs.png"
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error al enviar el mensaje en el grupo " + name + " con el JID " + groupJID)
		return
	}

	// Subir la imagen a WhatsApp
	resp, err := client.Upload(context.Background(), imageData, whatsmeow.MediaImage)
	if err != nil {
		fmt.Println("Error al enviar el mensaje en el grupo " + name + " con el JID " + groupJID)
		return
	}

	imageMsg := &waProto.ImageMessage{
		Caption: proto.String(
			"*Disponibles cuentas premium sin caidas 🔥*" +
				"\n\n - *Netflix original:* $ 2.99" +
				"\n - *HBO Max:* $ 1.99" +
				"\n - *Disney Plus:* $ 2.49" +
				"\n - *Spotify:* desde $ 2.49" +
				"\n - *Prime video:* $ 1.49" +
				"\n - *Crunchyroll:* $ 1.99" +
				"\n - *Youtube Premium:* desde $ 2.49" +
				"\n\n😉 Aceptamos *Bolivares, Binance, Paypal, Zinli, " +
				"Bancolombia y efectivo* " +
				"\n\n🦇 *Siguenos en instagram para no perderte de nada!" +
				"* https://instagram.com/house.streamingxx"),
		Mimetype: proto.String("image/png"), // replace this with the actual mime type
		// you can also optionally add other fields like ContextInfo and JpegThumbnail here

		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	_, err = client.SendMessage(context.Background(), types.JID{
		User:   groupJID,
		Server: types.GroupServer,
	}, &waProto.Message{
		ImageMessage: imageMsg,
	})
	if err != nil {
		fmt.Println("Error al enviar el mensaje en el grupo " + name + " con el JID " + groupJID)
		SendMessageToLog(client, data.GroupLog().JID, data.GroupLog().Name,
			"🤨🚩 *El mensaje no se pudo enviar en:* "+name)
		return
	}
}
