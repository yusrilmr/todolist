package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Start a db session
	createDBConnection()
}