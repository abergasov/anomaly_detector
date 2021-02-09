package anomaly_analysator

type Analyser struct {
}

func InitAnalyser() *Analyser {
	return &Analyser{}
}

func (a *Analyser) GetCurrentState() []string {
	return []string{"1"}
}
