# gRPC Tracer with Stackdriver

This is an example repo for getting the Stackdriver Trace working in gRPC

# Create new GCP project

<a href="https://console.cloud.google.com">Google Cloud Console</a>

# Enable billing & add to project

Do some stuff here...

# Enable Stackdriver Trace API

<a href="https://console.cloud.google.com/apis/dashboard">API Dashboard</a>

- click "Enable API"
- search "trace"
- select "Stackdriver Trace API"
- click "Enable"

# Get code

    go get github.com/jonathankentstevens/grpc-tracer
    
# Setting up gRPC

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
wget https://github.com/google/protobuf/releases/download/v3.2.0rc2/protoc-3.2.0rc2-linux-x86_64.zip
unzip protoc-3.2.0rc2-linux-x86_64.zip
sudo cp bin/protoc /usr/local/bin
```

# set environment vars
      export GCP_PROJECT={gcp project id}
    
      export GCP_SVCACCT_KEY= {path to your service account key file}

# run server
    go run server/server.go
    
# run client
    go run client/client.go
