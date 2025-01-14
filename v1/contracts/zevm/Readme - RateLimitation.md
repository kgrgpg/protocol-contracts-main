Here's the properly formatted markdown documentation:

# RateLimitedZRC20 Documentation

## Introduction

**RateLimitedZRC20** is a proof-of-concept (PoC) demonstrating rate-limited withdrawals on top of ZetaChain's ZRC20 token standard. This ensures that within a 24-hour period:
- Users can only withdraw a limited number of times.
- Users can only withdraw up to a certain total volume of tokens.

This mitigates risks like flash-loan exploits or large outflows in a short time frame. For demonstration purposes, cross-chain gas fee logic is skipped, focusing solely on the rate-limit mechanism.

## Why Rate Limiting?

### Security & Stability
- Reduces the impact of compromised accounts or malicious usage.
- Prevents mass outflows in a short time frame.

### User Experience
- Encourages withdrawals aligned with governance or protocol constraints.

### Protocol-Level Enforcement
- Rate limiting is integrated into the token's `withdraw()` function, enforcing constraints for all callers.

## Architecture & Key Components

### Forking ZRC20

The PoC starts with ZetaChain's official ZRC20 contract, which includes:
- Standard ERC20 logic (balances, transfers).
- Cross-chain functionalities like `withdraw()` and `deposit()`.
- Constructor restricted to mainnet's "Fungible Module."

For local development, deployment checks are relaxed to allow a normal Hardhat account to deploy.

### Daily Window Tracking

Each user's withdrawal data is tracked via a struct:

```solidity
struct WithdrawLimit {
    uint256 count;
    uint256 volume;
    uint256 windowStart;
}
```

On withdraw():
1. Check if block.timestamp >= windowStart + 24 hours. If true, reset count, volume, and windowStart.
2. Ensure count + 1 <= maxWithdrawCountPerDay.
3. Ensure volume + amount <= maxWithdrawVolumePerDay.
4. Revert if any limit is exceeded.

### Bypassing Cross-Chain Gas Fees

Since no real SystemContract exists in local development:
- Calls to withdrawGasFee() are removed or replaced to avoid reverting.

## Code Walkthrough

### Constructor Modifications

Removed checks like:

```solidity
if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
```

to allow local deployment.

### Deposit Logic

The deposit logic is adjusted as follows:

```solidity
function deposit(address to, uint256 amount) external override returns (bool) {
    // Relaxed sender validation
    _mint(to, amount);
    emit Deposit(abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), to, amount);
    return true;
}
```

### Rate-Limited Withdraw Logic

```solidity
function withdraw(bytes memory to, uint256 amount) external override returns (bool) {
    _checkAndUpdateWithdrawLimits(msg.sender, amount);
    _burn(msg.sender, amount);
    emit Withdrawal(msg.sender, to, amount, 0, PROTOCOL_FLAT_FEE);
    return true;
}
```

### Skipping withdrawGasFee() Locally

Fee logic is skipped by either commenting out fee-related calls or overriding them:

```solidity
function withdrawGasFee() public view override returns (address, uint256) {
    revert ZeroGasCoin();
}
```

## Local Development Setup

### Prerequisites
- Node.js and Yarn (or npm)
- Hardhat
- Foundry (optional for localnet) or an Anvil node

### ZetaChain Localnet or Plain Anvil

You can either:
1. Run npx hardhat localnet from ZetaChain's example contracts repo
2. Spin up an Anvil node and connect Hardhat to it

## Deployment & Demo

### Compile & Deploy

1. Compile:
```bash
npx hardhat compile
```

2. Deploy:
```bash
npx hardhat run scripts/deploy-rate-limited.ts --network localhost
```

### Mint Tokens (Deposit)

Mint tokens and verify balances:

```javascript
const rateLimited = await ethers.getContractAt(
  "RateLimitedZRC20",
  "<DEPLOYED_ADDRESS>"
);

await rateLimited.deposit(
  "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
  ethers.utils.parseEther("50")
);
```

### Testing Withdrawals

1. Withdraw tokens:
```javascript
await rateLimited.withdraw("0xabc123", ethers.utils.parseEther("10"));
```

2. Confirm limits:
- Count exceeded: "Exceeded max daily withdrawal count"
- Volume exceeded: "Exceeded max daily withdrawal volume"

### Daily Window Reset (Optional)

Fast-forward time for testing:

```javascript
await network.provider.send("evm_increaseTime", [86401]);
await network.provider.send("evm_mine");
```

## Potential Enhancements

1. Gas Fee Mock: Deploy a mock SystemContract
2. User-Specific Limits: Enable per-user thresholds
3. Admin-Adjustable Windows: Allow dynamic adjustment of window size
4. Event Logging: Add granular logs for debugging

## Conclusion

This PoC demonstrates how to enforce rate limits on ZRC20 withdrawals. For production environments, integrate real cross-chain fees via the SystemContract. This foundation offers significant security and stability benefits.
