package traces

type TraceFile struct {
	// ResourceSpans contains spans grouped by resource - the highest level in the hierarchy
	ResourceSpans []ResourceSpan `json:"resourceSpans"`
}

type ResourceSpan struct {
	Resource   Resource    `json:"resource"`
	ScopeSpans []ScopeSpan `json:"scopeSpans"`
}

type Resource struct {
	Attributes []Attribute `json:"attributes"`
}

type ScopeSpan struct {
	Scope Scope  `json:"scope"`
	Spans []Span `json:"spans"`
}

type Scope struct {
	// Add fields if necessary
}

type Span struct {
	TraceID                string      `json:"traceId"`
	SpanID                 string      `json:"spanId"`
	ParentSpanID           string      `json:"parentSpanId,omitempty"`
	Name                   string      `json:"name"`
	StartTimeUnixNano      string      `json:"startTimeUnixNano"`
	EndTimeUnixNano        string      `json:"endTimeUnixNano"`
	DroppedAttributesCount int         `json:"droppedAttributesCount,omitempty"`
	Events                 []SpanEvent `json:"events,omitempty"`
	DroppedEventsCount     int         `json:"droppedEventsCount,omitempty"`
	Links                  []SpanLink  `json:"links,omitempty"`
	DroppedLinksCount      int         `json:"droppedLinksCount,omitempty"`
	Status                 *Status     `json:"status,omitempty"`
}

type SpanEvent struct {
	TimeUnixNano           string      `json:"timeUnixNano"`
	Name                   string      `json:"name"`
	Attributes             []Attribute `json:"attributes,omitempty"`
	DroppedAttributesCount int         `json:"droppedAttributesCount,omitempty"`
}

type SpanLink struct {
	TraceID                string      `json:"traceId"`
	SpanID                 string      `json:"spanId"`
	Attributes             []Attribute `json:"attributes,omitempty"`
	DroppedAttributesCount int         `json:"droppedAttributesCount,omitempty"`
}

type Status struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type Attribute struct {
	Key   string     `json:"key"`
	Value ValueField `json:"value"`
}

type ValueField struct {
	StringValue string `json:"stringValue,omitempty"`
}
