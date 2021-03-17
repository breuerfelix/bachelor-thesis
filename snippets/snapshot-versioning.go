type HiveMQ struct {
	ID          string
	HostName    string
	Port        uint32
	Weight      uint32
	Healthy     bool
}

func NewHiveMQ(host string, port uint32) *HiveMQ {
	o := new(HiveMQ)

	o.HostName = host
	o.Port = port
	o.MetricsPort = metricsPort

	// default values
	o.Weight = 1
	o.Healthy = false

	o.ID = fmt.Sprintf("%v:%v", host, port)

	return o
}

func (h *HiveMQ) GetHashString() string {
	return fmt.Sprintf("%v-%v-%v", h.ID, h.Weight, h.Healthy)
}

func generateVersion(endpoints []*HiveMQ) string {
	// sort array so nodes are always in the same order
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].ID < endpoints[j].ID
	})

	s := sha256.New()
	for _, endpoint := range endpoints {
		s.Write([]byte(
			fmt.Sprintf("%v", endpoint.GetHashString())
		))
	}

	return fmt.Sprintf("%x", s.Sum(nil))
}
