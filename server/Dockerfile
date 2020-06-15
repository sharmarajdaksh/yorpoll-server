FROM golang:1.14

WORKDIR /usr/app/

RUN go get -u github.com/cosmtrek/air

COPY . .

RUN go get

CMD [ "air", "-c", ".air.conf" ]
