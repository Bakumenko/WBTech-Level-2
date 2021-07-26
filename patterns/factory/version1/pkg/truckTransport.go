package pkg

type truck struct {
	transport
}

func newTruck() ITransport {
	return &truck{
		transport{
			name:  "truck",
			speed: 90,
		},
	}
}
