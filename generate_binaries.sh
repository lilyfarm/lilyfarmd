# Binaries for this to be deployed on AWS Lightsail
# This is a cost efficient way of not having to compile on 
# the server. Currently, we are only paying for 0.5 GB of
# RAM and very limited compute.
GOOS=linux GOARCH=amd64 go build -o binaries/lilyfarm