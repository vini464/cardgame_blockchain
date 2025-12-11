
account:
	geth --datadir . account new

init:
	geth --datadir . init genesis.json

firstnode:
	geth --datadir . --networkid 2025 --http --http.addr "0.0.0.0" --http.port 8545 --http.api "eth,net,web3,personal" --allow-insecure-unlock --port 30303 --mine --miner.threads 1 --unlock $(ACCOUNT) --password /dev/null

secondnode:
	geth --datadir . --networkid 2025 --http --http.addr "0.0.0.0" --http.port 8545 --http.api "eth,net,web3,personal" --allow-insecure-unlock --port 30303 --bootnodes enode://$(key)@$(ip):30303
