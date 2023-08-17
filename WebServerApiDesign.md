# BurnFeed - A Decentralized Twitter-Like DApp

## API

# BurnFeed HTTP API

**Base URL:** ``

## General Authentication

To access endpoints other than `Authenticate`, `Sign-in`, and `Register`, you need to include an `Authorization header` in your requests with a valid token. Tokens can be obtained through the `Authenticate` procedure.
- **Possible request (with header):**
```json
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_TOKEN" -H "Nonce: YOUR_NONCE" -H "User-Wallet-Address: YOUR_WALLET_ADDRESS" -d '{
    "type": "artifact",
    "version": "0.1",
    "subtype": "tweet",
    "user": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
    "content": "This serves as an example of a tweet."
}' http://localhost:8080/create-artifact
```


## Register User

- **URL:** `/register`
- **Method:** `POST`
- **Purpose:** Register a new user with their Ethereum wallet address.
- **Request Format:**
```json
  {
      "walletAddress": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ"
  }
```
- **Response Format:**
```json
Status message confirming user registration
```

## Sign In
- **URL:** `/sign-in`
- **Method:** `POST`
- **Purpose:** Initiate the signing process by providing the wallet address.
- **Request Format:**
```json
  {
    "walletAddress": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ"
  }
```
- **Response Format:**
```json
  {
    "nonce": "123456789"
  }
```


## Authenticate
- **URL:** `/authenticate`
- **Method:** `POST`
- **Purpose:** Authenticate the user's signature using their nonce and grant a token.
- **Request Format:**
```json
{
    "walletAddress": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
    "signedNonce": "0x..."
}
```
- **Response Format:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## Create artifact
- **URL:** `/artifact`
- **Method:** `POST`
- **Purpose:** Create a new artifact (tweet, message, etc.) and store it on IPFS.
- **Request Format:**
```json
{
    "type": "artifact",
    "version": "0.1",
    "subtype": "tweet",
    "user": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
    "content": "This serves as an example of a tweet."
}
```
- **Response Format:**
```json
{
    "ipfsCID": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
}
```
- **Possible testing:**
1. Run main.go for the server (localhost:8080 by default)
2. Open a terminal and use curl to test the endpoint
```
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer YOUR_TOKEN" -d '{
    "type": "artifact",
    "version": "0.1",
    "subtype": "tweet",
    "user": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
    "content": "This serves as an example of a tweet."
}' http://localhost:8080/create-artifact
```
3. Check the response, something similar
```
{"ipfsCID":"0x123456"}
```

## Create action
- **URL:** `/action`
- **Method:** `POST`
- **Purpose:** Create a new action (follow, retweet, etc.) and store it on IPFS.
- **Request Format:**
```json
{
    "type": "action",
    "version": "0.1",
    "subtype": "follow",
    "user": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
    "followee": "0xDef4567ghiJ891klMno234PQR567stUv890WXYz12aB"
}
```
- **Response Format:**
```json
{
    "ipfsCID": "ipfs:QmYzZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqGUXPNzXdC9"
}
```

## Get user
- **URL:** `/user/{address}`
- **Method:** `GET`
- **Purpose:** Retrieve user details by Ethereum address.
- **Response Format:**
```json
{
    "address": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
    "username": "burnUser123",
    "createdAt": "2023-08-05T12:00:00Z"
}
```

## Get suggested artifacts
- **URL:** `/suggested-artifacts/{address}`
- **Method:** `GET`
- **Purpose:** Retrieve suggested artifacts (equivalent to a feed on Twitter) based on user interactions/preferences.
- **Response Format:**
```json
List of suggested artifacts
```

## Like
- **URL:** `/like`
- **Method:** `POST`
- **Purpose:** Handle the "like" action for a tweet and update token burns.
- **Request Format:**
```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "like",
    "likeOf": "ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX",
    "burn": "10000000000000"
}
```
- **Response Format:**
```json
Status message confirming like action
```

## Get popular tweets
- **URL:** `/popular-tweets`
- **Method:** `GET`
- **Purpose:** Retrieve popular tweets based on token burns.
- **Response Format:**
```json
List of popular tweets
```

## Get popular tweets
- **URL:** `/tweets`
- **Method:** `GET`
- **Purpose:** Retrieve most recent tweets.
- **Response Format:**
```json
List of most recent tweets
```

## Get artifact
- **URL:** `/artifact/{ipfsCID}`
- **Method:** `GET`
- **Purpose:** Retrieve the content of an artifact by its IPFS CID.
- **Response Format:**
```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "like",
    "likeOf": "ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX",
    "burn": "10000000000000"
}
```

## Get action
- **URL:** `/action/{ipfsCID}`
- **Method:** `GET`
- **Purpose:** Retrieve the content of an action by its IPFS CID.
- **Response Format:**
```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "like",
    "likeOf": "ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX",
    "burn": "10000000000000"
}
```