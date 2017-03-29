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

# Get repo
```bash
git clone http://github.com/jonathankentstevens/grpc-tracing-lab $GOPATH/src/github.com/jonathankentstevens/grpc-tracing-lab
```
    
# Setting up gRPC

- Get code
```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
```

- Install protoc
```bash
wget https://github.com/google/protobuf/releases/download/v3.2.0rc2/protoc-3.2.0rc2-linux-x86_64.zip
unzip protoc-3.2.0rc2-linux-x86_64.zip
sudo cp bin/protoc /usr/local/bin
```

# set environment vars
      export GCP_PROJECT={gcp project id}
    
      export GCP_SVCACCT_KEY= {path to your service account key file}

# Labs
<a href="https://github.com/jonathankentstevens/grpc-tracing-lab/blob/master/helloworld/README.md">Hello World</a><br>
<a href="https://github.com/jonathankentstevens/grpc-tracing-lab/blob/master/weather-search/README.md">Weather Search</a>