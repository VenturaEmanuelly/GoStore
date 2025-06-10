#!/bin/bash

set -e

echo "⛔ Parando e removendo containers + volumes..."
docker compose down -v

echo "✅ Subindo container novamente..."
docker compose up -d

echo "⌛ Aguardando o banco iniciar..."
sleep 5

echo "🚀 Rodando a aplicação Go..."
go run main.go
