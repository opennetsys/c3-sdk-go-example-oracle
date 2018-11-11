pragma solidity ^0.4.23;

import "./SafeMath.sol";
import "./Ownable.sol";

contract Exchange is Ownable {
  using SafeMath for uint256;

  event LogDeposit(address sender, uint256 value);
  event LogBuy(address sender, uint256 amount, uint256 price, uint256 value);
  event LogWithdrawal(address receiver, uint256 value);

  mapping (address => uint256) deposits;

  constructor(address _owner) public {
    require(_owner != address(0));
    transferOwnership(_owner);
  }

  /// amount of EOS to buy and at what price
  function deposit() public payable {
    deposits[msg.sender].add(msg.value);

    emit LogDeposit(msg.sender, msg.value);
  }

  /// amount of EOS to buy and at what price
  function buy(uint256 amount, uint256 price) public payable {
    deposits[msg.sender].add(msg.value);

    emit LogBuy(msg.sender, amount, price, msg.value);
  }

  function withdraw(address receiver, uint256 value) onlyOwner public {
    require(deposits[receiver] >= value);
    deposits[receiver].sub(value);

    emit LogWithdrawal(receiver, value);
  }
}
