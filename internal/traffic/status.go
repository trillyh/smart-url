package traffic

type TrafficStatus uint8

const (
	Valid TrafficStatus = iota
	Fraud
	Duplicate
)

func (s TrafficStatus) String() string {
	switch s {
	case Valid:
		return "VALID"
	case Fraud:
		return "FRAUD"
	case Duplicate:
		return "DUPLICATE"
	}
}
