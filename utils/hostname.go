package utils

var hostname string = "localhost"

// GetHostname function gets the name of host to be used while obtaining telemetry data
func GetHostname() string {
	return hostname
}

// SetHostname function sets the hostname to be used while performing queries
func SetHostname(newHostname string) {
	hostname = newHostname
}
