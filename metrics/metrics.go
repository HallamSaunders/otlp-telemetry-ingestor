package metrics

type MetricsFile struct {
	ResourceMetrics []ResourceMetric `json:"resourceMetrics"`
}

type ResourceMetric struct {
	Resource     Resource      `json:"resource"`
	ScopeMetrics []ScopeMetric `json:"scopeMetrics"`
}

type Resource struct {
	Attributes []Attribute `json:"attributes"`
}

type ScopeMetric struct {
	Scope   Scope    `json:"scope"` // Can remain empty, or define fields as needed
	Metrics []Metric `json:"metrics"`
}

type Scope struct {
	// Add fields if necessary
}

type Metric struct {
	Name string `json:"name"`
	Unit string `json:"unit,omitempty"`
	Sum  Sum    `json:"sum,omitempty"`
}

type Sum struct {
	DataPoints             []DataPoint `json:"dataPoints"`
	AggregationTemporality int         `json:"aggregationTemporality,omitempty"`
	IsMonotonic            bool        `json:"isMonotonic,omitempty"`
}

type DataPoint struct {
	Attributes        []Attribute `json:"attributes,omitempty"`
	StartTimeUnixNano string      `json:"startTimeUnixNano,omitempty"`
	TimeUnixNano      string      `json:"timeUnixNano,omitempty"`
	AsInt             string      `json:"asInt,omitempty"`
}

type Attribute struct {
	Key   string     `json:"key"`
	Value ValueField `json:"value"`
}

type ValueField struct {
	StringValue string `json:"stringValue,omitempty"`
}
