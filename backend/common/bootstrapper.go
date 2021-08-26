package common

// StartUp initiates the configuration and setup
func StartUp() {
	// Initialize AppConfig variable
	loadAppConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Start a db session
	createDBConnection()
}