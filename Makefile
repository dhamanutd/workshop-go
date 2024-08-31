# Run the application
.PHONY: run
run-http:
	go run main.go serveHttp --config configs/.config.yaml --secret configs/.secret.yaml
