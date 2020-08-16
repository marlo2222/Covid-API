FROM golang:onbuild

RUN cd ~
RUN apt update
RUN apt install vim -y
RUN apt install unzip -y
RUN wget https://releases.hashicorp.com/consul/1.7.2/consul_1.7.2_linux_amd64.zip
RUN unzip -f consul_1.7.2_linux_amd64.zip
RUN mv consul /usr/bin && rm -f consul*
RUN mkdir /etc/consul.d && mkdir /var/consul.d

EXPOSE 8080