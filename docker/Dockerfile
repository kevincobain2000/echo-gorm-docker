FROM centos:7.3.1611

RUN yum update -y && yum install -y \
    sudo which git-core zlib zlib-devel gcc-c++ patch readline readline-devel libyaml-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison curl groupinstall "Development Tools" epel-release

RUN yum install -y golang httpd monit


COPY ./docker/Entrypoint.sh /
RUN chmod +x /Entrypoint.sh

RUN systemctl enable httpd

RUN go get github.com/lestrrat/go-server-starter/cmd/start_server
RUN go get github.com/cespare/reflex

ENTRYPOINT ["/Entrypoint.sh"]

CMD ["-c"]