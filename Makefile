SED=sed
ifeq ($(shell uname), Darwin)
	SED=gsed
endif

default: build

test:
	go test -cover -v

testnew:
	go test -cover -run=TestMatchNew -v

bench:
	go test -run=XXX -bench=BenchmarkMatch -benchtime=1m -v

cover:
	go test -cover -coverpkg github.com/liamJunkermann/manta,github.com/liamJunkermann/manta/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatch2159568145 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

memprofile:
	go test -v -run=TestMatch2159568145 -test.memprofile=/tmp/manta.memprof -test.memprofilerate=1
	go tool pprof --alloc_space manta.test /tmp/manta.memprof

update: update-protobufs generate

update-protobufs:
	./update_protos.sh

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
