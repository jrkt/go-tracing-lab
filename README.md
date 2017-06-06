# Tracing requests with Google Stackdriver
https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/trace/trace.go

This is an example repo for getting Google's Stackdriver Trace tracing requests across microservices.

# Create new GCP project

<a href="https://console.cloud.google.com">Google Cloud Console</a>

# Enable billing

Go to <a href="https://console.cloud.google.com/billing/">Billing</a> to set up a billing account & make sure it is enabled for the new project you created.

# Enable Stackdriver Trace API

<a href="https://console.cloud.google.com/apis/dashboard">API Dashboard</a>

- click "Enable API"
- search "trace"
- select "Stackdriver Trace API"
- click "Enable"

# Create service account

Go to <a href="https://console.cloud.google.com/iam-admin/serviceaccounts">Service Accounts</a>

- click "Create service account"
- name it "tracer"
- select the Role of "Cloud Trace Agent"
- select "Furnish a new private key"
- click "CREATE"

The JSON key will be downloaded.

# set environment vars
      export GCP_PROJECT={project id of the new project you created}
    
      export GCP_KEY={full path of the service account key file that you just downloaded}
      
# Get repo
```
git clone http://github.com/jonathankentstevens/go-tracing-lab $GOPATH/src/github.com/jonathankentstevens/go-tracing-lab
```

# HTTP REST Labs
<a href="https://github.com/jonathankentstevens/go-tracing-lab/tree/master/rest/helloworld">Hello World</a><br>
<a href="https://github.com/jonathankentstevens/go-tracing-lab/tree/master/rest/convo">Silicon Valley Conversation</a>

# gRPC Labs
<a href="https://github.com/jonathankentstevens/go-tracing-lab/tree/master/grpc/helloworld">Hello World</a><br>
<a href="https://github.com/jonathankentstevens/go-tracing-lab/tree/master/grpc/weather-search">Weather Search</a>

# Reference: Setting up gRPC (linux)

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