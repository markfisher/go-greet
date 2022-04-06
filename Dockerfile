FROM golang
	WORKDIR /workspace
	COPY go.mod go.mod
	COPY app.go app.go
	RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM scratch
	COPY --from=0 /workspace/app .
	ENTRYPOINT [ "/app" ]
