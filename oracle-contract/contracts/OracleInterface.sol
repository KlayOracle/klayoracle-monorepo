// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface OracleInterface {
    struct Request {
        bytes32 requestId;
        address nodeAddress;
        string adapterId;
        bytes4 callbackFunctionId;
        address callBackContract;
        bytes32 data;
    }

    /**
     * @dev This function is called by the node to fulfill a request
     * @param requestId The request id
     * @param data The data to be returned
     * @param dataHash The hash of the signed data
     * @param signature The signature of the signed data
     */
    function fulfillOracleRequest(
        bytes32 requestId,
        bytes32 data,
        bytes32 dataHash,
        bytes memory signature
    ) external returns (bool);

    /**
     * @dev This function is called by the consumer contract to request data from the node
     * @param callbackFunctionId bytes4 of the consumer contract callback function
     * @param adapterID The adapater id defined by the data provider
     * @param callBackContract The contract with the callback function requesting offchain data
     */
    function newOracleRequest(
        bytes4 callbackFunctionId,
        string memory adapterID,
        address callBackContract
    ) external returns (bool);

    /**
     * @dev This function is called by the consumer contract to request data from the node
     * @param requestId The request id
     */
    function getOracleRequest(
        bytes32 requestId
    ) external view returns (Request memory);

    function 
}
