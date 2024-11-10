run-flags:
	go run ./locks/flags
	
run-peterson:
	go run ./locks/peterson

run-proproducer-consumer-buffer:
	go run ./structures/producer-consumer-buffer

run-tas:
	go run ./locks/tas
	
run-ttas:
	go run ./locks/ttas
	
run-exonential-backoff:
	go run ./locks/exonential-backoff

run-anderson-queue:
	go run ./locks/anderson-queue
	
run-clh:
	go run ./locks/clh

run-mcs:
	go run ./locks/mcs