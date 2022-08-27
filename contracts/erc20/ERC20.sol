// SPDX-License-Identifier: MIT
pragma solidity >=0.8.15;


/// @title Standard ERC20 interface
interface IERC20 {
    function totalSupply() external view returns (uint);

    function balanceOf(address account) external view returns (uint);

    function transfer(address recipient, uint amount) external returns (bool);

    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint amount) external returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint amount
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint value);
    event Approval(address indexed owner, address indexed spender, uint value);
}

/**
 * @title A simulator for trees
 * @author Shiun
 * @dev Implementation of the {IERC20} interface.
 *
 * This implementation is agnostic to the way tokens are created. This means
 * that a supply mechanism has to be added in a derived contract using {mint}.
 *
 * Additionally, an {Approval} event is emitted on calls to {transferFrom}.
 * This allows applications to reconstruct the allowance for all accounts just
 * by listening to said events. Other implementations of the EIP may not emit
 * these events, as it isn't required by the specification.
 */
contract ERC20 is IERC20 {
    uint public totalSupply;
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;
    address internal owner;
    string public symbol = "CVM";
    uint8 public decimals = 18;

    /**
     * @dev Sets the values for {owner} and mint 1000 tokens to the msg.sender.
     *
     * The default value of {decimals} is 18. To select a different value for
     * {decimals} you should overload it.
     *
     * These values are immutable: they can only be set once during
     * construction.
     */
    constructor() {
        owner = msg.sender;

        // Mint 1000 tokens to msg.sender
        // 1 token = 1 * (10 ** decimals)
        mint(1000 * 10**uint(decimals));
    }

    /**
     * @dev Restricts the user, only contract deployer is able to use the function
     * which has this restriction.
     */
    modifier OnlyWoner() {
        require(
            owner == msg.sender,
            "You are not the contract owner"
        );
        _;
    }

    /**
     * @dev trans amount of tokens from msg.sender to recipient account.
     *
     * Requirements:
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(address recipient, uint amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }

    /**
     * @dev See {IERC20-approve}.
     *
     * @notice: this function still have issue!
     *
     * Requirements:
     * - `spender` cannot be the zero address.
     */
    function approve(address spender, uint amount) external returns (bool) {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    /**
     * @dev transfer tokens from sender account to recipient account.
     *
     * NOTE: before using this function, sender must first use approve function
     * to set the allowance.
     *
     * Requirements:
     * - `sender` must have a balance of at least `amount`.
     * - the caller must have allowance for ``from``'s tokens of at least
     * `amount`.
     */
    function transferFrom(address sender, address recipient, uint amount) external returns (bool) {
        allowance[sender][msg.sender] -= amount;
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }

    /** @dev Creates `amount` tokens and assigns them to msg.sender,
     * it increases the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     */
    function mint(uint amount) public {
        balanceOf[msg.sender] += amount;
        totalSupply += amount;
        emit Transfer(address(0), msg.sender, amount);
    }

    /**
     * @dev Destroys `amount` tokens from the sender, reducing the
     * total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     */
    function burn(uint amount) external OnlyWoner {
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        emit Transfer(msg.sender, address(0), amount);
    }
}