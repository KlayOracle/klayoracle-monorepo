// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "./KlayOracleInterface.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "hardhat/console.sol";

abstract contract KlayOracle is KlayOracleInterface {
    using ECDSA for bytes32;

    string constant VERSION = "v1.0.0"; // Node can use this for compatibility checks

    address public immutable nodeAddress;

    bytes32 public immutable adapterId;

    mapping(bytes32 => Request) requests;

    uint256 public fulfilledCount = 0;

    event NewOracleRequest(bytes32 requestId);

    constructor(address _nodeAddress, bytes32 _adapterId) {
        nodeAddress = _nodeAddress;
        adapterId = _adapterId;
    }

    function newOracleRequest(
        bytes4 callbackFunctionId,
        bytes32 adapterID,
        address callBackContract
    ) external override returns (bool) {
        require(
            _isWhitelisted(msg.sender),
            "Oracle: Permission needed for consumer"
        );

        bytes32 requestId = keccak256(
            abi.encodePacked(
                nodeAddress,
                callbackFunctionId,
                callBackContract,
                adapterID,
                block.timestamp,
                fulfilledCount
            )
        );

        console.logBytes32(requestId);

        requests[requestId] = Request(
            requestId,
            nodeAddress,
            adapterID,
            callbackFunctionId,
            callBackContract,
            bytes32("0x"),
            0
        );

        emit NewOracleRequest(requestId);

        return true;
    }

    function fulfillOracleRequest(
        bytes32 requestId,
        bytes32 data,
        bytes memory signature
    ) external returns (bool) {
        require(
            msg.sender == nodeAddress,
            "Oracle: Permission needed for node"
        );

        address signer = data.recover(signature);

        require(signer == nodeAddress, "Oracle: Invalid data");

        Request storage request = requests[requestId];
        request.timestamp = block.timestamp;
        request.data = data;

        require(_beforeFulfill(), "Oracle: Permission needed for consumer");

        fulfilledCount++;

        (bool success, ) = request.callBackContract.call(
            abi.encode(request.callbackFunctionId, data)
        );

        return success;
    }

    function getOracleRequest(
        bytes32 requestId
    ) external view returns (Request memory) {
        return requests[requestId];
    }

    function _isWhitelisted(address _address) virtual internal returns (bool);

    function _beforeFulfill() virtual internal returns (bool);
}
