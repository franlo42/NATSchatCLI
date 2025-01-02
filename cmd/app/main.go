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

func printLogo() {
	logo := `
     __  _   _____  __      _           _     ___   __   _____ 
  /\ \ \/_\ /__   \/ _\ ___| |__   __ _| |_  / __\ / /   \_   \
 /  \/ //_\\  / /\/\ \ / __| '_ \ / _` + "`" + ` | __|/ /   / /     / /\/
/ /\  /  _  \/ /   _\ \ (__| | | | (_| | |_/ /___/ /___/\/ /_  
\_\ \/\_/ \_/\/    \__/\___|_| |_|\__,_|\__\____/\____/\____/  
`
	fmt.Println(logo)
}

func main() {
	// Definir flags
	address := flag.String("a", "", "Dirección del servidor NATS (alias: --address)")
	channel := flag.String("c", "", "Nombre del canal (alias: --channel)")
	name := flag.String("n", "", "Tu nombre en el chat (alias: --name)")

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

	printLogo()

	fmt.Printf("Conectado a NATS en %s\n", *address)
	fmt.Printf("Canal: %s, Nombre: %s\n", *channel, *name)

	// Configurar JetStream
	js, err := nc.JetStream()
	if err != nil {
		fmt.Printf("Error al configurar JetStream: %v\n", err)
		os.Exit(1)
	}

	// Crear un stream por canal si no existe
	streamName := fmt.Sprintf("chat_%s", strings.Replace(*channel, ".", "_", -1)) // Usar el nombre del canal como el nombre del stream
	_, err = js.StreamInfo(streamName)
	if err != nil {
		// Si el stream no existe, crearlo
		fmt.Println("El stream no existe, creando nuevo stream...")
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{*channel},
			Storage:  nats.FileStorage,
			MaxAge:   time.Hour, // Guardar mensajes por una hora
		})
		if err != nil {
			fmt.Printf("Error al crear el stream: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Stream creado correctamente.")
	} else {
		fmt.Println("Stream existente encontrado.")
	}

	// Recuperar mensajes pasados (última hora)
	fmt.Println("Recuperando mensajes pasados del último período...")
	subscription, err := js.Subscribe(*channel, func(msg *nats.Msg) {
		fmt.Printf("[%s]: %s\n", msg.Header.Get("name"), string(msg.Data))
	}, nats.DeliverAll())
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
