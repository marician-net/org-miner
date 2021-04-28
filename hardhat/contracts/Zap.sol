pragma solidity ^0.5.0;

// import "./libraries/SafeMath.sol";
import "./libraries/ZapStorage.sol";
// import "./libraries/ZapTransfer.sol";
import "./libraries/ZapDispute.sol";
import "./libraries/ZapStake.sol";
import "./libraries/ZapLibrary.sol";
// import "./libraries/Upgradable.sol";
import "./token/ZapToken.sol";

/**
 * @title Zap Oracle System
 * @dev Oracle contract where miners can submit the proof of work along with the value.
 * The logic for this contract is in ZapLibrary.sol, ZapDispute.sol, ZapStake.sol, 
 * and ZapTransfer.sol
 */
contract Zap {

    using SafeMath for uint256;

    using ZapDispute for ZapStorage.ZapStorageStruct;
    using ZapLibrary for ZapStorage.ZapStorageStruct;
    using ZapStake for ZapStorage.ZapStorageStruct;
    using ZapTransfer for ZapStorage.ZapStorageStruct;

    ZapStorage.ZapStorageStruct zap;
    ZapToken public token;

    constructor (address zapToken) public {
        token = ZapToken(zapToken);
    }

    function balanceOf(address _user) public view returns (uint256 balance){
        return token.balanceOf(_user);
    }

    /*Functions*/
    
    /*This is a cheat for demo purposes, will delete upon actual launch*/
   function theLazyCoon(address _address, uint _amount) public {
        zap.theLazyCoon(_address,_amount);
    }


    /**
    * @dev Helps initialize a dispute by assigning it a disputeId 
    * when a miner returns a false on the validate array(in Zap.ProofOfWork) it sends the 
    * invalidated value information to POS voting
    * @param _requestId being disputed
    * @param _timestamp being disputed
    * @param _minerIndex the index of the miner that submitted the value being disputed. Since each official value 
    * requires 5 miners to submit a value.
    */
    function beginDispute(uint _requestId, uint _timestamp,uint _minerIndex) external {
        zap.beginDispute(_requestId,_timestamp,_minerIndex);
    }


    /**
    * @dev Allows token holders to vote
    * @param _disputeId is the dispute id
    * @param _supportsDispute is the vote (true=the dispute has basis false = vote against dispute)
    */
    function vote(uint _disputeId, bool _supportsDispute) external {
        zap.vote(_disputeId,_supportsDispute);
    }


    /**
    * @dev tallies the votes.
    * @param _disputeId is the dispute id
    */
    function tallyVotes(uint _disputeId) external {
        zap.tallyVotes(_disputeId);
    }


    /**
    * @dev Allows for a fork to be proposed
    * @param _propNewZapAddress address for new proposed Zap
    */
    function proposeFork(address _propNewZapAddress) external {
        zap.proposeFork(_propNewZapAddress);
    }


   /**
    * @dev Add tip to Request value from oracle
    * @param _requestId being requested to be mined
    * @param _tip amount the requester is willing to pay to be get on queue. Miners
    * mine the onDeckQueryHash, or the api with the highest payout pool
    */
    function addTip(uint _requestId, uint _tip) external {
        zap.addTip(_requestId,_tip);
    }


    /**
    * @dev Request to retreive value from oracle based on timestamp. The tip is not required to be 
    * greater than 0 because there are no tokens in circulation for the initial(genesis) request 
    * @param _c_sapi string API being requested be mined
    * @param _c_symbol is the short string symbol for the api request
    * @param _granularity is the number of decimals miners should include on the submitted value
    * @param _tip amount the requester is willing to pay to be get on queue. Miners
    * mine the onDeckQueryHash, or the api with the highest payout pool
    */
    function requestData(string calldata _c_sapi,string calldata _c_symbol,uint _granularity, uint _tip) external {
        zap.requestData(_c_sapi,_c_symbol,_granularity,_tip);
    }


    /**
    * @dev Proof of work is called by the miner when they submit the solution (proof of work and value)
    * @param _nonce uint submitted by miner
    * @param _requestId the apiId being mined
    * @param _value of api query
    */
    function submitMiningSolution(string calldata _nonce, uint _requestId, uint _value) external{
        zap.submitMiningSolution(_nonce,_requestId,_value);
    }


    /**
    * @dev Allows the current owner to transfer control of the contract to a newOwner.
    * @param _newOwner The address to transfer ownership to.
    */
    // function transferOwnership(address payable _newOwner) public {
    //     // zap.transferOwnership(_newOwner);
    //     token.transferOwnership(_newOwner);
    // }


    /**
    * @dev This function allows miners to deposit their stake.
    */
    function depositStake() external {
        // require balance is >= here before it hits NewStake()
        require(token.balanceOf(msg.sender) >= zap.uintVars[keccak256("stakeAmount")]);
        zap.depositStake();
    }


    /**
    * @dev This function allows stakers to request to withdraw their stake (no longer stake) 
    * once they lock for withdraw(stakes.currentStatus = 2) they are locked for 7 days before they 
    * can withdraw the stake
    */
    function requestStakingWithdraw() external {
        zap.requestStakingWithdraw();
    }


    /**
    * @dev This function allows users to withdraw their stake after a 7 day waiting period from request 
    */
    function withdrawStake() external {
        zap.withdrawStake();
    }


    /**
    * @dev This function approves a _spender an _amount of tokens to use
    * @param _spender address
    * @param _amount amount the spender is being approved for
    * @return true if spender appproved successfully
    */
    function approve(address _spender, uint _amount) public returns (bool) {
        // return zap.approve(_spender,_amount);
        return token.approve(_spender, _amount);
    }


    /**
    * @dev Allows for a transfer of tokens to _to
    * @param _to The address to send tokens to
    * @param _amount The amount of tokens to send
    * @return true if transfer is successful
    */
    function transfer(address _to, uint256 _amount) public returns (bool) {
        // return zap.transfer(_to,_amount);
        return token.transfer(_to, _amount);
    }


    /**
    * @notice Send _amount tokens to _to from _from on the condition it
    * is approved by _from
    * @param _from The address holding the tokens being transferred
    * @param _to The address of the recipient
    * @param _amount The amount of tokens to be transferred
    * @return True if the transfer was successful
    */
    function transferFrom(address _from, address _to, uint256 _amount) public returns (bool) {
        // return zap.transferFrom(_from,_to,_amount);
        return token.transferFrom(_from, _to, _amount);
    }

}
