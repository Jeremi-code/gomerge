package app

func InitiateDocker() {
	OpenTerminal("docker compose up --build -w")
}

func UploadMigration() {
	InitiateDocker()
	OpenTerminal("yarn migrate")
}
