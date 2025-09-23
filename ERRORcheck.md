Login Succeeded!
[1/2] STEP 1/7: FROM golang:1.24-alpine AS builder
Resolving "golang" using unqualified-search registries (/etc/containers/registries.conf)
Trying to pull docker.io/library/golang:1.24-alpine...
Getting image source signatures
Copying blob sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1
Copying blob sha256:9824c27679d3b27c5e1cb00a73adb6f4f8d556994111c12db3c5d61a0c843df8
Copying blob sha256:72e8fc27da55cce2fb2e46c294e2373522d6729d092352a2d7221eb8f9a77a2d
Copying blob sha256:994678d1c2a93bf605782bddb80bc3a0c17db79ca705f3c2205cc880671e0dc7
Copying blob sha256:40e451c10b31532d8616ff155adb8383bcfe95b7388ad69363af6512986430f9
Copying config sha256:1d954e402d242cf11bd1ab7e108b59e92c43e4f1d5cb7bda41a77bb3dcffdd32
Writing manifest to image destination
[1/2] STEP 2/7: ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64
--> abe2b099d9a8
[1/2] STEP 3/7: WORKDIR /app
--> 734c749a198c
[1/2] STEP 4/7: COPY go.mod go.sum ./
--> 33133bfb88b1
[1/2] STEP 5/7: RUN go mod download && go mod verify
all modules verified
--> 536624c96ebd
[1/2] STEP 6/7: COPY . .
--> 91a4164763bf
[1/2] STEP 7/7: RUN go build -a -installsuffix cgo -o main .
database.go:12:2: string literal not terminated
Error: building at STEP "RUN go build -a -installsuffix cgo -o main .": while running runtime: exit status 1
time="2025-09-23T09:33:15.645Z" level=info msg="sub-process exited" argo=true error="<nil>"
time="2025-09-23T09:33:15.645Z" level=info msg="not saving outputs - not main container" argo=true
Error: exit status 1