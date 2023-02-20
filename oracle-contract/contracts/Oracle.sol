// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "./OracleInterface.sol";

contract Oracle is OracleInterface {

    string constant VERSION = "1.0.0";

    address public nodeAddress;

    string public adapterId;

    mapping (bytes32 => Request) requests;

    event NewOracleRequest(string adapterId, bytes32 requestId);

    constructor(address _nodeAddress,string memory _adapterId) {
        nodeAddress = _nodeAddress;
        adapterId = _adapterId;
    }

    function newOracleRequest(
        bytes4 callbackFunctionId,
        string memory adapterID,
        address callBackContract
    ) external override returns(bool) {
        bytes32 requestId = keccak256(abi.encodePacked(nodeAddress, callbackFunctionId, callBackContract, adapterID));

        requests[requestId] = Request(
            requestId,
            nodeAddress,
            adapterID,
            callbackFunctionId,
            callBackContract,
            bytes32("")
        );

        emit NewOracleRequest(adapterID, requestId);

        return true;
    }

    function fulfillOracleRequest(
        bytes32 requestId,
        bytes32 data,
        bytes32  messageHash, 
        bytes memory signature
    ) external override returns(bool) {

        require(msg.sender == nodeAddress,"Oracle: Permission needed for node");

        Request memory request = requests[requestId];

        (bool success, ) = request.callBackContract.call(abi.encodeWithSelector(request.callbackFunctionId, data));

        return success;
    }

}