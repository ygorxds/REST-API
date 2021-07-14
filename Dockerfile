#conteiner para API-REST
#Autores: Ygor Pereira e Dejair Santos
#Data: 02/07/2021 - 07/07/2021
 
 FROM golang:latest

 WORKDIR /app

 RUN export GO111MODULE=ON

 COPY go.mod .
 COPY go.sum .
 COPY main.go .
 COPY . .
 RUN go get github.com/gorilla/mux
 RUN go get google.golang.org/api
 RUN go build
 EXPOSE 8000
 CMD [ "./main.go" ]
 
