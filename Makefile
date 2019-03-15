binary:
	go build -o have-i-been-pwned .
image:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o have-i-been-pwned . \
	&& docker build --no-cache -t have-i-been-pwned:latest -t rekzi/have-i-been-pwned:latest . \
	&& rm have-i-been-pwned
push:
	docker push rekzi/have-i-been-pwned:latest