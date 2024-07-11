package main

//go:generate godotenv -f ./.env tern migrate --migrations .\internal\pgstore\migrations\ --config .\internal\pgstore\migrations\tern.conf
