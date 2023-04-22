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
        require(_beforeFulfill(request), "KlayOracle: subscribe to DP");    //Provider can perform: charge

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
        bytes32 roundAnswer,
        bytes memory signature
    ) external returns (bool) {
        require(msg.sender == nodeAddress, "Oracle: node unknown");

        bytes32 msgHash = _getHash(roundAnswer, roundTime);

        (bytes32 r, bytes32 s, uint8 v) = splitSignature(signature);
        address signer = VerifyKlaytnMessage(msgHash, v, r, s);

        require(signer == nodeAddress, "Oracle: Invalid data");

        Round memory round = Round(roundAnswer, roundTime, block.timestamp);

        rounds.push(round);

        latestRound = round;
        latestResponse = roundAnswer;

        return true;
    }

    function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) public pure returns (address) {
        bytes memory prefix = "\x19Klaytn Signed Message:\n32";
        bytes32 prefixedHashMessage = keccak256(abi.encodePacked(prefix, _hashedMessage));
        address signer = ecrecover(prefixedHashMessage, _v, _r, _s);
        return signer;
    }

    function _getHash(bytes32 roundAnswer, uint256 roundTime) public view returns (bytes32) {
        return keccak256(abi.encodePacked(msg.sender, roundAnswer, roundTime));
    }

    function splitSignature(bytes memory sig) public pure returns (bytes32 r, bytes32 s, uint8 v) {
        require(sig.length == 65, "invalid signature length");

        assembly {
        /*
        First 32 bytes stores the length of the signature

        add(sig, 32) = pointer of sig + 32
        effectively, skips first 32 bytes of signature

        mload(p) loads next 32 bytes starting at the memory address p into memory
        */

        // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
        // second 32 bytes
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        if(v == 0){v = 27;}
        if(v == 1){v = 28;}

        return (r, s, v);
    }

    function _isWhitelisted(address _address) internal virtual returns (bool);

    function _beforeFulfill(
        Request memory request
    ) internal virtual returns (bool);
}
