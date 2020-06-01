pragma solidity >=0.4.13;
pragma experimental ABIEncoderV2;

import "./ConvertLib.sol";

contract ProofDataPossession {

	struct filePledge{
		address  owner; //文件所有者
		uint32 amount; //为该文件的质押量
		string hash; //文件hash
		bool valid; //质押是否有效
	}
	struct challenge {
		address owner; //挑战者
		string file; //挑战文件
		bool valid; //挑战是否有效
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

	//节点退出, 从合约地址回退剩余token
	function NodeLeave() public {
		msg.sender.transfer(NodePledge[msg.sender]);
		NodePledge[msg.sender] = 0;
	}

	//为了保证存储服务的稳定性，存储节点需要完成一定量的质押, 质押地址为合约本身
	function PledgeForNode(uint32 amount, address node) public payable {
		require(msg.sender.balance>amount, "balance is not enough");
		NodePledge[node] = amount;
	}

	//存储方需要为了存储的文件，质押一定的ETH
	function PledgeForFile(uint32 amount, string memory file) public payable {
		require(msg.sender.balance>amount, "balance is not enough");
		FilePledge[file] = filePledge({owner: msg.sender, amount:amount, valid:true, hash:file});
	}

	//下载方可以发起文件存在性挑战
	function Challenge(string memory _file) public {
		ChallengeList[_file] = challenge({owner:msg.sender, file:_file, valid:true});
	}

	//存储节点获得挑战申请

	function GetChallengeList() public {
/*
		challenge[] memory cList;
				for(uint32 i=0;i<ChallengeList.keys.length; i++){
					if(ChallengeList[ChallengeList.keys[i]].valid == true){
						cList.push(ChallengeList[ChallengeList.keys[i]]);
					}
				}
*/
		//return cList;
	}

	function GetChallengeResponse(string memory _file) public view  returns(address){
		if(ProofList[_file].exist==0){
			return address(0);
		}
		return ProofList[_file].owner;
	}

	//存储节点提供数据持有性证明, 如果已经存在，则直接退出；
	function ProofProvide(string memory _file,  string memory _proof) public {
		if(ProofList[_file].exist==1){
			return;
		}
		ChallengeList[_file].valid = false;
		ProofList[_file] = proof({owner:msg.sender, file:_file, value:_proof, exist:1});
	}

	//当存储节点提供假数据时，会有惩罚措施; 举报node提供的file是错误的，合约会发起挑战验证；
	function Punish(address node, address file)public{

	}

	//合约销毁
	function destroy() public {
		require(msg.sender == owner);
		selfdestruct(msg.sender);
	}

}
