package app

func UploadMigration() {
	OpenTerminal("docker compose up --build -d && yarn migrate")
}
