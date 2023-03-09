FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go build -o thedudebridge
EXPOSE 5758
CMD ./thedudebridge