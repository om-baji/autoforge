package main

type Ec2Response struct {
}

type Ec2Request struct {
	region   string
	instance string
	ami      string
	gateway  bool
	ssh      string
}
