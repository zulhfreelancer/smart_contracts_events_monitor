pragma solidity ^0.4.23;

import "./zeppelin/token/ERC20/MintableToken.sol";

contract ABCtoken is MintableToken {
  string public constant name = "ABCtoken";
  string public constant symbol = "ABC";
  uint8 public constant decimals = 18;

  constructor() public {
    // make the deployer rich
    balances[msg.sender] = 1000 * 1 ether;
  }
}
