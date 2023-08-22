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
        "type": "action",
        "timestamp": "2023-08-20 13:02:09.846505 +0800 CST m=+0.285431793",
        "actions": [
            {
                "subtype": "tweet",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "tweet",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
                "retweetOf": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "follow",
                "user": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
                "followee": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
            },
            {
                "subtype": "like",
                "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            },
            {
                "subtype": "send_message",
                "to": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
                "message": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
            }
        ]
    }' http://localhost:8181/create-artifact

```
- **Example response:**
```json
{"ipfsCID":"0x123456"}
```