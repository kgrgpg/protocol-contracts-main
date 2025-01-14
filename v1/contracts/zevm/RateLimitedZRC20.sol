// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "./interfaces/IZRC20.sol";
import "./interfaces/ISystem.sol";

/**
 * @title RateLimitedZRC20
 * @dev A fork of the ZRC20 contract with a daily withdrawal rate limit (both frequency and volume).
 *
 * Inherits from original ZRC20 but overrides withdraw(...) to include rate-limiting checks.
 */
contract RateLimitedZRC20 is IZRC20Metadata, ZRC20Events {
    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~ Original ZRC20 Storage ~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    /// @notice Fungible address is always the same, maintained at the protocol level
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    /// @notice Chain ID for this ZRC20
    uint256 public immutable CHAIN_ID;
    /// @notice Coin type (Gas, ERC20, etc). See enum in IZRC20.sol
    CoinType public immutable COIN_TYPE;
    /// @notice System contract address
    address public SYSTEM_CONTRACT_ADDRESS;
    /// @notice Gas limit used to compute gas fees for withdrawals
    uint256 public GAS_LIMIT;
    /// @notice Protocol flat fee added to gas fee
    uint256 public override PROTOCOL_FLAT_FEE;

    // Balance and allowance tracking
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;

    // ERC20 standard metadata
    string private _name;
    string private _symbol;
    uint8 private _decimals;

    /// Custom errors from original ZRC20
    error CallerIsNotFungibleModule();
    error InvalidSender();
    error GasFeeTransferFailed();
    error ZeroGasCoin();
    error ZeroGasPrice();
    error LowAllowance();
    error LowBalance();
    error ZeroAddress();

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~ Rate Limiting Storage ~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    struct WithdrawLimit {
        uint256 count;       // how many times user withdrew in current window
        uint256 volume;      // total tokens user withdrew in current window
        uint256 windowStart; // when current daily window started
    }

    mapping(address => WithdrawLimit) private _withdrawInfo;

    uint256 public constant WINDOW_SIZE = 86400;  // 24h in seconds
    uint256 public maxWithdrawCountPerDay = 3;
    uint256 public maxWithdrawVolumePerDay = 1000 ether;

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Modifiers ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    /**
     * @dev Only Fungible Module can call certain admin-type functions.
     *      We comment out the check for local dev to let us deploy & admin freely.
     */
    modifier onlyFungible() {
        // if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        _;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~ Constructor ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 chainid_,
        CoinType coinType_,
        uint256 gasLimit_,
        address systemContractAddress_
    ) {
        // if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        // (commented out for local dev)

        _name = name_;
        _symbol = symbol_;
        _decimals = decimals_;

        CHAIN_ID = chainid_;
        COIN_TYPE = coinType_;
        GAS_LIMIT = gasLimit_;
        SYSTEM_CONTRACT_ADDRESS = systemContractAddress_;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~ ZRC20 Standard Functions ~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    function name() public view override returns (string memory) {
        return _name;
    }

    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function decimals() public view override returns (uint8) {
        return _decimals;
    }

    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) public view override returns (uint256) {
        return _balances[account];
    }

    function transfer(address recipient, uint256 amount) public override returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function allowance(address owner, address spender) public view override returns (uint256) {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 amount) public override returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public override returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][msg.sender];
        if (currentAllowance < amount) revert LowAllowance();

        _approve(sender, msg.sender, currentAllowance - amount);
        return true;
    }

    /**
     * @dev Burns tokens from sender’s balance.
     */
    function burn(uint256 amount) external override returns (bool) {
        _burn(msg.sender, amount);
        return true;
    }

    /**
     * @dev Mints tokens (e.g., deposit from an external chain).
     */
    function deposit(address to, uint256 amount) external override returns (bool) {
        // if (msg.sender != FUNGIBLE_MODULE_ADDRESS && msg.sender != SYSTEM_CONTRACT_ADDRESS) {
        //     revert InvalidSender();
        // }
        // (commented out for local dev)
        _mint(to, amount);
        emit Deposit(abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), to, amount);
        return true;
    }

    /**
     * @dev For local PoC, we skip the real system contract logic so this reverts.
     *      Do not call it in your demo; it’s left here for completeness.
     */
    function withdrawGasFee() public view override returns (address, uint256) {
        revert ZeroGasCoin(); 
    }

    /**
     * @dev Rate-limited withdrawal override.
     */
    function withdraw(bytes memory to, uint256 amount) external override returns (bool) {
        // 1. Rate-limiting check
        _checkAndUpdateWithdrawLimits(msg.sender, amount);

        // 2. Skip gas fee logic in local PoC
        // e.g. (address gasZRC20, uint256 gasFee) = withdrawGasFee();
        // ...

        // 3. Burn user’s tokens
        _burn(msg.sender, amount);

        // 4. Emit event with dummy gasFee = 0
        emit Withdrawal(msg.sender, to, amount, 0, PROTOCOL_FLAT_FEE);
        return true;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~ Internal ZRC20 Helpers ~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    function _transfer(address sender, address recipient, uint256 amount) internal {
        if (sender == address(0)) revert ZeroAddress();
        if (recipient == address(0)) revert ZeroAddress();

        uint256 senderBalance = _balances[sender];
        if (senderBalance < amount) revert LowBalance();

        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    function _mint(address account, uint256 amount) internal {
        if (account == address(0)) revert ZeroAddress();

        _totalSupply += amount;
        _balances[account] += amount;

        emit Transfer(address(0), account, amount);
    }

    function _burn(address account, uint256 amount) internal {
        if (account == address(0)) revert ZeroAddress();

        uint256 accountBalance = _balances[account];
        if (accountBalance < amount) revert LowBalance();

        _balances[account] = accountBalance - amount;
        _totalSupply -= amount;

        emit Transfer(account, address(0), amount);
    }

    function _approve(address owner, address spender, uint256 amount) internal {
        if (owner == address(0)) revert ZeroAddress();
        if (spender == address(0)) revert ZeroAddress();

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~ Rate-Limiting Logic ~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    function _checkAndUpdateWithdrawLimits(address user, uint256 amount) internal {
        WithdrawLimit storage limits = _withdrawInfo[user];

        // If the daily window ended, reset usage
        if (block.timestamp >= limits.windowStart + WINDOW_SIZE) {
            limits.count = 0;
            limits.volume = 0;
            limits.windowStart = block.timestamp;
        }

        // Check count
        if (limits.count + 1 > maxWithdrawCountPerDay) {
            revert("RateLimitedZRC20: Exceeded max daily withdrawal count");
        }

        // Check volume
        if (limits.volume + amount > maxWithdrawVolumePerDay) {
            revert("RateLimitedZRC20: Exceeded max daily withdrawal volume");
        }

        // Update usage
        limits.count += 1;
        limits.volume += amount;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~ Admin Functions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    function setMaxWithdrawCountPerDay(uint256 _newCount) external onlyFungible {
        maxWithdrawCountPerDay = _newCount;
    }

    function setMaxWithdrawVolumePerDay(uint256 _newVolume) external onlyFungible {
        maxWithdrawVolumePerDay = _newVolume;
    }

    function updateSystemContractAddress(address addr) external onlyFungible {
        SYSTEM_CONTRACT_ADDRESS = addr;
        emit UpdatedSystemContract(addr);
    }

    function updateGasLimit(uint256 gasLimit) external onlyFungible {
        GAS_LIMIT = gasLimit;
        emit UpdatedGasLimit(gasLimit);
    }

    function updateProtocolFlatFee(uint256 protocolFlatFee) external onlyFungible {
        PROTOCOL_FLAT_FEE = protocolFlatFee;
        emit UpdatedProtocolFlatFee(protocolFlatFee);
    }
}