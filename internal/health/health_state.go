package health


type HealthState int

const (
	Healthy HealthState = iota
	Unhealthy
	Trial
)

func (hs HealthState) String() string {
	switch hs {
	case Healthy:
		return "Healthy"
	case Unhealthy:
		return "Unhealthy"
	case Trial:
		return "Trial"
	default:
		return "Unknown"
	}
}

