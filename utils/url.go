package utils

// GetBaseURL function retrieves the base url to be used while redirecting requests
func GetBaseURL() string {
	return "http://$hostname$:8111/$path$"
}
