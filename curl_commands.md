# BurnFeed - A Decentralized Twitter-Like DApp

## API

# curl command examples

## Get nonce (First part of the signing in)

- **Example request:**
```
curl -X POST -H "Content-Type: application/json" -H "" -d '{}' http://localhost:8080/get-nonce
```
- **Example response:**
```json
{"nonce":"xHL5/YhoDO18iq7MUlhKmocUlY8QXciMhOAp1K2RIJU="}
```

## Sign in (user signs the nonce + wallet address)

- **Example request (with valid data):**
```
curl -X POST -H "Content-Type: application/json" -H "Signature: 0xe3b93fbb3195843e501c6e7c7def51d7b3332d0e19ef0502cabcc04d64a8d80b281601d281ce5e403037404e600add036042c0a3d4f3eac1bbd03258f9ca1f631c" -d '{
    "walletAddress": "0xacF3411e61d4D2C8e839f29252984983a48f212D",
    "nonce": "xHL5/YhoDO18iq7MUlhKmocUlY8QXciMhOAp1K2RIJU="
}' http://localhost:8080/sign-in

```
- **Example response:**
```json
{"nonce":"xHL5/YhoDO18iq7MUlhKmocUlY8QXciMhOAp1K2RIJU="}
```

## Create artifact

- **Example request:**
```
curl -X POST -H "Content-Type: application/json" -d '{
    "type": "artifact",
    "version": "0.1",
    "subtype": "tweet",
    "user": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
    "content": "This serves as an example of a tweet."
}' http://localhost:8080/create-artifact

```
- **Example response:**
```json
{"ipfsCID":"0x123456"}
```