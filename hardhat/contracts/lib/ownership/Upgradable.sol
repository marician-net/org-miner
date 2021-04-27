pragma solidity ^0.5.1;

import "./ZapCoordinatorInterface.sol";

contract Upgradable {

    address coordinatorAddr;
    ZapCoordinatorInterface coordinator;

    constructor(address c) public{
        coordinatorAddr = c;
        coordinator = ZapCoordinatorInterface(c);
    }

    function updateDependencies() external coordinatorOnly {
        _updateDependencies();
    }

    function _updateDependencies() internal;

    modifier coordinatorOnly() {
        require(msg.sender == coordinatorAddr, "Error: Coordinator Only Function");
        _;
    }
}
