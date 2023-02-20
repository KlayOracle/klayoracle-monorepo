// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "./KlayOracle.sol";

contract OracleSample is KlayOracle {
    constructor(
        address _nodeAddress,
        bytes32 _adapterId
    ) KlayOracle(_nodeAddress, _adapterId) {}

    function _isWhitelisted(
        address _address
    ) internal pure override returns (bool) {
        require(_address == address(0));
        return true;
    }

    function _beforeFulfill() internal pure override returns (bool) {
        return true;
    }
}
