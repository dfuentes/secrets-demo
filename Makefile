stealth:
	@go get github.com/Clever/stealth

build:
	@go get github.com/Masterminds/glide
	@GO15VENDOREXPERIMENT=1 glide up
	@go build -o bootstrap

plan:
	@terraform get tf/
	@terraform plan tf/

apply:
	@terraform apply tf/
