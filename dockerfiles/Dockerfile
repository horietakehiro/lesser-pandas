FROM ubuntu:18.04

COPY config/ /config
COPY codes/ /codes
COPY entrypoint.sh /entrypoint.sh

RUN apt update -y \
    && apt install -y `cat /config/apt.txt` \
    && apt clean all

ENV PATH ${PATH}:/usr/local/go/bin
ENV GOPATH /codes/lesser-pandas:${GOPATH}

RUN wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz \
    && rm -fr go1.14.2.linux-amd64.tar.gz

RUN chmod +x /entrypoint.sh

WORKDIR /codes/lesser-pandas
ENTRYPOINT ["/entrypoint.sh"]
