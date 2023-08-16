# BurnFeed - A Decentralized Twitter-Like DApp

## Overview

BurnFeed is a decentralized application (DApp) that mimics the functionality of Twitter. This project will use a [simple smart contract](https://github.com/notfound81/burnfeedprotocol/blob/main/src/BurnFeedProtocol.sol) inspired by [Simpubprotocol](https://github.com/simpubprotocol/simpubprotocol). It aims to achieve the following:

- Handle around 10 million active users per day on an Ethereum Layer 2 or Layer 3, corresponding to approximately 600 transactions per second (assuming each user executes 5 transactions daily),
- Provide a user experience akin to Twitter for disseminating information,
- Resist censorship by enabling users to download software and operate a node, creating a localized view of network information. It does not utilize any public DNS or website,
- Enable users to extend their reach or attention from potential viewers by burning specific tokens or Ether.
- Allow users to send encrypted messages.

The following are explicitly excluded from our goals:
- Tokenizing user content,
- Built-in anonymity - users must create new addresses to maintain anonymity and use applications like Tornado Cash to transfer assets to new addresses.

## Data Organization

The app's data is hierarchically organized into four types.

### Users

Every Ethereum address is a potential user. To prevent spamming, the app may require an address to possess a special NFT before transacting on-chain.

### Artifacts

Each artifact is saved as a file on IPFS, and its IPFS CID serves as the artifact's ID. Artifacts are stored on a user's local IPFS and are retrievable by others. If the artifact's hosting IPFS node is offline, some users may fail to fetch the artifacts, potentially resulting in shadow-banning of the artifact's creator. Users are advised to keep the app running in the background to avoid being shadow banned. Here's an example of an artifact as a tweet:

```json
{
    "type": "artifact",
    "version":"0.1",
    "subtype": "tweet",
    "user": "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5",
    "content": "This serves as an example of a tweet."
}
```
Once saved on IPFS, the file will have a CID `ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9` and you can use this as the tweet's URI.

Another example is an encrypted message:

```json
{
    "type": "artifact",
    "version":"0.1",
    "subtype": "message",
    "key_id": 1,
    "content": "U2FsdGVkX1+vupppZksvRf5pq5g5XjFRlipRkwB0K1Y96Qsv2Lm+31cmzaAILwyt"
}
```

### Actions

Actions form artifact-artifact, artifact-user, and user-user relations. For instance, if a user posts a tweet, we have:

```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "tweet",
    "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
}
```

If user `ethereum:0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ` followed `ethereum:0xDef4567ghiJ891klMno234PQR567stUv890WXYz12aB`, the action is defined as:

```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "follow",
    "user": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
    "followee": "0xDef4567ghiJ891klMno234PQR567stUv890WXYz12aB"
}
```

Users have the ability to follow others discreetly, without the need to publish any "follow" actions on the blockchain. This feature facilitates private or undisclosed following of individuals.

If a tweet `ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9` is a retweet of `ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX`, we should have an action defined as:

```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "tweet",
    "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
    "retweetOf": "ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX"
}
```

The action for sending an encrypted message:
```json
{
    "type": "action",
    "version":"0.1",
    "subtype": "send_message",
    "to": "0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
    "message": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
    "burn": "25000000000000"
}
```


Not every action is saved on IPFS, instead, many actions within a UI session are aggregated into an *action file* (also known as the session file) on IPFS. Aggregating the above three actions simplifies them as:

```json
{
    "type":"action",
    "version":"0.1",
    "actions":[
        {
            "subtype":"tweet",
            "tweet":"ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
        },
        {
            "subtype":"follow",
            "user":"0xAbC1234def5678ghI901jklmNoPq234Rstuv567WXyZ",
            "followee":"0xDef4567ghiJ891klMno234PQR567stUv890WXYz12aB"
        },
        {
            "subtype":"retweet",
            "tweet":"ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
            "retweetOf":"ipfs:QmQeKSUqZoGUXPNzZkH64BFAkVVhoFAPA8uBkfNydC9xX"
        }
    ]
}
```

### Publications

For each session, the URI of an action file is published on-chain. We refer to data published on-chain as a publication. A publication is defined as an event in the contract as:

```solidity
    event Actions(
        address indexed user,
        string indexed uri,
        uint256 burn
    );
```

BurnFeed will use the keccak hash of the JSON schema for validating actions and artifacts as the value for `spec` and use the keccak hash of "IPFS" (`0x2eff44a149f4fdc0ebc091d05846eece9649b64f4c485b1f548136aaaf6483ac`) as the value of `uriType`.

## Burn-to-Promote

As social media platforms thrive on capturing people's attention, BurnFeed offers users the functionality to sort artifacts and users based on the amount of tokens they burnt. If the user sets the minimum burn required per artifact to one token, then all artifacts with less than one token burnt will be hidden from the user. This design helps users filter out spam and lets willing users pay for attention or advertisement.

The app also lets users like certain artifacts (tweets), but the number of like actions is not as crucial. Each like action must burn some non-zero tokens, and the sum of all users' tokens burnt for liking a tweet is displayed for that tweet.

```json
   {
    "subtype": "like",
    "tweet": "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9",
    "burn": "10000000000000"
   }
```

## User Bans and Credibility Measures
Clients may implement various rules to enhance local content quality. Certain actions could lead to a decrease in a user's credibility score, including:

- Publishing artifacts that are unable to be fetched, excessively large, or do not comply with the client's specifications. Discrepancies between the action and artifact content, or between the total burn in the action file and the on-chain burn amount, can also trigger a decrease.
- A user can be explicitly banned via the user interface.

If a user's credibility score falls below a specified threshold, they may be subjected to a local ban for a predetermined period. Note that banned users may not be aware of their ban status, as these actions are performed locally.

## Client Implementation

### Client Variability
BurnFeed does not enforce specific client implementations, allowing for great flexibility in how local views of on-chain and off-chain published data are constructed. As a result, user experiences on BurnFeed may vary greatly depending on the chosen client. Essentially, users are grouped by the client implementations they choose. Furthermore, users have the option to utilize more than one client to observe different local views, thereby diversifying their user experience. Hence, BurnFeed embodies a truly decentralized and consensus-free design.



## Prototyping a Simple Client

### Technical Verifications

The following technical aspects have been verified (though this does not necessarily imply their adoption in the final implementation):

- The capability to determine the size of a file on IPFS without needing to fetch it remotely, effectively preventing spam. This can be accomplished utilizing the `ipfs object stat` command.
- A file on IPFS can be encapsulated using a 32-byte representation, given that the most frequently employed hash function within IPFS, SHA-256, generates a 32-byte hash.

### Modules

The client may have the following modules:

#### User Interface
The user interface should offer the following perspectives:

- Tweet Details: This mirrors Twitter's tweet view. It displays the tweet's content, recent retweets, and its appreciation measure, which is the total number of tokens burned (akin to likes).
- User Profile: Showcases the recent tweets posted by the user.
- Latest Tweets: Displays the most recent tweets, sorted by timestamp.
- Popular Tweets: Highlights the tweets with the highest token burns in the past 24 hours, emphasizing those with substantial engagement.
- (Optionally) Message Inbox: Lists the recent messages received by the user.
- (Optionally) Message Details: Showcases an individual message along with its associated metadata.

In terms of functionality, the UI should provide forms for users to like (burn for) a tweet, retweet, and post a new tweet. Optinally, the interface may facilitate the creation and delivery of encrypted messages.

#### Action Indexer
The Action Indexer subscribes to contract events to create and update a local database for indexing actions. While these events are quite lightweight, maintaining a historical record of all actions can consume substantial disk space and processing time. To mitigate this, the Action Indexer periodically purges outdated action data. When initiated, the Action Indexer doesn't necessarily process events from the app's genesis block. Instead, it has the flexibility to focus on more recent events, as determined by the user or client configurations.

The Indexer may have to maintain the following relation tables (assuming using MySQL):

**Tweet Table**

| Column Name | Data Type | Description |
| ----------- | --------- | ----------- |
| `user` | VARCHAR(42) | Stores the address of the user. |
| `tweet` | TEXT | Stores the URI of the tweet. |
| `retweetOf` | TEXT | Stores the URI of the tweet that is retweeted. |
| `burn` | BIGINT | Records the number of tokens burned for the particular tweet. |
| `likes` | BIGINT | Records the number of tokens burned by users who liked this tweet. |
| `createdAt` | DATETIME | Timestamp of when the tweet was created. |

**Follow Table**

| Column Name | Data Type | Description |
| ----------- | --------- | ----------- |
| `user` | VARCHAR(42) | Stores the address of the user. |
| `followee` | TEXT | Stores the address of the followee. |
| `createdAt` | DATETIME | Timestamp of when the follow action was created. |

#### API Server

The API server functions as the intermediary between the User Interface and the source data, retrieving content from the Action Indexer's database and IPFS. To ensure efficient and prompt data retrieval for all views, the Indexer's database should be properly indexed.

In addition to this, the API server also communicates with the IPFS daemon to fetch files from remote peers based on demand. Once the artifact files are locally available, the UI is notified via WebSocket notifications.

## Third-party Services
Certain organizations may opt to index and retain all action and artifact data for analysis. For instance, this data can be leveraged for AI-driven services, such as content recommendation or tagging. Client applications can offer users the option to use these services to enhance content quality and overall user experience.





