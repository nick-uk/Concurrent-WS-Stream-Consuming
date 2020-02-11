# Concurrent WS Stream Consuming/Processing 
This project demonstrates the power of Go for fast concurrent stream processing with channels and waitgroups.

### What is this repository for?
Demostrates how to process a mysterious WebSocket stream using dynamically a number of workers equal with thew number of the CPU cores.

### How to test it?
```bash
$ go run test -v ./...
2020/02/11 16:14:34 Worker 1 is pending...
2020/02/11 16:14:34 Worker 6 is pending...
2020/02/11 16:14:34 Worker 4 is pending...
2020/02/11 16:14:34 Worker 5 is pending...
2020/02/11 16:14:34 Worker 9 is pending...
2020/02/11 16:14:34 Worker 7 is pending...
2020/02/11 16:14:34 Worker 3 is pending...
2020/02/11 16:14:34 Worker 8 is pending...
2020/02/11 16:14:34 Worker 2 is pending...
2020/02/11 16:14:34 Worker 12 is pending...
2020/02/11 16:14:34 Worker 10 is pending...
2020/02/11 16:14:34 Worker 11 is pending...
=== RUN   Test_wsClient
2020/02/11 16:14:34 Service listen on :12345 under /ws
=== RUN   Test_wsClient/Connectivity_check
2020/02/11 16:14:35 => Connecting to ws://127.0.0.1:1234
=== RUN   Test_wsClient/Mock_server
2020/02/11 16:14:35 => Connecting to ws://127.0.0.1:12345
=== RUN   Test_wsClient/Mock_server2
2020/02/11 16:14:35 => Connecting to ws://127.0.0.1:12345/ws
2020/02/11 16:14:35 Send: 26cf1f5a-76a2-bbbf-6062-a602f8b5c670  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 26cf1f5a-76a2-bbbf-6062-a602f8b5c670 which takes 1 sec & 780 ms from worker 1
2020/02/11 16:14:35 Send: a9f24247-ec41-dc3c-14d7-bc268b92ab76  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for a9f24247-ec41-dc3c-14d7-bc268b92ab76 which takes 1 sec & 780 ms from worker 6
2020/02/11 16:14:35 Send: e98ca7f6-1fa4-92df-3612-bf692bd22eb8  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for e98ca7f6-1fa4-92df-3612-bf692bd22eb8 which takes 1 sec & 780 ms from worker 4
2020/02/11 16:14:35 Send: 56265289-75e2-92ff-8e24-70139e1a1988  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 56265289-75e2-92ff-8e24-70139e1a1988 which takes 1 sec & 780 ms from worker 5
2020/02/11 16:14:35 Send: 7309c4ca-af0f-6818-921a-b80affa9a528  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 7309c4ca-af0f-6818-921a-b80affa9a528 which takes 1 sec & 780 ms from worker 9
2020/02/11 16:14:35 Send: da28ba96-ad4b-7309-8de9-51cb9881abcb  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for da28ba96-ad4b-7309-8de9-51cb9881abcb which takes 1 sec & 780 ms from worker 7
2020/02/11 16:14:35 Send: d926a595-45e8-d6d0-a52c-074dc039d3f5  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for d926a595-45e8-d6d0-a52c-074dc039d3f5 which takes 1 sec & 780 ms from worker 3
2020/02/11 16:14:35 Send: 80a31eca-196c-2fc7-db50-27378bf392fd  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 80a31eca-196c-2fc7-db50-27378bf392fd which takes 1 sec & 780 ms from worker 8
2020/02/11 16:14:35 Send: 69198547-16d4-6071-7eb6-35ca52a10819  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 69198547-16d4-6071-7eb6-35ca52a10819 which takes 1 sec & 780 ms from worker 2
2020/02/11 16:14:35 Send: 1246e3ec-1af6-c14d-852a-b7899c2485f7  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 1246e3ec-1af6-c14d-852a-b7899c2485f7 which takes 1 sec & 780 ms from worker 12
2020/02/11 16:14:35 Send: 67987760-578e-b117-2f59-90e93dafb2bd  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 67987760-578e-b117-2f59-90e93dafb2bd which takes 1 sec & 780 ms from worker 10
2020/02/11 16:14:35 Send: 535086c0-8826-088d-1b46-3bf42f52ff8b  of type: 1 for processing
2020/02/11 16:14:35 Fake processing in progress for 535086c0-8826-088d-1b46-3bf42f52ff8b which takes 1 sec & 780 ms from worker 11
2020/02/11 16:14:35 Send: 7bbe7c8e-9bc0-b65f-319a-a4140b07506a  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for 7bbe7c8e-9bc0-b65f-319a-a4140b07506a which takes 1 sec & 780 ms from worker 6
2020/02/11 16:14:37 Result: 588c2502-7058-61ab-c716-8a8c29006ea1 1
2020/02/11 16:14:37 Result: 098f5b6a-e6a3-4922-2406-b139d9781bfc 2
2020/02/11 16:14:37 Result: 0c90873f-958a-b5fc-f174-bd0a9baa9d7b 3
2020/02/11 16:14:37 Result: e28e222e-c0fb-7adc-2040-d62eda37bc25 4
2020/02/11 16:14:37 Result: 4e06631e-fe3e-742b-01a1-13a01a90f22f 5
2020/02/11 16:14:37 Result: 798e71af-c9a5-7c2f-d097-81dcc012d746 6
2020/02/11 16:14:37 Send: 36128896-1916-892c-56ae-f6180ad77190  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for 36128896-1916-892c-56ae-f6180ad77190 which takes 1 sec & 780 ms from worker 9
2020/02/11 16:14:37 Send: b0aea541-0b97-2060-1f09-b34388cfa104  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for b0aea541-0b97-2060-1f09-b34388cfa104 which takes 1 sec & 780 ms from worker 5
2020/02/11 16:14:37 Send: 4a925111-e1f9-8392-9aa6-d8e43323bb68  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for 4a925111-e1f9-8392-9aa6-d8e43323bb68 which takes 1 sec & 780 ms from worker 1
2020/02/11 16:14:37 Send: a480503b-3b42-d474-89ea-c95d258b5f52  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for a480503b-3b42-d474-89ea-c95d258b5f52 which takes 1 sec & 780 ms from worker 4
2020/02/11 16:14:37 Send: 5270a1d9-5522-39b0-b731-19c42c761fbb  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for 5270a1d9-5522-39b0-b731-19c42c761fbb which takes 1 sec & 780 ms from worker 8
2020/02/11 16:14:37 Send: 81ae315b-650d-46fb-2f9a-a87d878678dc  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for 81ae315b-650d-46fb-2f9a-a87d878678dc which takes 1 sec & 780 ms from worker 10
2020/02/11 16:14:37 Result: a7c68865-40ca-943b-9134-665ac36f01e5 7
2020/02/11 16:14:37 Result: 6f34e4c4-97c8-88c0-57c7-7fea40311a93 8
2020/02/11 16:14:37 Result: 3bbf9cc8-a022-34e1-9041-1989543aa5d0 9
2020/02/11 16:14:37 Result: 1e93fdb1-b713-88f3-83f5-61ff70889354 10
2020/02/11 16:14:37 Result: 048a0935-35af-6aa6-2acf-b6902ac0df13 11
2020/02/11 16:14:37 Result: fdaf7b32-6f86-1847-49f1-6f4ea8ccb7f2 12
2020/02/11 16:14:37 Send: dab9a74d-53ce-c235-e3f1-da480e5fa27b  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for dab9a74d-53ce-c235-e3f1-da480e5fa27b which takes 1 sec & 780 ms from worker 11
2020/02/11 16:14:37 closing WS
2020/02/11 16:14:37 Send: c4204f89-16b0-7f70-fbb5-8f43d5ca73a0  of type: 1 for processing
2020/02/11 16:14:37 Fake processing in progress for c4204f89-16b0-7f70-fbb5-8f43d5ca73a0 which takes 1 sec & 780 ms from worker 7
2020/02/11 16:14:37 ReadMessage() websocket: close 1000 (normal): stream is over
2020/02/11 16:14:39 Result: 1891501a-3f28-9db9-2879-5008b7e932e5 13
2020/02/11 16:14:39 Result: 164bdce9-8149-11cf-77d3-5a3b6cfdc4c6 14
2020/02/11 16:14:39 Result: 7d30d191-9e45-9465-3cd9-2299e12ee070 15
2020/02/11 16:14:39 Result: 7ca2e41f-5119-606a-4953-a44d5c30519a 16
2020/02/11 16:14:39 Result: e3dfe6e5-568e-23b6-6249-8a6889b4e89b 17
2020/02/11 16:14:39 Result: 86c411eb-e6f1-5fa1-7bf3-62457db4fe88 18
2020/02/11 16:14:39 Result: 0cc6e800-9d4b-dd86-21b1-ddba3ba8150a 19
2020/02/11 16:14:39 Result: 919aa245-6ca7-43c1-1afd-4a32006f3922 20
2020/02/11 16:14:39 Result: b51cb45a-5a80-ab8b-f75c-23b490c13b7d 21
2020/02/11 16:14:39 Stream is over & processing is done. Cheers
--- PASS: Test_wsClient (4.57s)
    --- PASS: Test_wsClient/Connectivity_check (0.00s)
        ws_test.go:93: wsClient() OK expected error: Connecting: dial tcp 127.0.0.1:1234: connect: connection refused
    --- PASS: Test_wsClient/Mock_server (0.00s)
        ws_test.go:93: wsClient() OK expected error: Connecting: websocket: bad handshake
    --- PASS: Test_wsClient/Mock_server2 (1.79s)
        ws_test.go:93: wsClient() OK expected error: <nil>
PASS
ok      consume-ws-with-chans   5.554s
```

### Contribution guidelines
* The usual please. Test it well locally, then create a new branch then do a pull request.

### Who do I talk to?
* If you watching this page probably you know already how to contact me. 

