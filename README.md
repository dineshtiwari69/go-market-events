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

\*\*Check `watcher/watcher.go` to see how to parse the incoming events
