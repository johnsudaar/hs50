package client

type Client struct {
	IP   string
	Port int
}

// Generate a new HS50 client
// ip:   IP of the console   (Ex: 192.168.0.8)
// port: Port of the console (Ex: 60040)
func NewClient(ip string, port int) (*Client, error) {
	c := &Client{
		IP:   ip,
		Port: port,
	}

	return c, nil
}
