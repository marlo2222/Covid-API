FROM nginx:1.18.0-alpine
RUN apk update && apk add bash bind-tools vim curl
RUN wget https://releases.hashicorp.com/consul/1.7.2/consul_1.7.2_linux_amd64.zip
RUN unzip consul_1.7.2_linux_amd64.zip && mv consul /usr/bin && rm -f consul*
RUN mkdir /etc/consul.d && mkdir /var/consul.d