// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "./KlayOracleInterface.sol";
import "hardhat/console.sol";

contract OracleConsumer {
    address public immutable oracleAddress;

    uint256 public klayOutput;

    constructor(address _oracleAddress) {
        require(_oracleAddress != address(0));
        oracleAddress = _oracleAddress;
    }

    function swapEthtoKlay() public returns (bool) {
        KlayOracleInterface oracle = KlayOracleInterface(oracleAddress);

        bool replied = oracle.newOracleRequest(
            this.swap.selector,
            address(this)
        );

        return replied;
    }

    function swap(uint256 _klayOutput) public {
        klayOutput = _klayOutput;

        console.logUint(klayOutput);
        //Swap eth to klay
    }
}
