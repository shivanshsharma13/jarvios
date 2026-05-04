package provider

import "net"

// Current returns "ollama" if no network, else "cloud"
func Current() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "ollama"
	}
	for _, i := range ifaces {
		if i.Flags&net.FlagUp != 0 &&
			i.Flags&net.FlagLoopback == 0 {
			addrs, _ := i.Addrs()
			if len(addrs) > 0 {
				return "cloud"
			}
		}
	}
	return "ollama"
}

// Chat routes to correct provider
func Chat(message string) (string, error) {
	if Current() == "ollama" {
		return OllamaChat(message)
	}
	// cloud provider goes here later
	return OllamaChat(message) // fallback for now
}
