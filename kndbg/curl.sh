#!/bin/bash
curl -v "msgdumper-channel-4w6gn.kn-dbg.svc.cluster.local" \
  -X POST \
  -H "Ce-Id: c568f3e2-f8bd-4d5d-b90d-a873b3effb00" \
  -H "Ce-Knativebrokerttl: 10" \
  -H "Ce-Specversion: 0.2" \
  -H "Ce-Time: 2019-04-02T18:31:15.860829687Z" \
  -H "Ce-Type: type1" \
  -H "Content-Type: application/json" \
  -H "X-B3-Sampled: 1" \
  -H "X-B3-Spanid: 7e7d70d4d0e7de6d" \
  -H "X-B3-Traceid:7e7d70d4d0e7de6d" \
  -H "X-Request-Id: 9a5b97df-fbfe-9f48-95db-5ccdf5dbf4d8" \
  -H "Ce-Source: source2" \
  -d '{Reader:{"msg":"Body-type1-source2","sequence":"1"}}' 