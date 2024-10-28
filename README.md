# go-getting-started
Learn Golang

## Run api

```bash
go run main.go api
```

## Run migration scripts

```bash
go run main.go migrate
```

## create migration file

```bash
make migration name=<input-name>
```

## Install mockery

```bash
go install github.com/vektra/mockery/v2@latest
```

## Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate swagger docs

```bash
swag init
```

## Generate JWT public and private key

```bash
# Generate a private key (private.pem) with 2048-bit RSA encryption
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048

# Extract the public key (public.pem) from the private key
openssl rsa -pubout -in private.pem -out public.pem
```

## Google authentication qr format

```text
otpauth://totp/Example:alice@google.com?secret=TTETFQGNNTNRHHSY&issuer=Example
```

use `qr` tool to generate qr code image

```bash
pip install qrcode
```

gen image

```bash
qr "otpauth://totp/Example:alice@google.com?secret=TTETFQGNNTNRHHSY&issuer=Example"
```