package constants

const (
	JWTKeyEmail         = "email"
	JWTKeyUserID        = "user_id"
	JWTKeyExpiration    = "exp"
	JWTKeyCertificateID = "certificate_id"
	JWTKeyPOAID         = "poa_id"

	// Format: "1s", "13h" (max - h, d - days - not allowed)
	// See "time.ParseDuration"
	ViperJWTTTLKey    = "service.jwt_ttl"
	ViperJWTSecretKey = "service.jwt_secret"
)
