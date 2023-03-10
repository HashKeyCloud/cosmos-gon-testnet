package types

/*
// tx_responses:{height:403 txhash:"31310083D8DA2C43FB9EF0291D9F72F4768FAFFC5D0CC942A6C2C88F249AC566"
//          data:"12240A222F697269736D6F642E6E66742E4D7367497373756544656E6F6D526573706F6E7365"
//          raw_log:"[{\"msg_index\":0,\"events\":[{\"type\":\"issue_denom\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"testIndividual\"}

	raw_log:"[{\"msg_index\":0,\"events\":[{\"type\":\"issue_denom\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"issutnametest\"},{\"key\":\"denom_name\",\"value\":\"name\"},{\"key\":\"creator\",\"value\":\"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/irismod.nft.MsgIssueDenom\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld\"}]}]}]" logs:{events:{type:"issue_denom" attributes:{key:"denom_id" value:"issutnametest"} attributes:{key:"denom_name" value:"name"} attributes:{key:"creator" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld"}} events:{type:"message" attributes:{key:"action" value:"/irismod.nft.MsgIssueDenom"} attributes:{key:"module" value:"nft"} attributes:{key:"sender" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld"}}} gas_wanted:200000 gas_used:73773 tx:{[/cosmos.tx.v1beta1.Tx]:{body:{messages:{type_url:"/irismod.nft.MsgIssueDenom" value:"\n\rissutnametest\x12\x04name\"*iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyldJUhttps://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.pngZ\x96\x01{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}"}} auth_info:{signer_infos:{public_key:{type_url:"/cosmos.crypto.secp256k1.PubKey" value:"\n!\x03\x0b\xab\x18\xf2\x0c\x17\x90\xb7e\xc7\xdf\xe7\xe3\xfe\x86\x897\xf6\xe6?\xef;\xb5T\x8d|kt>J\xba\x15"} mode_info:{single:{mode:SIGN_MODE_DIRECT}} sequence:58} fee:{amount:{denom:"uiris" amount:"20000"} gas_limit:200000}} signatures:"\xe3ΐ\xc8؁|0\xd5\x02ޑ\xa0H\x80>\xa2\x9e\x8c\x97\xc8\x02`\x02F\xc1\x81\x0f\xd7\x12\xa7+ \xe6\xba\xd8\xee\x85d@3+߈\xb0\xbb\x8a\xbe\x9bL\xfau\xf2\x8c\xa3\xe0\xe0\x0b?\x8a\x1046\x86"}} timestamp:"2023-03-08T08:38:47Z" events:{type:"coin_spent" attributes:{key:"spender" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true} attributes:{key:"amount" value:"20000uiris" index:true}} events:{type:"coin_received" attributes:{key:"receiver" value:"iaa17xpfvakm2amg962yls6f84z3kell8c5l9mr3fv" index:true} attributes:{key:"amount" value:"20000uiris" index:true}} events:{type:"transfer" attributes:{key:"recipient" value:"iaa17xpfvakm2amg962yls6f84z3kell8c5l9mr3fv" index:true} attributes:{key:"sender" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true} attributes:{key:"amount" value:"20000uiris" index:true}} events:{type:"message" attributes:{key:"sender" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true}} events:{type:"tx" attributes:{key:"fee" value:"20000uiris" index:true} attributes:{key:"fee_payer" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true}} events:{type:"tx" attributes:{key:"acc_seq" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld/58" index:true}} events:{type:"tx" attributes:{key:"signature" value:"486QyNiBfDDVAt6RoEiAPqKejJfIAmACRsGBD9cSpysg5rrY7oVkQDMr34iwu4q+m0z6dfKMo+DgCz+KEDQ2hg==" index:true}} events:{type:"message" attributes:{key:"action" value:"/irismod.nft.MsgIssueDenom" index:true}} events:{type:"issue_denom" attributes:{key:"denom_id" value:"issutnametest" index:true} attributes:{key:"denom_name" value:"name" index:true} attributes:{key:"creator" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true}} events:{type:"message" attributes:{key:"module" value:"nft" index:true} attributes:{key:"sender" value:"iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" index:true}}

	BlockHeight: tx.Height,
			TxHash:      tx.Txhash,
			ClassID:     denomId,
			Name:   denomName,
			Sender: creator,
*/
type IssueNFT struct {
	BlockHeight int64
	TxHash      string
	ClassID     string
	Name        string
	Sender      string
}

type MintNFT struct {
	BlockHeight int64
	TxHash      string
	ClassID     string
	NFTName     string
	Recipient   string
	Owner       string
	TokenUri    string
}

type TransferNFT struct {
	BlockHeight int64
	TxHash      string
	ClassID     string
	BaseClassID string
	NFTName     []string
	Sender      string
	Receiver    string
}

type PacketData struct {
	ClassData string   `json:"classData"`
	ClassID   string   `json:"classId"`
	ClassURI  string   `json:"classUri"`
	Receiver  string   `json:"receiver"`
	Sender    string   `json:"sender"`
	TokenData []string `json:"tokenData"`
	TokenIds  []string `json:"tokenIds"`
	TokenUris []string `json:"tokenUris"`
}
