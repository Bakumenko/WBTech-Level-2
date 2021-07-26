package pkg

type ship struct {
	transport
}

func newShip() ITransport {
	return &ship{
		transport{
			name:  "Ship",
			speed: 50,
		},
	}
}
