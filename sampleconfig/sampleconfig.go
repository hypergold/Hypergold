@@ -433,7 +433,7 @@ the method name for further details such as parameter and return information.
 |Parameters|None|
 |Description|Returns data about each connected network peer as an array of json objects.|
 |Returns|`(json array)`
`addr`: (string) the ip address and port of the peer
`services`: (string) the services supported by the peer
`lastrecv`: (numeric) time the last message was received in seconds since 1 Jan 1970 GMT
`lastsend`: (numeric) time the last message was sent in seconds since 1 Jan 1970 GMT
`bytessent`: (numeric) total bytes sent
`bytesrecv`: (numeric) total bytes received
`conntime`: (numeric) time the connection was made in seconds since 1 Jan 1970 GMT
`pingtime`: (numeric) number of microseconds the last ping took
`pingwait`: (numeric) number of microseconds a queued ping has been waiting for a response
`version`: (numeric) the protocol version of the peer
`subver`: (string) the user agent of the peer
`inbound`: (boolean) whether or not the peer is an inbound connection
`startingheight`: (numeric) the latest block height the peer knew about when the connection was established
`currentheight`: (numeric) the latest block height the peer is known to have relayed since connected
`syncnode`: (boolean) whether or not the peer is the sync peer
`[{"addr": "host:port", "services": "00000001", "lastrecv": n, "lastsend": n, "bytessent": n, "bytesrecv": n, "conntime": n, "pingtime": n, "pingwait": n, "version": n, "subver": "useragent", "inbound": true_or_false, "startingheight": n, "currentheight": n, "syncnode": true_or_false }, ...]`|
-|Example Return|`[{"addr": "178.172.xxx.xxx:11008", "services": "00000001", "lastrecv": 1388183523, "lastsend": 1388185470, "bytessent": 287592965, "bytesrecv": 780340, "conntime": 1388182973, "pingtime": 405551, "pingwait": 183023, "version": 70001, "subver": "/hcashd:0.4.0/", "inbound": false, "startingheight": 276921, "currentheight": 276955, "syncnode": true }, ...]`|
+|Example Return|`[{"addr": "178.172.xxx.xxx:14008", "services": "00000001", "lastrecv": 1388183523, "lastsend": 1388185470, "bytessent": 287592965, "bytesrecv": 780340, "conntime": 1388182973, "pingtime": 405551, "pingwait": 183023, "version": 70001, "subver": "/hcashd:0.4.0/", "inbound": false, "startingheight": 276921, "currentheight": 276955, "syncnode": true }, ...]`|
 [Return to Overview](#MethodOverview)
  
 ***
 +; All interfaces on port 14008:
+; listen=:14008
+; All ipv4 interfaces on port 14008:
+; listen=0.0.0.0:14008
+; All ipv6 interfaces on port 14008:
+; listen=[::]:14008
+; Only ipv4 localhost on port 14008:
+; listen=127.0.0.1:14008
+; Only ipv6 localhost on port 14008:
+; listen=[::1]:14008
 ; Only ipv4 localhost on non-standard port 8336:
 ; listen=127.0.0.1:8336
 ; All interfaces on non-standard port
