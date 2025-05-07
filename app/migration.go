package app

func UploadMigration() {
	OpenTerminal("docker compose up --build -w && yarn migrate")
}
