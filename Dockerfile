FROM golang:1.21

WORKDIR ["/opt/du"]

COPY . .

RUN ["apt", "update"]
RUN ["apt", "install", "-y", "make"]
RUN ["make", "all"]

CMD ["duapp"]
