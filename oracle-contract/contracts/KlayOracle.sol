// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "./KlayOracleInterface.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

abstract contract KlayOracle is KlayOracleInterface {
    using ECDSA for bytes32;

    string public constant VERSION = "v1.0.0"; // Node can use this for compatibility checks

    address public immutable nodeAddress;

    bytes32 public immutable adapterId;

    //Requests variables
    mapping(bytes32 => Request) public requests;

    uint256 public fulfilledCount = 0;

    bytes32 private latestResponse; // solhint-disable-line private-vars-leading-underscore

    bytes32 private latestRequestId; // solhint-disable-line private-vars-leading-underscore

    event NewOracleRequest(bytes32 requestId);

    // Round variables
    struct Round {
        bytes32 answer;
        uint256 roundTime;
        uint256 timestamp; //The timestamp of the round
    }

    Round[] public rounds;

    Round public latestRound;

    constructor(address _nodeAddress, bytes32 _adapterId) {
        nodeAddress = _nodeAddress;
        adapterId = _adapterId;
    }

    //Called by the consumer contract
    function newOracleRequest(
        bytes4 callbackFunctionId,
        address callBackContract
    ) external override returns (bool) {
        require(
            _isWhitelisted(msg.sender),
            "Oracle: consumer is not whitelisted"
        );

        bytes32 requestId = keccak256(
            abi.encodePacked(
                latestRequestId,
                nodeAddress,
                adapterId,
                callbackFunctionId,
                callBackContract,
                latestResponse,
                block.timestamp
            )
        );

        Request memory request = Request(
            requestId,
            nodeAddress,
            adapterId,
            callbackFunctionId,
            callBackContract,
            latestRound.answer,
            block.timestamp
        );

        emit NewOracleRequest(requestId);

        bool fulfilled = _fulfill(requestId, request);

        require(fulfilled, "KlayOracle: request not fulfilled");

        return true;
    }

    function _fulfill(
        bytes32 requestId,
        Request memory request
    ) internal returns (bool) {
        require(_beforeFulfill(request), "KlayOracle: subscribe to DP"); //Provider can perform: charge

        requests[requestId] = request;

        fulfilledCount++;

        (bool success, ) = request.callBackContract.call(
            abi.encodePacked(request.callbackFunctionId, request.data)
        );

        return success;
    }

    //Called by the node to submit latest data
    function newRoundData(
        uint256 roundTime, //milliseconds
        bytes32 data,
        bytes32 dataHash,
        bytes memory signature
    ) external returns (bool) {
        require(msg.sender == nodeAddress, "Oracle: node unknown");

        address signer = dataHash.recover(signature);

        require(signer == nodeAddress, "Oracle: Invalid data");

        Round memory round = Round(data, roundTime, block.timestamp);

        rounds.push(round);

        latestRound = round;
        latestResponse = data;

        return true;
    }

    function _isWhitelisted(address _address) internal virtual returns (bool);

    function _beforeFulfill(
        Request memory request
    ) internal virtual returns (bool);
}
