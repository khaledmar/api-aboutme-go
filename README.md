# api-aboutme-go
Api desenvolvido em golang + webframework beego + mongodb drive mgov2

Para executar o projeto é necessário instalar e configurar GOLANG: https://golang.org/dl/

Após concluir as configurações abra seu terminal/prompt e execute os comandos para instalar os pacotes:

go get gopkg.in/mgo.v2
go get github.com/astaxie/beego
go get github.com/beego/bee

Acesse a pasta do projeto e execute o comando:
bee run -downdoc=true -gendoc=true

Abra o navegador e digite: http://localhost:8080/swagger
