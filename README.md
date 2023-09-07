# JWT Authentication

##
Este projeto tem como objetivo testar a implementação de JSON Web Tokens (JWT) em uma API utilizando a linguagem Go (Golang).

## Instalação

Para executar este projeto, siga as etapas abaixo:

### Clonar o repositório

Você pode clonar o repositório usando o seguinte comando:

```bash
git clone https://github.com/matheusferreira165/jwt-authentication.git
```
### Configurar variaveis de ambiente
Crie um arquivo .env com as seguintes variaveis 

```bash
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

DB_HOST=db
DB_PORT=5432
DB_NAME=
DB_USER=
DB_PASSWORD=
```

### Utilizando Docker Compose

Certifique-se de que o docker e docker-compose esta instalado na maquina

```bash
docker-compose up
```
Isso iniciará o aplicativo, incluindo a api e o db PostgreSQL, em containers Docker.

### SE