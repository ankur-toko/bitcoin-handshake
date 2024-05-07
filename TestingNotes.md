# How to test

1. Download implementation of bitcoin node from https://github.com/btcsuite/btcd. 
2. `[Optional]` Comment out all DNSSeeds from file: `btcd/chaincfg/params.go.` This will prevent unnecessary downloading of the blockchain. we do not need the blockchain for this testing
3. Start the bitcoin node using "go run ." inside the btdc folder. This will start a server on 127.0.0.1:8333
4. Run the script under test using "go run .". Our client will try to eastablish connection with the node spawned on 8333 and start handshake by sending the "version" message. 
5. If everything runs well, nodes will exchange version and verack message and the node will accept out client as a valid peer. You should see the message below in the node's logs :
```
[2024-05-07 20:18:41.895 [INF] SYNC: New valid peer 127.0.0.1:63667 (inbound) (AnkurHandshakerV1.0)]"
```