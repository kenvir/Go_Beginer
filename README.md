# Go_Beginner

# Package Go

- https://pkg.go.dev/std

# Another Module

- Create two folder
- In folder main:
  - go mod init example.com/hello (example.com can change, hello is folder main)
  - create file hello.go
  - go mod edit -replace example.com/greetings=../greetings
  - go mod tidy
- In folder extra
  - go mod init example.com/greetings (example.com can change, hello is folder extra)
  - create file greetings.go
