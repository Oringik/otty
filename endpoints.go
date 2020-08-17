package otty

// CreateEndpoint creates endpoint for call function by this endpoint after
func (otty *Otty) CreateEndpoint(endpoint string, function func([]byte)) {
	otty.Endpoints[endpoint] = function
}

// ResolveEndpoint call function of endpoint by endpoint name
func (otty *Otty) ResolveEndpoint(endpoint []byte, data []byte) {
	endpointString := string(endpoint)

	functionToCall := otty.Endpoints[endpointString]
	if functionToCall == nil {
		return
	}

	functionToCall(data)
	return

}
