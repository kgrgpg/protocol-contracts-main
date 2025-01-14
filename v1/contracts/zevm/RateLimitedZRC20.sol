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
    /// @notice System contract address.
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

    /**
     * @dev Custom error definitions replicated from original ZRC20 for brevity
     */
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

    /**
     * @dev A struct that tracks user’s withdrawal usage within a daily window
     *  - count: how many times user withdrew in the current window
     *  - volume: how many tokens user withdrew in the current window
     *  - windowStart: when the current daily window started
     */
    struct WithdrawLimit {
        uint256 count;
        uint256 volume;
        uint256 windowStart;
    }

    /// @dev Mapping from user => their withdrawal usage
    mapping(address => WithdrawLimit) private _withdrawInfo;

    /// @dev The “daily” time window, in seconds (24h = 86400)
    uint256 public constant WINDOW_SIZE = 86400;

    /// @notice Maximum times a user can withdraw in a single 24-hour window
    uint256 public maxWithdrawCountPerDay = 3;

    /// @notice Maximum total volume a user can withdraw in a single 24-hour window
    uint256 public maxWithdrawVolumePerDay = 1000 ether;

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Modifiers ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    /**
     * @dev Only Fungible Module can call certain functions (admin-type).
     *      This is the same as the original ZRC20’s version.
     */
    modifier onlyFungible() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        _;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~ Constructor ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    /**
     * @dev On deployment, we configure standard ZRC20 parameters. 
     *      Only the fungible module can deploy new ZRC20 tokens.
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 chainid_,
        CoinType coinType_,
        uint256 gasLimit_,
        address systemContractAddress_
    ) {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();

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
     * @dev Mints tokens (e.g., representing a deposit from an external chain).
     */
    function deposit(address to, uint256 amount) external override returns (bool) {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS && msg.sender != SYSTEM_CONTRACT_ADDRESS) {
            revert InvalidSender();
        }
        _mint(to, amount);
        emit Deposit(abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), to, amount);
        return true;
    }

    /**
     * @dev Returns the gas fee for a withdrawal, including the address of the chain’s
     *      gas ZRC20 token and the total computed fee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE.
     */
    function withdrawGasFee() public view override returns (address, uint256) {
        address gasZRC20 = ISystem(SYSTEM_CONTRACT_ADDRESS).gasCoinZRC20ByChainId(CHAIN_ID);
        if (gasZRC20 == address(0)) revert ZeroGasCoin();

        uint256 gasPrice = ISystem(SYSTEM_CONTRACT_ADDRESS).gasPriceByChainId(CHAIN_ID);
        if (gasPrice == 0) revert ZeroGasPrice();

        uint256 gasFee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE;
        return (gasZRC20, gasFee);
    }

    /**
     * @dev Rate-limited withdrawal of ZRC20 tokens to external chain.
     *      We override the original withdraw() to inject `_checkAndUpdateWithdrawLimits(...)`.
     */
    function withdraw(bytes memory to, uint256 amount) external override returns (bool) {
        // 1. Perform rate-limiting checks
        _checkAndUpdateWithdrawLimits(msg.sender, amount);

        // 2. Pay gas fee in gasZRC20
        (address gasZRC20, uint256 gasFee) = withdrawGasFee();
        bool feeSuccess = IZRC20(gasZRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee);
        if (!feeSuccess) revert GasFeeTransferFailed();

        // 3. Burn the user’s tokens
        _burn(msg.sender, amount);

        // 4. Emit event
        emit Withdrawal(msg.sender, to, amount, gasFee, PROTOCOL_FLAT_FEE);
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

    /**
     * @dev Internal function that checks if the user can withdraw
     *      based on their daily usage. If their daily window is expired,
     *      reset it. Then ensure they do not exceed daily count or volume.
     */
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

        // Update tracking
        limits.count += 1;
        limits.volume += amount;
    }

    // ------------------------------------------------------------------------
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~ Admin Functions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ------------------------------------------------------------------------

    /**
     * @dev (Optional) Only the fungible module can update the max daily withdraw count.
     */
    function setMaxWithdrawCountPerDay(uint256 _newCount) external onlyFungible {
        maxWithdrawCountPerDay = _newCount;
    }

    /**
     * @dev (Optional) Only the fungible module can update the max daily withdraw volume.
     */
    function setMaxWithdrawVolumePerDay(uint256 _newVolume) external onlyFungible {
        maxWithdrawVolumePerDay = _newVolume;
    }

    /**
     * @dev Only Fungible Module can update the system contract address.
     */
    function updateSystemContractAddress(address addr) external onlyFungible {
        SYSTEM_CONTRACT_ADDRESS = addr;
        emit UpdatedSystemContract(addr);
    }

    /**
     * @dev Only Fungible Module can update the gas limit.
     */
    function updateGasLimit(uint256 gasLimit) external onlyFungible {
        GAS_LIMIT = gasLimit;
        emit UpdatedGasLimit(gasLimit);
    }

    /**
     * @dev Only Fungible Module can update the protocol flat fee.
     */
    function updateProtocolFlatFee(uint256 protocolFlatFee) external onlyFungible {
        PROTOCOL_FLAT_FEE = protocolFlatFee;
        emit UpdatedProtocolFlatFee(protocolFlatFee);
    }
}
