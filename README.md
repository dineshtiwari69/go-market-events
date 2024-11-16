# Markets events

## How to run

```
make run
```

## Events

### Proposed price

When someone thinks a market is ready to be resolved, proposes a price. This means that the market is close to be resolved.
This action emits this event:

```solidity

event ProposePrice(
    address indexed requester,
    address indexed proposer,
    bytes32 identifier,
    uint256 timestamp,
    bytes ancillaryData,
    int256 proposedPrice,
    uint256 expirationTimestamp,
    address currency
);
```

#### How to relate this event with a market in the API

- Parse the `ancillaryData` to question id

```go
questionId := crypto.Keccak256Hash(log.AncillaryData)
```

- Then use that value here:

```
curl --location 'https://clob.polymarket.com/markets-by-question-id/{question id}'
```

### Disputed price

If the resolution is accepted the market gets resolved and that's all. However, if someone disagres is possible to dispute the proposed resolution. This means that the resolution will take a bit longer.

```solidity
event DisputePrice(
    address indexed requester,
    address indexed proposer,
    address indexed disputer,
    bytes32 identifier,
    uint256 timestamp,
    bytes ancillaryData,
    int256 proposedPrice
);
```

#### How to relate this event with a market in the API

- Parse the `ancillaryData` to question id

```go
questionId := crypto.Keccak256Hash(log.AncillaryData)
```

- Then use that value here:

```
curl --location 'https://clob.polymarket.com/markets-by-question-id/{question id}'
```

### Condition resolved

When the market actually gets resolved, this event is emitted:

```solidity
event ConditionResolution(
    bytes32 indexed conditionId,
    address indexed oracle,
    bytes32 indexed questionId,
    uint outcomeSlotCount,
    uint[] payoutNumerators
);
```

#### How to relate this event with a market in the API

- Use the `conditionId` field to filter in the API request

```
curl --location 'https://clob.polymarket.com/markets/{condition id}'
```

\*\*Check `watcher/watcher.go` to see how to parse the incoming events


## Docker

### Build
To build the Docker image:


⌘K to generate a command
@Codebase Add a section in Readme.md on how to build a docker file and run it

I'll help you add a Docker section to the README.md. Based on the provided code, I can see there's already a Dockerfile and the project structure.
Add this section after the "How to run" section in the README.md (after line 8):
## Docker

### Build
To build the Docker image:

```bash
docker build -t market-events .
```


### Run
To run the container:
```bash
docker run market-events

```
