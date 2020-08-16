# API para dados de covid no Ceará. :bar_chart:


## Descrição do projeto :pencil2:

API que fornece a consulta aos dados disponibilizados pelo governo do ceará sobre a situação da covid no estado.

Entre as informações disponibizadas extraidas do CSV:
- Casos por municipios(confirmados, obitos, curados).
- Incidência de casos.
- Entre outras informações.

## Executando o projeto :rocket:

### Clone o repositorio: 
execute `https://github.com/marlo2222/Covid-API.git`
`cd Covid-API`

### Executando e instalando depencias do projeto: 

- `docker-compose build` 

- `docker-compose up -d`

:warning: OBS :warning: Caso não tenha o `docker` ou `docker-compose` em sua maquina utilize o seguinte [tutorial](https://phoenixnap.com/kb/install-docker-compose-ubuntu) :thumbsup:.

:pushpin: OBS : Todas as dependências incluindo uma imagem `Golang` e um imagem do `Consul` já foram definidas no dockerfile e nos arquivos do docker-compose. :smile:

### Executando o Agente Consul : 

dentro do container da imagem do consul execute:

- `consul agent -server -ui -bootstrap-expect 1 -data-dir /var/consul.d -config-dir /etc/consul.d -node consul-master -client 0.0.0.0`

:pushpin: OBS: para entrar no container onde está executandoa imagem consul você pode usar : `docker-compose exec {container} bash`

### Sugestões e melhorias: :heart:

- Está é a primeira versão do projeto. Uma serie de melhorias ainda precisam ser feitas.

- Em breve estarei disponibilizando uma doc para os endpoints disponíveis  bem como um `link` para acesso.

### este projeto foi realizado junto com os colegas: 

- [Artur Castro](https://github.com/ArturCRS) :octocat:
