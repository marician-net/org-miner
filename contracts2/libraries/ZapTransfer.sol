pragma solidity >=0.5.0 <0.7.0;

import "./SafeMath.sol";
import "./BerryStorage.sol";

/**
* @title Berry Transfer
* @dev Contains the methods related to transfers and ERC20. Berry.sol and BerryGetters.sol
* reference this library for function's logic.
*/
library BerryTransfer {
    using SafeMath for uint256;

    event Approval(address indexed _owner, address indexed _spender, uint256 _value); //ERC20 Approval event
    event Transfer(address indexed _from, address indexed _to, uint256 _value); //ERC20 Transfer Event

    /*Functions*/

    /**
    * @dev Allows for a transfer of tokens to _to
    * @param _to The address to send tokens to
    * @param _amount The amount of tokens to send
    */
    function transfer(BerryStorage.BerryStorageStruct storage self, address _to, uint256 _amount) public returns (bool success) {
        doTransfer(self, msg.sender, _to, _amount);
        return true;
    }

    /**
    * @notice Send _amount tokens to _to from _from on the condition it
    * is approved by _from
    * @param _from The address holding the tokens being transferred
    * @param _to The address of the recipient
    * @param _amount The amount of tokens to be transferred
    */
    function transferFrom(BerryStorage.BerryStorageStruct storage self, address _from, address _to, uint256 _amount)
        public
        returns (bool success)
    {
        require(self.allowed[_from][msg.sender] >= _amount, "Allowance is wrong");
        self.allowed[_from][msg.sender] -= _amount;
        doTransfer(self, _from, _to, _amount);
        return true;
    }

    /**
    * @dev This function approves a _spender an _amount of tokens to use
    * @param _spender address
    * @param _amount amount the spender is being approved for
    * @return true if spender appproved successfully
    */
    function approve(BerryStorage.BerryStorageStruct storage self, address _spender, uint256 _amount) public returns (bool) {
        require(_spender != address(0), "Spender is 0-address");
        self.allowed[msg.sender][_spender] = _amount;
        emit Approval(msg.sender, _spender, _amount);
        return true;
    }

    /**
    * @param _user address of party with the balance
    * @param _spender address of spender of parties said balance
    * @return Returns the remaining allowance of tokens granted to the _spender from the _user
    */
    function allowance(BerryStorage.BerryStorageStruct storage self, address _user, address _spender) public view returns (uint256) {
        return self.allowed[_user][_spender];
    }

    /**
    * @dev Completes POWO transfers by updating the balances on the current block number
    * @param _from address to transfer from
    * @param _to addres to transfer to
    * @param _amount to transfer
    */
    function doTransfer(BerryStorage.BerryStorageStruct storage self, address _from, address _to, uint256 _amount) public {
        require(_amount > 0, "Tried to send non-positive amount");
        require(_to != address(0), "Receiver is 0 address");
        //allowedToTrade checks the stakeAmount is removed from balance if the _user is staked
        require(allowedToTrade(self, _from, _amount), "Stake amount was not removed from balance");
        uint256 previousBalance = balanceOfAt(self, _from, block.number);
        updateBalanceAtNow(self.balances[_from], previousBalance - _amount);
        previousBalance = balanceOfAt(self, _to, block.number);
        require(previousBalance + _amount >= previousBalance, "Overflow happened"); // Check for overflow
        updateBalanceAtNow(self.balances[_to], previousBalance + _amount);
        emit Transfer(_from, _to, _amount);
    }

    /**
    * @dev Gets balance of owner specified
    * @param _user is the owner address used to look up the balance
    * @return Returns the balance associated with the passed in _user
    */
    function balanceOf(BerryStorage.BerryStorageStruct storage self, address _user) public view returns (uint256) {
        return balanceOfAt(self, _user, block.number);
    }

    /**
    * @dev Queries the balance of _user at a specific _blockNumber
    * @param _user The address from which the balance will be retrieved
    * @param _blockNumber The block number when the balance is queried
    * @return The balance at _blockNumber specified
    */
    function balanceOfAt(BerryStorage.BerryStorageStruct storage self, address _user, uint256 _blockNumber) public view returns (uint256) {
        if ((self.balances[_user].length == 0) || (self.balances[_user][0].fromBlock > _blockNumber)) {
            return 0;
        } else {
            return getBalanceAt(self.balances[_user], _blockNumber);
        }
    }

    /**
    * @dev Getter for balance for owner on the specified _block number
    * @param checkpoints gets the mapping for the balances[owner]
    * @param _block is the block number to search the balance on
    * @return the balance at the checkpoint
    */
    function getBalanceAt(BerryStorage.Checkpoint[] storage checkpoints, uint256 _block) public view returns (uint256) {
        if (checkpoints.length == 0) return 0;
        if (_block >= checkpoints[checkpoints.length - 1].fromBlock) return checkpoints[checkpoints.length - 1].value;
        if (_block < checkpoints[0].fromBlock) return 0;
        // Binary search of the value in the array
        uint256 min = 0;
        uint256 max = checkpoints.length - 1;
        while (max > min) {
            uint256 mid = (max + min + 1) / 2;
            if (checkpoints[mid].fromBlock <= _block) {
                min = mid;
            } else {
                max = mid - 1;
            }
        }
        return checkpoints[min].value;
    }

    /**
    * @dev This function returns whether or not a given user is allowed to trade a given amount
    * and removing the staked amount from their balance if they are staked
    * @param _user address of user
    * @param _amount to check if the user can spend
    * @return true if they are allowed to spend the amount being checked
    */
    function allowedToTrade(BerryStorage.BerryStorageStruct storage self, address _user, uint256 _amount) public view returns (bool) {
        if (self.stakerDetails[_user].currentStatus > 0 && self.stakerDetails[_user].currentStatus < 4) {
            //Removes the stakeAmount from balance if the _user is staked
            if (balanceOf(self, _user).sub(self.uintVars[keccak256("stakeAmount")]).sub(_amount) >= 0) {
                return true;
            }
        } else if (balanceOf(self, _user).sub(_amount) >= 0) {
            return true;
        }
        return false;
    }

    /**
    * @dev Updates balance for from and to on the current block number via doTransfer
    * @param checkpoints gets the mapping for the balances[owner]
    * @param _value is the new balance
    */
    function updateBalanceAtNow(BerryStorage.Checkpoint[] storage checkpoints, uint256 _value) public {
        if ((checkpoints.length == 0) || (checkpoints[checkpoints.length - 1].fromBlock < block.number)) {
            BerryStorage.Checkpoint storage newCheckPoint = checkpoints[checkpoints.length + 1];
            newCheckPoint.fromBlock = uint128(block.number);
            newCheckPoint.value = uint128(_value);
        } else {
            BerryStorage.Checkpoint storage oldCheckPoint = checkpoints[checkpoints.length - 1];
            oldCheckPoint.value = uint128(_value);
        }
    }
}
