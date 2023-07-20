// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "https://github.com/KlayOracle/klayoracle-monorepo/blob/development/oracle-contract/contracts/KlayOracleInterface.sol";

contract OracleConsumerSample {
    address public immutable oracleAddress;

    uint256 public klayOutput;

    constructor(address _oracleAddress) {
        require(_oracleAddress != address(0));
        oracleAddress = _oracleAddress;
    }

    function swapUsdtoKlay() public returns (bool) {
        KlayOracleInterface oracle = KlayOracleInterface(oracleAddress);

        bool replied = oracle.newOracleRequest(
            this.swap.selector,
            address(this)
        );

        return replied;
    }

    function swap(uint256 _klayOutput) external {
        require(msg.sender == oracleAddress, "not allowed");    //ensure only Oracle contract can set price
        klayOutput = _klayOutput;
        //Swap usd to klay
    }
}