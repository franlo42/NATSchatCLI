#!/bin/bash

# Imagen y contenedor para la construcción de binarios ejecutables.
IMAGE_NAME="nats-chat-builder"
CONTAINER_NAME="nats-chat-builder-container"


echo "Construyendo la imagen de compilación..."
docker build -t $IMAGE_NAME . || { echo "Error: No se pudo construir la imagen de compilación"; exit 1; }

echo "Creando un contenedor temporal..."
docker create --name $CONTAINER_NAME $IMAGE_NAME || { echo "Error: No se pudo crear el contenedor temporal"; exit 1; }

echo "Extrayendo ejecutables..."
docker cp $CONTAINER_NAME:/app/nats_chat_linux ./nats_chat_linux || { echo "Error: No se pudo extraer nats_chat_linux"; exit 1; }
docker cp $CONTAINER_NAME:/app/nats_chat_mac ./nats_chat_mac || { echo "Error: No se pudo extraer nats_chat_mac"; exit 1; }

echo "Limpiando recursos temporales..."
docker rm $CONTAINER_NAME || { echo "Error: No se pudo eliminar el contenedor temporal"; exit 1; }

echo "Construcción y extracción completadas. Ejecutables disponibles:"
echo "  - ./nats_chat_linux"
echo "  - ./nats_chat_mac"

echo "Iniciando el servidor NATS con docker-compose..."
docker-compose up -d || { echo "Error: No se pudo iniciar el servidor NATS"; exit 1; }
echo "Servidor NATS en ejecución."

echo "¡Listo!"
