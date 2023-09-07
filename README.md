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

### Sem Docker
1. Instale o Postgres, mais informacoes em: https://www.postgresql.org/download/
2. Fazer a build do projeto
    ```bash 
        go build -o server
    ```
3. Executar o arquivo gerado
   ```bash
   ./server
   ```

## Uso

Após a instalação e execução bem-sucedidas, você pode usar a API para realizar a autenticação com JWT. Aqui estão algumas operações de exemplo:

- **Registrar um usuário:** Envie uma solicitação POST para `/api/v1/register` com os dados do usuário EX:
   ```json
   {
	"name": "Rosa",
	"email": "rosa@gmail.com",
	"password": "rosa123"
    }
   ```
  

- **Fazer login:** Envie uma solicitação POST para `/api/v1/login` com suas credenciais para obter um token JWT EX:
   ```json
   {
	"email": "rosa@gmail.com",
	"password": "rosa123"
    }
   ```

- **Acessar recurso protegido que retorna os dados do usuario:** Envie uma solicitação GET para `/api/v1/user`
  Possiveis resultados: 

  Usuario Logado:
  ```json
    {
	"id": 1,
	"name": "Rosa",
	"email": "rosa@gmail.com"
    }
  ```
  Usuario deslogado:
  ```json
   {
	"message": "not logged"
   }
  ```
  Usuario nao autorizado:
  ```json
    {
	"message": "unauthentication"
    }
  ```
  **Deslogar:** Envie uma solicitação GET para `/api/v1/logout` com suas credenciais para obter um token JWT

### Tecnologias Utilizadas

Neste projeto, foram utilizadas as seguintes tecnologias:

- Go (Golang) 
- PostgreSQL 
- JSON Web Tokens (JWT)


