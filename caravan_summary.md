# Caravan: Incentive-Driven Account Migration via Transaction Aggregation in Sharded Blockchain

## Overview
Caravan is an optimized account migration scheme for sharded blockchains that improves upon the Fine-tuned Lock mechanism. It introduces transaction aggregation for migrating accounts and an incentive-driven priority mechanism to accelerate migration processes.

## Core Innovation

### 1. Transaction Aggregation Mechanism
Caravan addresses the inefficiency in Fine-tuned Lock where only "payer" transactions are locked during migration. Instead, Caravan:

- **Gathers "payer" transactions** for migrating accounts during migration phase
- **Aggregates them into consolidated transactions** (big transactions) per account
- **Processes aggregated transactions** after migration completion
- **Reduces post-migration overhead** by minimizing the number of transactions to process

### 2. Incentive-Driven Priority System
- Increases revenue generated from migration transactions
- Incentivizes miners to prioritize migration transactions
- Accelerates the overall migration process through economic incentives

## Architecture Components

### Core Data Structures

#### Transaction Types
1. **Transaction2** (`core/transaction.go`): Enhanced transaction structure supporting multiple recipients
2. **TXmig1** (`core/txmig1.go`): Migration request transaction
3. **TXmig2** (`core/txmig2.go`): Aggregated migration transaction with proof data
4. **TXns** (`core/txns.go`): Notification transaction for incoming accounts
5. **TXann** (`core/txann.go`): Announcement transaction

#### Transaction Pools
1. **Tx_pool** (`core/txpool.go`): Main transaction pool with specialized sub-pools:
   - `Outing_Before_Announce_TX_Pools`: Transactions from accounts before migration announcement
   - `Outing_After_Announce_TX_Pools`: Transactions from accounts after migration announcement  
   - `Locking_TX_Pools`: Transactions involving locked accounts
   - `Coming_TX_Pools`: Transactions for incoming accounts
   - `Relay_Pools`: Cross-shard relay transactions

2. **Specialized Pools**:
   - `TXmig1_pool`: Migration request pool
   - `TXmig2_pool`: Aggregated migration transaction pool
   - `TXann_pool`: Announcement pool
   - `TXns_pool`: Notification pool

### Account Management (`account/`)
- **Account2Shard mapping**: Tracks which shard each account belongs to
- **Account state tracking**: Manages account balances and migration status
- **Lock mechanisms**: Controls account locking during migration

### Blockchain Core (`chain/`)
- **BlockChain structure**: Manages the blockchain state and transaction processing
- **Multi-level Merkle tree**: Modified structure for security in transaction aggregation
- **Storage layer**: Persistent storage using LevelDB

### Shard Management (`shard/`)
- **Node structure**: Represents nodes within shards
- **PBFT consensus**: Practical Byzantine Fault Tolerance for intra-shard consensus
- **Cross-shard communication**: Handles transaction relay between shards

## Migration Workflow

### Phase 1: Pre-Migration
1. **Account selection**: Accounts identified for migration using CLPA algorithm
2. **Transaction gathering**: "Payer" transactions for migrating accounts are collected
3. **Lock initiation**: Accounts enter migration state with appropriate locking

### Phase 2: Migration Execution  
1. **TXmig1 creation**: Migration requests created with proof data
2. **Transaction aggregation**: Payer transactions aggregated into TXmig2 structures
3. **Cross-shard announcement**: Migration announced to destination shard via TXann
4. **State transfer**: Account state transferred with Merkle proofs

### Phase 3: Post-Migration
1. **Aggregated transaction processing**: Consolidated TXmig2 transactions executed
2. **Account unlocking**: Migrated accounts unlocked in destination shard
3. **Notification completion**: TXns transactions confirm migration completion

## Key Configuration Parameters (`params/config.go`)

### Migration Settings
- `Stop_When_Migrating`: Whether to stop normal transactions during migration
- `Lock_Acc_When_Migrating`: Whether to lock accounts during migration
- `Not_Lock_immediately`: Delayed locking for optimization
- `RelayLock`: Lock mechanism for relay transactions

### Performance Parameters
- `MaxBlockSize`: Maximum transactions per block (default: 2000)
- `MaxMigSize`: Maximum migration transactions per block (default: 1000)
- `Inject_speed`: Transaction injection rate (default: 500 TXs/sec)
- `Block_interval`: Time between blocks (default: 6 seconds)

### Experimental Controls
- `Bu_Tong_Bi_Li`: Different transaction proportion experiments
- `Bu_Tong_Shi_Jian`: Different migration timing experiments
- `Pressure`: Stress testing with multiple account migrations
- `Cross_Chain`: Cross-chain migration experiments

## Performance Improvements

### Compared to Fine-tuned Lock
1. **Reduced locking overhead**: Only aggregates payer transactions instead of locking all
2. **Post-migration efficiency**: Processes aggregated transactions instead of individual ones
3. **Incentive alignment**: Economic incentives ensure faster migration processing
4. **Scalability**: Better handles high-volume migration scenarios

### Security Enhancements
1. **Modified Merkle tree**: Multi-level structure ensures proof integrity
2. **Consensus integration**: PBFT ensures agreement on migration state
3. **Atomic completion**: Migration either fully completes or fully rolls back

## Implementation Details

### Transaction Flow
```
Normal TX → Tx_pool → [If sender migrating] → Aggregation → TXmig2 → Post-migration execution
Migration request → TXmig1_pool → Announcement → State transfer → Completion
```

### Lock Management
- **Fine-grained locking**: Different lock types for different migration phases
- **Delayed locking**: `Not_Lock_immediately` option for optimization
- **Selective locking**: Only payer transactions affected during aggregation phase

### Cross-Shard Communication
- **Relay mechanism**: Transactions relayed between shards via TCP
- **Proof verification**: Merkle proofs verify cross-shard state transitions
- **Atomic updates**: Ensures consistency across shards

## Code Structure

```
block-caravan/
├── main.go                    # Entry point with test execution
├── core/                      # Core data structures
│   ├── transaction.go         # Transaction definitions
│   ├── txpool.go             # Main transaction pool
│   ├── txmig1.go             # Migration request
│   ├── txmig2.go             # Aggregated migration transaction
│   └── [other pool types]
├── chain/                     # Blockchain implementation
│   └── blockchain.go         # Blockchain state management
├── shard/                     # Shard management
│   └── shard.go              # Node and shard operations
├── account/                   # Account management
│   ├── account.go            # Account definitions
│   └── account_state.go      # Account state tracking
├── params/                    # Configuration
│   └── config.go             # System parameters
├── test/                      # Test suites
└── utils/                     # Utility functions
```
