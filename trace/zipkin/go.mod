module github.com/xuqil/observability/trace/zipkin

go 1.19

require (
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/zipkin v1.15.1
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.15.1
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	golang.org/x/sys v0.7.0 // indirect
)
