FROM golang:alpine3.23
LABEL authors="sergey.deynego@gmail.com"




ENTRYPOINT ["top", "-b"]