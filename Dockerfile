FROM golang
MAINTAINER HackerZ hackerzgz@gmail.com
ADD main.go /
CMD ["go","run","main.go"]
