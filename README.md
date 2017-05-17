# Tracing requests with Google Stackdriver
https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/trace/trace.go

This is an example repo for getting Google's Stackdriver Tracer tracing requests across microservices.

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
```
git clone http://github.com/jonathankentstevens/go-tracing-lab $GOPATH/src/github.com/jonathankentstevens/go-tracing-lab
```

# set environment vars
      export GCP_PROJECT={gcp project id}
    
      export GCP_KEY={path to your service account key file}

# HTTP REST Labs
<a href="https://github.com/jonathankentstevens/go-tracing-lab/rest/blob/master/helloworld/README.md">Hello World</a><br>
<a href="https://github.com/jonathankentstevens/go-tracing-lab/rest/blob/master/weather-search/README.md">Weather Search</a>

# gRPC Lab
<a href="https://github.com/jonathankentstevens/go-tracing-lab/grpc/blob/master/weather-search/README.md">Weather Search</a>

# Reference: Setting up gRPC

- Get code
```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
```

- Install protoc
```
wget https://github.com/google/protobuf/releases/download/v3.2.0rc2/protoc-3.2.0rc2-linux-x86_64.zip
unzip protoc-3.2.0rc2-linux-x86_64.zip
sudo cp bin/protoc /usr/local/bin
```