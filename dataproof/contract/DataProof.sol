pragma solidity >=0.4.13;
pragma experimental ABIEncoderV2;

import "./ConvertLib.sol";

contract ProofDataPossession {

	struct filePledge{
		address  owner; //file owner
		uint32 amount; //pledge eth amount for the file
		string hash; //file hash
		bool valid; //does pledge is valid
	}
	struct challenge {
		address owner; //who challenge start
		string file; //challenge file
		bool valid;
	}

	struct challIndex{
		string file;
		uint32 status; /*0:unverified; 1: verifiying; 2:done*/
	}

	enum ProofStatus {Init, ToBeVerified, Success, Fail}

	struct proof{
		address owner;
		string file;
		string value;
        uint32 exist;
	}

	mapping(address => uint32) public NodePledge;
	mapping(string => filePledge) public FilePledge;
	mapping(string => challenge) public ChallengeList;
	mapping (string => proof) public ProofList;
	challIndex[4096] challArray;
	address public owner;
	uint32 _now = 0;
	uint32 constant REWARD_AMOUNT_PER_FILE = 1000;
	uint32 constant PUNISH_AMOUNT_PER_FILE = 3000;
	constructor() public{
		owner = msg.sender;
	}

	//when node exit from the network, smart contract will give back left eth to the node.
	function NodeLeave() public {
		msg.sender.transfer(NodePledge[msg.sender]);
		NodePledge[msg.sender] = 0;
	}

    //in order to keep the store service stable, store node need to pledge some eth token.
	function PledgeForNode(uint32 amount, address node) public payable {
		require(msg.sender.balance>amount, "balance is not enough");
		NodePledge[node] = amount;
	}

	//when a node want to store file to network, need to pledge some eth for the file.
	function PledgeForFile(uint32 amount, string memory file) public payable {
		require(msg.sender.balance>amount, "balance is not enough");
		FilePledge[file] = filePledge({owner: msg.sender, amount:amount, valid:true, hash:file});
	}

	//downloader can start a file challenge to verify that someone has the entry.
	function Challenge(string memory _file) public returns(bool) {
		if (challArray.length>4096) {
			return false;
		}
		challArray[now]=challIndex({file:_file, status:0});
		ChallengeList[_file] = challenge({owner:msg.sender, file:_file, valid:true});
		return true;
	}

	//store will get challenge list round time.
	function GetChallengeList() public view returns(address) {
		for(uint32 i=0;i<challArray.length;i++){
			if(challArray[i].status==0){
				//return challArray[i].file;
			}
		}
		return address(0);
	}

	function GetChallengeResponse(string memory _file) public view  returns(address){
		if(ProofList[_file].exist==0){
			return address(0);
		}
		return ProofList[_file].owner;
	}


	function utilCompareInternal(string memory a, string memory b) internal pure returns (bool) {
		if (bytes(a).length != bytes(b).length) {
			return false;
		}
		for (uint i = 0; i < bytes(a).length; i ++) {
			if(bytes(a)[i] != bytes(b)[i]) {
				return false;
			}
		}
		return true;
	}

    //store node provide the data-provide-proof for the file; if this proof has been provided, exit directly.
	function ProofProvide(string memory _file,  string memory _proof) public {
		if(ProofList[_file].exist==1){
			return;
		}
		ChallengeList[_file].valid = false;
		ProofList[_file] = proof({owner:msg.sender, file:_file, value:_proof, exist:1});
	}

    //when storage node provide error data for downloader, punish will be done.
	function PunishOrReward(address payable node, string memory file, bool valid)public{
		if(valid==true){
			for(uint32 i=0;i<challArray.length;i++){
				if(utilCompareInternal(file, challArray[i].file)){
					node.transfer(REWARD_AMOUNT_PER_FILE);
					delete challArray[i];
					delete ChallengeList[file];
				}
			}
		}else{
			for(uint32 i=0;i<challArray.length;i++){
				if(utilCompareInternal(file, challArray[i].file)){
					//node.send(PUNISH_AMOUNT_PER_FILE,owner);
					challArray[i].status=0;
				}
			}
		}
	}

	//destroy the smart contract
	function destroy() public {
		require(msg.sender == owner);
		selfdestruct(msg.sender);
	}

}
