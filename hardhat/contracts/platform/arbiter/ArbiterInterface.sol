pragma solidity ^0.5.1;

contract ArbiterInterface {
    function initiateSubscription(address, bytes32, bytes32[] memory, uint256, uint64) public;
    function getSubscription(address, address, bytes32) public view returns (uint64, uint96, uint96);
    function endSubscriptionProvider(address, bytes32) public;
    function endSubscriptionSubscriber(address, bytes32) public;
    function passParams(address receiver, bytes32 endpoint, bytes32[] memory params) public;
}
