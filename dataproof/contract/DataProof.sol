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

	address public owner;
	uint32 ProfitChallengeSuccess = 1000;

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
	function Challenge(string memory _file) public {
		ChallengeList[_file] = challenge({owner:msg.sender, file:_file, valid:true});
	}

	//store will get challenge list round time.
	function GetChallengeList() public pure returns(address) {
/*
		challenge[] memory cList;
				for(uint32 i=0;i<ChallengeList.keys.length; i++){
					if(ChallengeList[ChallengeList.keys[i]].valid == true){
						cList.push(ChallengeList[ChallengeList.keys[i]]);
					}
				}
*/
		//return cList;
		return address(0);
	}

	function GetChallengeResponse(string memory _file) public view  returns(address){
		if(ProofList[_file].exist==0){
			return address(0);
		}
		return ProofList[_file].owner;
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
	function Punish(string memory node, string memory file)public{

	}

	//destroy the smart contract
	function destroy() public {
		require(msg.sender == owner);
		selfdestruct(msg.sender);
	}

}
