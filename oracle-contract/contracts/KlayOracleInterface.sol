// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface KlayOracleInterface {
    struct Request {
        bytes32 requestId;
        address nodeAddress;
        bytes32 adapterId;
        bytes4 callbackFunctionId;
        address callBackContract;
        bytes32 data;
        uint256 timestamp;
    }

    /**
     * @dev This function is called by the node to fulfill a request
     * @param roundTime Time to aggregate the data
     * @param data The data to be returned
     * @param dataHash UTF-8 bytes, prefixed with \x19Ethereum Signed Message\n32
     * @param signature The signature of the signed data
     */
    function newRoundData(
        uint256 roundTime,
        bytes32 data,
        bytes32 dataHash,
        bytes memory signature
    ) external returns (bool);

    /**
     * @dev This function is called by the consumer contract to request data from the node
     * @param callbackFunctionId bytes4 of the consumer contract callback function
     * @param callBackContract The contract with the callback function requesting offchain data
     */
    function newOracleRequest(
        bytes4 callbackFunctionId,
        address callBackContract
    ) external returns (bool);
}
