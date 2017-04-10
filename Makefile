stealth:
	@go get github.com/Clever/stealth

build:
	@go build -o bootstrap

plan:
	@terraform get tf/
	@terraform plan tf/

apply:
	@terraform apply tf/
