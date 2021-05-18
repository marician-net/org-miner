pragma solidity ^0.5.0;

import "./ZapStorage.sol";
import "./ZapTransfer.sol";

/**
* @title Zap Dispute
* @dev Contais the methods related to disputes. Zap.sol references this library for function's logic.
*/


library ZapDispute {
    using SafeMathM for uint256;

    event NewDispute(uint indexed _disputeId, uint indexed _requestId, uint _timestamp, address _miner);//emitted when a new dispute is initialized
    event Voted(uint indexed _disputeID, bool _position, address indexed _voter);//emitted when a new vote happens
    event DisputeVoteTallied(uint indexed _disputeID, int _result,address indexed _reportedMiner,address _reportingParty, bool _active);//emitted upon dispute tally
    event NewZapAddress(address _newZap); //emmited when a proposed fork is voted true

    /*Functions*/

    /**
    * @dev Allows token holders to vote
    * @param _disputeId is the dispute id
    * @param _supportsDispute is the vote (true=the dispute has basis false = vote against dispute)
    */
    function vote(ZapStorage.ZapStorageStruct storage self, uint _disputeId, bool _supportsDispute) public {
        ZapStorage.Dispute storage disp = self.disputesById[_disputeId];
        
        //Get the voteWeight or the balance of the user at the time/blockNumber the disupte began
        uint voteWeight = ZapTransfer.balanceOfAt(self,msg.sender,disp.disputeUintVars[keccak256("blockNumber")]);
        
        //Require that the msg.sender has not voted
        require(disp.voted[msg.sender] != true);
        
        //Requre that the user had a balance >0 at time/blockNumber the disupte began
        require(voteWeight > 0);
        
        //ensures miners that are under dispute cannot vote
        require(self.stakerDetails[msg.sender].currentStatus != 3);
        
        //Update user voting status to true
        disp.voted[msg.sender] = true;
        
        //Update the number of votes for the dispute
        disp.disputeUintVars[keccak256("numberOfVotes")] += 1;
        
        //Update the quorum by adding the voteWeight
        disp.disputeUintVars[keccak256("quorum")] += voteWeight; 
        
        //If the user supports the dispute increase the tally for the dispute by the voteWeight
        //otherwise decrease it
        if (_supportsDispute) {
            disp.tally = disp.tally + int(voteWeight);
        } else {
            disp.tally = disp.tally - int(voteWeight);
        }
        
        //Let the network know the user has voted on the dispute and their casted vote
        emit Voted(_disputeId,_supportsDispute,msg.sender);
    }


    /**
    * @dev tallies the votes.
    * @param _disputeId is the dispute id
    */
    function tallyVotes(ZapStorage.ZapStorageStruct storage self, uint _disputeId) public {
        ZapStorage.Dispute storage disp = self.disputesById[_disputeId];
        ZapStorage.Request storage _request = self.requestDetails[disp.disputeUintVars[keccak256("requestId")]];

        //Ensure this has not already been executed/tallied
        require(disp.executed == false);

        //Ensure the time for voting has elapsed
        require(now > disp.disputeUintVars[keccak256("minExecutionDate")]);  

        //If the vote is not a proposed fork 
        if (disp.isPropFork== false){
        ZapStorage.StakeInfo storage stakes = self.stakerDetails[disp.reportedMiner];  
            //If the vote for disputing a value is succesful(disp.tally >0) then unstake the reported 
            // miner and transfer the stakeAmount and dispute fee to the reporting party 
            if (disp.tally > 0 ) { 

                //Changing the currentStatus and startDate unstakes the reported miner and allows for the
                //transfer of the stakeAmount
                stakes.currentStatus = 0;
                stakes.startDate = now -(now % 86400);

                //Decreases the stakerCount since the miner's stake is being slashed
                self.uintVars[keccak256("stakerCount")]--;
                updateDisputeFee(self);

                //Transfers the StakeAmount from the reporded miner to the reporting party
                ZapTransfer.doTransfer(self, disp.reportedMiner,disp.reportingParty, self.uintVars[keccak256("stakeAmount")]);
                
                //Returns the dispute fee to the reportingParty
                ZapTransfer.doTransfer(self, address(this),disp.reportingParty,disp.disputeUintVars[keccak256("fee")]);
                
                //Set the dispute state to passed/true
                disp.disputeVotePassed = true;

                //If the dispute was succeful(miner found guilty) then update the timestamp value to zero
                //so that users don't use this datapoint
                if(_request.inDispute[disp.disputeUintVars[keccak256("timestamp")]] == true){
                    _request.finalValues[disp.disputeUintVars[keccak256("timestamp")]] = 0;
                }

            //If the vote for disputing a value is unsuccesful then update the miner status from being on 
            //dispute(currentStatus=3) to staked(currentStatus =1) and tranfer the dispute fee to the miner
            } else {
                //Update the miner's current status to staked(currentStatus = 1)
                stakes.currentStatus = 1;              
                //tranfer the dispute fee to the miner
                ZapTransfer.doTransfer(self,address(this),disp.reportedMiner,disp.disputeUintVars[keccak256("fee")]);
                if(_request.inDispute[disp.disputeUintVars[keccak256("timestamp")]] == true){
                    _request.inDispute[disp.disputeUintVars[keccak256("timestamp")]] = false;
                }
            }
        //If the vote is for a proposed fork require a 20% quorum before exceduting the update to the new zap contract address
        } else {
            if(disp.tally > 0 ){
                require(disp.disputeUintVars[keccak256("quorum")] >  (self.uintVars[keccak256("total_supply")] * 20 / 100));
                self.addressVars[keccak256("zapContract")] = disp.proposedForkAddress;
                disp.disputeVotePassed = true;
                emit NewZapAddress(disp.proposedForkAddress);
            }
        }
        
        //update the dispute status to executed
        disp.executed = true;
        emit DisputeVoteTallied(_disputeId,disp.tally,disp.reportedMiner,disp.reportingParty,disp.disputeVotePassed);
    }


    /**
    * @dev Allows for a fork to be proposed
    * @param _propNewZapAddress address for new proposed Zap
    */
    function proposeFork(ZapStorage.ZapStorageStruct storage self, address _propNewZapAddress) public {
        bytes32 _hash = keccak256(abi.encodePacked(_propNewZapAddress));
        require(self.disputeIdByDisputeHash[_hash] == 0);
        ZapTransfer.doTransfer(self, msg.sender,address(this), self.uintVars[keccak256("disputeFee")]);//This is the fork fee
        self.uintVars[keccak256("disputeCount")]++;
        uint disputeId = self.uintVars[keccak256("disputeCount")];
        self.disputeIdByDisputeHash[_hash] = disputeId;
        self.disputesById[disputeId] = ZapStorage.Dispute({
            hash: _hash,
            isPropFork: true,
            reportedMiner: msg.sender, 
            reportingParty: msg.sender, 
            proposedForkAddress: _propNewZapAddress,
            executed: false,
            disputeVotePassed: false,
            tally: 0
            }); 
        self.disputesById[disputeId].disputeUintVars[keccak256("blockNumber")] = block.number;
        self.disputesById[disputeId].disputeUintVars[keccak256("fee")]  = self.uintVars[keccak256("disputeFee")];
        self.disputesById[disputeId].disputeUintVars[keccak256("minExecutionDate")] = now + 7 days;
    }
    

    /**
    * @dev this function allows the dispute fee to fluctuate based on the number of miners on the system.
    * The floor for the fee is 15.
    */
    function updateDisputeFee(ZapStorage.ZapStorageStruct storage self) public {
            //if the number of staked miners divided by the target count of staked miners is less than 1
            if(self.uintVars[keccak256("stakerCount")]*1000/self.uintVars[keccak256("targetMiners")] < 1000){
                //Set the dispute fee at stakeAmt * (1- stakerCount/targetMiners)
                //or at the its minimum of 15 
                self.uintVars[keccak256("disputeFee")] = SafeMathM.max(15,self.uintVars[keccak256("stakeAmount")].mul(1000 - self.uintVars[keccak256("stakerCount")]*1000/self.uintVars[keccak256("targetMiners")])/1000);
            }
            else{
                //otherwise set the dispute fee at 15 (the floor/minimum fee allowed)
                self.uintVars[keccak256("disputeFee")] = 15;
            }
    }
}
