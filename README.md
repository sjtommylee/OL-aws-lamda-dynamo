## AWS Lambda + DynamoDB + Go

CRUD, auth
`serverless.yml`

## notes

- Argon2-based password hashing [Argon2](https://auth0.com/blog/hashing-in-action-understanding-bcrypt/).
- Input validation
- Data consistency with DynamoDB transactions

# Serverless offline

requirements:

- node.js
  - [serverless framework](https://www.npmjs.com/package/serverless)
  - [serverless offline](https://www.npmjs.com/package/serverless-offline)
- golang
- docker
- [make](https://formulae.brew.sh/formula/make)

Install npm dependencies:

```
npm install // install serverless offline
```

use [make](https://formulae.brew.sh/formula/make) to run commands in `Makefile`

```
make start // clean, build and start the project
```
