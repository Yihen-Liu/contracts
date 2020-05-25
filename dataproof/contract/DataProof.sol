pragma solidity >=0.4.25 <0.7.0;
pragma experimental ABIEncoderV2;

import "./ConvertLib.sol";

contract ProofDataPossession {

	struct filePledge{
		address  owner; //文件所有者
		uint32 amount; //为该文件的质押量
		bytes32 hash; //文件hash
		bool valid; //质押是否有效
	}
	struct challenge {
		address owner; //挑战者
		bytes32 file; //挑战文件
		bytes32 nonce; //随机值
		bool valid; //挑战是否有效
	}

	enum ProofStatus {Init, ToBeVerified, Success, Fail}
	struct proof{
		address owner;
		bytes32 file;
		bytes32 nonce;
		bytes32 value;
		ProofStatus status;
	}

	mapping(address => uint32) public NodePledge;
	mapping(bytes32 => filePledge) public FilePledge;
	mapping(string => challenge) public ChallengeList;
	mapping (string => proof[]) public ProofList;

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
	function PledgeForFile(uint32 amount, bytes32 file) public payable {
		require(msg.sender.balance>amount, "balance is not enough");
		FilePledge[file] = filePledge({owner: msg.sender, amount:amount, valid:true, hash:file});
	}

	//下载方可以发起文件存在性挑战
	function Challenge(bytes32 file, bytes32 _nonce) public {
		string memory index; // = strConcat(_nonce, file);
		ChallengeList[index] = challenge({owner:msg.sender, file:file, nonce:_nonce, valid:true});
	}

	//存储节点获得挑战申请
	//function GetChallengeList() public pure returns (challenge[] memory){
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

	//存储节点提供数据持有性证明
	function ProofProvide(bytes32 _file, bytes32 _nonce, bytes32 _proof) public {
/*
		string memory index = strConcat(_nonce, _file);
		mapping (string => proof[]) storage ref = ProofList;
		ref[index] = proof({owner:msg.sender, file:_file, nonce:_nonce, value:_proof, status:ProofStatus.ToBeVerified});
*/
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
