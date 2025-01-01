package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Definir flags
	address := flag.String("a", "nats://localhost:4222", "Dirección del servidor NATS (alias: --address)")
	channel := flag.String("c", "chat.room1", "Nombre del canal (alias: --channel)")
	name := flag.String("n", "Usuario", "Tu nombre en el chat (alias: --name)")

	// Soporte para nombres largos de flags
	flag.StringVar(address, "address", *address, "Dirección del servidor NATS")
	flag.StringVar(channel, "channel", *channel, "Nombre del canal")
	flag.StringVar(name, "name", *name, "Tu nombre en el chat")

	flag.Parse()

	// Validar flags requeridos
	if *address == "" || *channel == "" || *name == "" {
		fmt.Println("Uso: ./nats_chat -a <address> -c <channel> -n <name>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Conexión al servidor NATS
	nc, err := nats.Connect(*address)
	if err != nil {
		fmt.Printf("Error conectándose a NATS: %v\n", err)
		os.Exit(1)
	}
	defer nc.Close()

	fmt.Printf("Conectado a NATS en %s\n", *address)
	fmt.Printf("Canal: %s, Nombre: %s\n", *channel, *name)

	// Configurar JetStream
	js, err := nc.JetStream()
	if err != nil {
		fmt.Printf("Error al configurar JetStream: %v\n", err)
		os.Exit(1)
	}

	// Crear stream para persistencia de mensajes
	streamName := "CHAT_STREAM"
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{*channel},
		Storage:  nats.FileStorage,
		MaxAge:   time.Hour,
	})
	if err != nil {
		fmt.Printf("Error al crear el stream: %v\n", err)
	}

	// Recuperar mensajes pasados (última hora)
	fmt.Println("Recuperando mensajes pasados del último período...")
	subscription, err := js.Subscribe(*channel, func(msg *nats.Msg) {
		fmt.Printf("[%s]: %s\n", msg.Header.Get("name"), string(msg.Data))
	}, nats.DeliverLast())
	if err != nil {
		fmt.Printf("Error al suscribirse al canal: %v\n", err)
		os.Exit(1)
	}
	defer subscription.Unsubscribe()

	// Publicar mensajes
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escribe mensajes para el chat (Ctrl+C para salir):")
	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) == "" {
			continue
		}

		msg := &nats.Msg{
			Subject: *channel,
			Data:    []byte(text),
			Header:  nats.Header{"name": []string{*name}},
		}
		if err := nc.PublishMsg(msg); err != nil {
			fmt.Printf("Error publicando mensaje: %v\n", err)
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("Error leyendo entrada: %v\n", scanner.Err())
	}
}
