FROM ubuntu:18.04

ENV DEBIAN_FRONTEND=noninteractive
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
ENV POSTGRES_URL postgres://docker:docker@localhost:5432/db?sslmode=disable

RUN mkdir -p /go /go/bin /go/src /go/src/github.com/c3systems/Hackathon-EOS-SF-2018 /go/pkg &&\
  apt-get update -y && apt-get upgrade -y &&\
  apt-get install -y --no-install-recommends --fix-missing make curl python gnupg2 dirmngr golang-go build-essential ca-certificates &&\
  apt-get autoremove -y &&\
  apt-get update -y --no-install-recommends

# Add the PostgreSQL PGP key to verify their Debian packages.
# It should be the same key as https://www.postgresql.org/media/keys/ACCC4CF8.asc
RUN ( apt-key adv --keyserver ha.pool.sks-keyservers.net --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8 \
  || apt-key adv --keyserver pgp.mit.edu --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8 \
  || apt-key adv --keyserver keyserver.pgp.com --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8 ) &&\
  echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list &&\
  apt-get update -y --no-install-recommends && apt-get install -y --no-install-recommends postgresql-9.3 postgresql-client-9.3 postgresql-contrib-9.3

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
# And add ``listen_addresses`` to ``/etc/postgresql/9.3/main/postgresql.conf``
RUN rm /etc/postgresql/9.3/main/pg_hba.conf &&\
    echo "local all all trust" >> /etc/postgresql/9.3/main/pg_hba.conf &&\
    echo "host all all 127.0.0.1/32 trust" >> /etc/postgresql/9.3/main/pg_hba.conf &&\
    echo "host all all ::1/128 trust" >> /etc/postgresql/9.3/main/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /etc/postgresql/9.3/main/postgresql.conf &&\
    /etc/init.d/postgresql restart

# Cd into the api code directory
WORKDIR /go/src/github.com/c3systems/Hackathon-EOS-SF-2018

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/c3systems/Hackathon-EOS-SF-2018

# note: this is insecure and just for demo purposes...
ENV ETH_PRIVATE_KEY 18afdddf061ab614bdacff35bbe3c58b4b464a95db18aecfa7c266fd61b4850c
ENV ETH_CONTRACT_ADDRESS 0xb366b07c070c380051893a33681a757116c9c685
ENV ETH_NODE_URL wss://rinkeby.infura.io/ws

RUN chmod +x /go/src/github.com/c3systems/Hackathon-EOS-SF-2018/docker-entrypoint.sh
ENTRYPOINT /go/src/github.com/c3systems/Hackathon-EOS-SF-2018/docker-entrypoint.sh
