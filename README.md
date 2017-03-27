# gRPC Tracer

This is an example repo for getting the Stackdriver Trace working in gRPC

# set environment vars
      export GCP_PROJECT={gcp project id}
    
      export GCP_SVCACCT_KEY= {path to your service account key file}

# run server
    go run server/server.go
    
# run client
    go run client/client.go
