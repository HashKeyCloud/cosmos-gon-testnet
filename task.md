# cosmos-gon-testnet
## 安装客户端
```shell
https://github.com/game-of-nfts/gon-testnets/blob/main/doc/installation.md#iris
```

- 配置远程客户端信息
```shell
vim /data/.iris/config/client.toml
```
- 远程客户端信息
```shell
https://github.com/game-of-nfts/gon-testnets/blob/main/doc/testnet-info.md
```

## 领取水
- https://faucet.ping.pub/
## Create one Collection on IRISnet
- 原始命令：
```shell
iris tx nft issue <denom-id> \
  --name <name> \
  --uri <uri> \
  --data <json-string-denom-data> \
  --schema <denom-schema> \
  --symbol <denom-symbol> \
  --mint-restricted <true|false> \
  --update-restricted <true|false>
 ```
- 执行命令：
```shell
ubuntu@10-7-95-161:~$ iris tx nft issue hashquarklogo --name="the hashquark logo" --from=hashquark --node="tcp://34.80.93.133:26657" --mint-restricted=false --update-restricted=false --uri="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgIssueDenom
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    description: ""
    id: hashquarklogo
    mint_restricted: false
    name: the hashquark logo
    schema: ""
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    symbol: ""
    update_restricted: false
    uri: https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: CB3BD702DFF7FA8B819AB6FC6E277F3766908E7CB68C31FFC5542012FDFC6760
```

## mint(至少需要mint两个nft)
- 原始命令：
```shell
iris tx nft mint <denom-id> <nft-id> \
--name <nft-name> \
--uri <nft-uri> \
--uri-hash <nft-uri-hash> \
--data <json-string-token-data> \
--recipient <nft-recipient>
```
- 执行命令：
```shell
ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo thehashquarklog --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: thehashquarklog
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 10AAAEDE795C8866A46B32E1AAD360BA058B6641E6A3D0DC092525F6E232FC54



ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo thehashquarklog1 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: thehashquarklog1
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: A255C7D6DD1FD112E70235C0D4AC2C43C0C30D92618198A4E5B4656F8E293DD5
```


## NFT发送
- 原始命令: classID:就是denom-id
```shell
simd tx nft-transfer transfer \
<src-chain-port> \
<src-chain-channel> \
<dst-chain-receiver> \
<classID> \
<nftID>  
```
- 执行命令
```shell
ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-24" "juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q" "hashquarklogo" "thehashquarklog"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "348292"
      revision_number: "6"
    timeout_timestamp: "1677751245179287129"
    token_ids:
    - thehashquarklog
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 336B9AE765EA80F88771BD3DBEFD059195615D361A59CE39BAE7432F5DE577FF


ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-0" "omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z" "hashquarklogo" "thehashquarklog1"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-0
    source_port: nft-transfer
    timeout_height:
      revision_height: "354003"
      revision_number: "1"
    timeout_timestamp: "1677751352487519991"
    token_ids:
    - thehashquarklog1
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 75BB202B053C053ED7A4696DA880ACE3DCCB032076A1B81168FCDF4473FA5F7F
```

## 从Juno转回IRIS
- 原始命令
```shell
junod tx wasm execute <cw-721-addr> \
{"send_nft": {"contract": "<ics-721-addr>", "token_id": "<token-id>", "msg": \ "<basa64_encoded_msg>"}}
```
- 其中`cw-721-addr`通过以下命令查询
```shell
simd q wasm contract-state smart <ics-721-addr> '{"nft_contract": {"class_id" : "<class-prefix/class-base-id>"}}'
```
备注
<ics-721-addr>: https://github.com/game-of-nfts/gon-testnets/blob/main/doc/port-channel-table.md
class_id:targetIcs721IbcPort + "/" + targetChannelId + "/" + sourceCw721Address
<base64_encoded_msg>: 
{
    "receiver": "<reciver-addr>",
    "channel_id": "<src-channel-id>",
    "timeout": {
        "block": {
            "revision": <chain-revision>,
            "height": <timeout-height>
        }
    }
}
- base64转化：https://www.base64encode.org/

- 执行命令

```shell
ubuntu@10-7-95-161:~$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9



ubuntu@10-7-95-161:~$ junod tx wasm execute juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "thehashquarklog", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDM0NzUxMAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"thehashquarklog","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDM0NzUxMAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F2E758646B4640215C252E4FDA7EE50084B62D79F16EC1B8D08C5E2B45611E51
```

验证：
```shell
iris query nft owner <address> --denom-id=<denom-id>
ubuntu@10-7-95-161:~$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld --denom-id=hashquarklogo --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
pagination:
  next_key: null
  total: "0"
```


## 从omniflixhub转回IRIS
- 执行命令

```shell
ubuntu@10-7-95-161:~$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-24/hashquarklogo --node=http://65.21.93.56:26657
hash: A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE

ubuntu@10-7-95-161:~$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" "iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" "ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE" "thehashquarklog1"  --from=hashquark --node=http://65.21.93.56:26657 --fees 2000uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "2000"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "265598"
      revision_number: "1"
    timeout_timestamp: "1677752766463877313"
    token_ids:
    - thehashquarklog1
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 1B5A97DD9B76007974950D4ABB51A46A760D91A045CC0531C6F20EC4DBA3C6AB
```



## A7~A20
### 背景知识
> Explanation 
> 
> Flow
> 
> i --(1)--> s means that you transfer an NFT from IRISnet(i) to Stargaze(s) through channel is-1(si-1). Check the channel table to find what is-1 represents. Id Each different flow has a unique id. In task we use flow-id to refer to a flow we want participants to follow and transfer NFT.
>
> Style
> 
>Never-go-back: No class-id prefix removal during the transfer flow (i.e. i --(1)--> s --(2)--> i ). This means when transferring the NFT back to iris from stargaze through is-2, it will result in a new Collection created on iris, even though the NFT is back on iris.
>
> Revisit: At least one class-id prefix removal during the transfer flow (i.e. i --(1)--> s --(1)--> j --(1)--> s --(2)--> i ). For example, when transferring the NFT back to stargaze from juno through js-1, there’s no new Collection created on stargaze. Hence a removal happened.
> 
> Backtrack: All class-id prefixes are finally removed in the transfer flow (i.e. i --(1)--> s --(1)--> i ). The original NFT on iris is back and unlocked.
### 所有任务
```shell
Flow Id	Style	Flow
a1	never-go-back	i --(1)--> s --(1)--> j --(1)--> i

a2	never-go-back	i --(1)--> u --(1)--> o --(1)--> i

a3	never-go-back	i --(1)--> s --(1)--> j --(1)--> u --(1)--> i

a4	never-go-back	i --(1)--> s --(1)--> o --(1)--> j --(1)--> i

a5	never-go-back	i --(1)--> s --(1)--> j --(1)--> u --(1)--> o --(1)--> s --(1)--> i

a6	never-go-back	i --(1)--> o --(1)--> s --(1)--> u --(1)--> o --(1)--> j --(1)--> i

b1	revisit	i --(1)--> s --(1)--> u --(1)--> s --(2)--> i

b2	revisit	i --(1)--> u --(1)--> o --(1)--> u --(2)--> i

b3	revisit	i --(1)--> j --(1)--> u --(1)--> j --(2)--> i

b4	revisit	i --(1)--> j --(1)--> s --(1)--> j --(2)--> i

c1	backtrack	i --(1)--> s --(1)--> j --(1)--> s --(1)--> i

c2	backtrack	i --(1)--> o --(1)--> u --(1)--> o --(1)--> i

c3	backtrack	i --(1)--> s --(1)--> j --(1)--> u --(1)--> j --(1)--> s --(1)--> i

c4	backtrack	i --(1)--> u --(1)--> s --(1)--> o --(1)--> s --(1)--> u --(1)--> i
```



### A7：never-go-back	i --(1)--> s --(1)--> j --(1)--> i
- 即保留所有的交易路径
- 执行命令
- i --(1)--> s

```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task7 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task7
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5D734DEA457CECE191A371C133A591AFF6255F035AA68CA837F6CF74337CDE44



ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" "stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd" "hashquarklogo" "task7"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3731304"
      revision_number: "1"
    timeout_timestamp: "1678099153827715758"
    token_ids:
    - task7
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 080E2364D6D7316CFF12DDA0B4F2DC78D26E8667D4A69AF95CDB6771F7CA669A
```

- s --(1)--> j
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d


ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task7", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogMzczMTcwMAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task7","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogMzczMTcwMAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: D703E63111AFCBC06EB8715C89C6640C2448E11F0A4375C5DB0A45214F93D592
```
- j --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task7", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task7","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: D417662340275371800A9FDA9E7B8FE9205F9B286E8F1758588E3CD4E0E62AD3
```

- 验证:地址下有task7
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "
```

### A8 i --(1)--> u --(1)--> o --(1)--> i
- i --(1)--> u
```shell
ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo task8 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task8
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 65CF9E7910F3392D403E95BB244F3F2D5A88F347EB128F7622B0196B21788FA4

ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-17" "uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3" "hashquarklogo" "task8"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-17
    source_port: nft-transfer
    timeout_height:
      revision_height: "2372549"
      revision_number: "2"
    timeout_timestamp: "1678159570574719329"
    token_ids:
    - task8
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 83B65168FAF5E92569CDD36BBF5FA68C886F59A096E346F1E5AD1898CB985FFC
```
- u --(1)--> o
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-3/hashquarklogo
hash: 855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F

uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F" "task8"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-5
    source_port: nft-transfer
    timeout_height:
      revision_height: "429748"
      revision_number: "1"
    timeout_timestamp: "1678161278646362748"
    token_ids:
    - task8
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 4196F7DC7BFA378F0FB3466569F34F00D1562C4A1A5753618AF5F27E3245A86B
```
- o --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-3/hashquarklogo
hash: B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" $iAddr "ibc/B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6" "task8"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "342393"
      revision_number: "1"
    timeout_timestamp: "1678161575562505625"
    token_ids:
    - task8
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: E1E745451777E913F3AB243E33349789AABA9F69051CE811EF56E548F2DCD60C

```
- 验证:地址下有task8
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "
```

### A9 i --(1)--> s --(1)--> j --(1)--> u --(1)--> i
- i --(1)--> s
```shell
ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo task9 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task9
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 6D4E0689149DC7A6C6523AB99817509C54D2E1678526EC1FDBAA8C4974D9762B

ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "task9"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3743415"
      revision_number: "1"
    timeout_timestamp: "1678168555468758130"
    token_ids:
    - task9
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: DD39209F2C0D5CF08E2ECF6DEF6CAA4D7097A3EB2A770F37E2202BC64486374D
```
- s --(1)--> j 
```shell
ubuntu@10-7-95-161:~$  starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task9", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task9","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F75F73362443CE3B51A7048F7B3217C0D209246B0F3506A1AB74E26DB827992E
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task9", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task9","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: EDAB4814E67C5DE731F3D66BCBA8EC85369768668B6CBDCA2162B1D1EC37E123
```
- u --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-3" $iAddr "ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613" "task9"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-3
    source_port: nft-transfer
    timeout_height:
      revision_height: "343892"
      revision_number: "1"
    timeout_timestamp: "1678169670647298322"
    token_ids:
    - task9
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5B5AD131F1891F552824DD6666F7FB74C985BAF0CBD8C41FE3B5F392AF69B62E
```
- 检查：iris地址下面有task9
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"
```

### A10：never-go-back	i --(1)--> s --(1)--> o --(1)--> j --(1)--> i
- i --(1)--> s
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo new1task10 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: new1task10
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: A37B82CA1F104F4D6F13A88B1F1422F965A5A7CE7F3A4288741288AA295CB6DD


ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "new1task10"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3746214"
      revision_number: "1"
    timeout_timestamp: "1678180290867759503"
    token_ids:
    - new1task10
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 366D41AAEBE0E84270D2E3CEDACC9E816B3421235B50278A8544433E3E437BD5
```
- s --(1)--> o
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "new1task10", "msg": "ewogICAgInJlY2VpdmVyIjogIm9tbmlmbGl4MTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtbDlsdjJ6IiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtMjA5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"new1task10","msg":"ewogICAgInJlY2VpdmVyIjogIm9tbmlmbGl4MTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtbDlsdjJ6IiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtMjA5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: E647A4E686EA4535F2ECCF48B93A644500A25C18852A2C2B8CEB328DA1087D94
```
- o --(1)--> j
```shell
ubuntu@10-7-95-161:~$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-44/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 292C9020747C9DC1AB878EB8030A9E87B654BDCFC2651A44A226304BAB46FF3E

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-46" $jAddr "ibc/292C9020747C9DC1AB878EB8030A9E87B654BDCFC2651A44A226304BAB46FF3E" "new1task10"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/292C9020747C9DC1AB878EB8030A9E87B654BDCFC2651A44A226304BAB46FF3E
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-46
    source_port: nft-transfer
    timeout_height:
      revision_height: "422587"
      revision_number: "6"
    timeout_timestamp: "1678180549940937846"
    token_ids:
    - new1task10
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 43C5DD0F09290350C78854AA91607698E27183716F743B2100800C0D91242A88
```
- j --(1)--> i
```shell

ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-91/nft-transfer/channel-44/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1cqytzwmjy8kqfktq25u632ku6qanq52zcs3caxf9wttwxvz8w2ts3xn0f2


ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1cqytzwmjy8kqfktq25u632ku6qanq52zcs3caxf9wttwxvz8w2ts3xn0f2 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "new1task10", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1a5zcyflc9lc0qllswmsdmhdfp62mq94u2c0dzkwrdszjvw88jjxqcwvyh9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"newtask10","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 1E0284004AACAB9F97E6160C667F9379F3F5FB4AA477F84FC565C4452EB4A78B
```
- 检查：iris地址下面有new1task10
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/660086CBB68FC1BD57F3839865E9FE525889021B9570088C3A80437A51D442AB
    token_ids:
    - newtask10
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"
```


### A11 i --(1)--> s --(1)--> j --(1)--> u --(1)--> o --(1)--> s --(1)--> i
- i --(1)--> s
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo newtask11 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: newtask11
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 01CC72B74D778E19C7BDD3981ECE2CC9E34474AC495841C78A9CCC94933212A3


ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "newtask11"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3745954"
      revision_number: "1"
    timeout_timestamp: "1678178820004845214"
    token_ids:
    - newtask11
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: D12B59822C7972DF8C262FE11B4346CDBABA6456B893F1306D501475EF6B96AA
```
- s --(1)--> j
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "newtask11", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"newtask11","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 952DE40738162D9E2B88DABE883841C69B034E6D0FF2181AA68A280DA94B1CBB
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "newtask11", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"newtask11","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 8F14A6D1414E7ECA33E45F2EF28468E5D15E703690FD8FE39C634653F3A3F426
```
- u --(1)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613" "newtask11"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-5
    source_port: nft-transfer
    timeout_height:
      revision_height: "432962"
      revision_number: "1"
    timeout_timestamp: "1678179274882902196"
    token_ids:
    - newtask11
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: ABCFCDDA79D353535C2CA0CEFCD36175F47BE7248E3557D45946D77BD87045CE

```
- o --(1)--> s
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 95C767F5357A83C97BAA43387CB297ECE2F47D04D2ABC7FF2DF493BE988CC22F

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-44" $sAddr "ibc/95C767F5357A83C97BAA43387CB297ECE2F47D04D2ABC7FF2DF493BE988CC22F" "newtask11"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/95C767F5357A83C97BAA43387CB297ECE2F47D04D2ABC7FF2DF493BE988CC22F
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-44
    source_port: nft-transfer
    timeout_height:
      revision_height: "3745986"
      revision_number: "1"
    timeout_timestamp: "1678179323235025586"
    token_ids:
    - newtask11
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F8887ED6D7978678871AFC55EA565C72C2280A6F4AFD04E93F4CB0CCDA83D87B
```
- s --(1)--> i
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-209/nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1kky8q9as0zw04xjtj7sgljnvuu3yhjp7c52ek9qe07jmegtgzw3qv4etld

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1kky8q9as0zw04xjtj7sgljnvuu3yhjp7c52ek9qe07jmegtgzw3qv4etld '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "newtask11", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NDgzOAogICAgICAgIH0KICAgIH0KfQo="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1kky8q9as0zw04xjtj7sgljnvuu3yhjp7c52ek9qe07jmegtgzw3qv4etld","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"newtask11","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NDgzOAogICAgICAgIH0KICAgIH0KfQo="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5BA9B47334E416181449622E905413F9B0CC52C965E5602F8E44EE19FA4088B1
```

- 检查：iris地址下面有newtask11
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/660086CBB68FC1BD57F3839865E9FE525889021B9570088C3A80437A51D442AB
    token_ids:
    - newtask10
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"
```

### A12 i --(1)--> o --(1)--> s --(1)--> u --(1)--> o --(1)--> j --(1)--> i
- i --(1)--> o 
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task12 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task12
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F2E2EF3B05488F60E3D224F0894162A25A4AA1F5575BAACBD3E6C33164985043

ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-0" $oAddr "hashquarklogo" "task12"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
password must be at least 8 characters
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-0
    source_port: nft-transfer
    timeout_height:
      revision_height: "432989"
      revision_number: "1"
    timeout_timestamp: "1678179924737692826"
    token_ids:
    - task12
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 681D0BDD9E64573A2CDDFA27C29202A6A6B7E90C1C704D4F907D2071B12ECBD7
```
- o --(1)--> s 
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-24/hashquarklogo
hash: A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-44" $sAddr "ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE" "task12"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-44
    source_port: nft-transfer
    timeout_height:
      revision_height: "3746460"
      revision_number: "1"
    timeout_timestamp: "1678182194450904734"
    token_ids:
    - task12
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 0673E51883DF35E0431235FF7594D7032F774F2AEC47561FCCF8D06C2CC73493
```
- s --(1)--> u 
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-209/nft-transfer/channel-24/hashquarklogo"}}'
data: stars1fuq0we8wz2kfsqgrs9agpm0kfcw05w0nst08j7rxec69e68mgnjsl384fr

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1fuq0we8wz2kfsqgrs9agpm0kfcw05w0nst08j7rxec69e68mgnjsl384fr '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task12", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1fuq0we8wz2kfsqgrs9agpm0kfcw05w0nst08j7rxec69e68mgnjsl384fr","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task12","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 38942234B76777F65F4C1E8AEA71D3EC8066CAC3F00C29F6FE4299C51C375618
```
- u --(1)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-6/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-209/nft-transfer/channel-24/hashquarklogo
hash: C22573F2D994DE81D5F0823D9C1DD8B5A20562E5078A8554F29409500CF6AF6D

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/C22573F2D994DE81D5F0823D9C1DD8B5A20562E5078A8554F29409500CF6AF6D" "task12"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/C22573F2D994DE81D5F0823D9C1DD8B5A20562E5078A8554F29409500CF6AF6D
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-5
    source_port: nft-transfer
    timeout_height:
      revision_height: "433769"
      revision_number: "1"
    timeout_timestamp: "1678183029673965154"
    token_ids:
    - task12
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 9B18B5021BFFE7C682B1335B3C7117547243572C707E4622A6DFF2FF85DD524A
```
- o --(1)--> j
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-6/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-209/nft-transfer/channel-24/hashquarklogo
hash: A2E1F234F062672170A84FEBCCE6CEB6C106E9296FA639FFC93C9665FD20C819

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-46" $jAddr "ibc/A2E1F234F062672170A84FEBCCE6CEB6C106E9296FA639FFC93C9665FD20C819" "task12"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/A2E1F234F062672170A84FEBCCE6CEB6C106E9296FA639FFC93C9665FD20C819
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-46
    source_port: nft-transfer
    timeout_height:
      revision_height: "423176"
      revision_number: "6"
    timeout_timestamp: "1678183127032927110"
    token_ids:
    - task12
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 45C7B2E913CB8CDE7A1471C2DBAC5908A7D80BCF21122CF5D40B7A2824784A6F
```
- j --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-91/nft-transfer/channel-41/nft-transfer/channel-6/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-209/nft-transfer/channel-24/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1k397vsrzj9gna62s024kcm2n7tkz6v8nxxjhqf2g8zx43yupkcfsxq5r4l

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1k397vsrzj9gna62s024kcm2n7tkz6v8nxxjhqf2g8zx43yupkcfsxq5r4l '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task12", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1k397vsrzj9gna62s024kcm2n7tkz6v8nxxjhqf2g8zx43yupkcfsxq5r4l","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task12","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDIzNjkwCiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 983389BD4895F01FD1A50267AA7796288765377168C4141FE016429C4464FC93
```
- 检查：iris地址下面有task12
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/660086CBB68FC1BD57F3839865E9FE525889021B9570088C3A80437A51D442AB
    token_ids:
    - newtask10
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"
```


### A13 i --(1)--> s --(1)--> u --(1)--> s --(2)--> i
- i --(1)--> s 
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task13 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task13
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 61F95D77A72651286C2229A39173EB472F5403419F41F862244011C4D934637A


ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "task13"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3746753"
      revision_number: "1"
    timeout_timestamp: "1678183570550186594"
    token_ids:
    - task13
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: A80B5FB1229307D147473CED5E8854696792AD838E4901214F0FA7BC0C0F1598
```
- s --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task13", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task13","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: E0BD1CB933DB3A0AA2D85C0C7FC1F1DE3C4744B8EAA3C58DFC44688C01576528
```
- u --(1)--> s
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-6/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 1347D2FC2CF401CF7D73E1F3A50AAF762312169EE3459E4FEFF6FBFA1585D769

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-6" $sAddr "ibc/1347D2FC2CF401CF7D73E1F3A50AAF762312169EE3459E4FEFF6FBFA1585D769" "task13"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/1347D2FC2CF401CF7D73E1F3A50AAF762312169EE3459E4FEFF6FBFA1585D769
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-6
    source_port: nft-transfer
    timeout_height:
      revision_height: "3746863"
      revision_number: "1"
    timeout_timestamp: "1678184036824717564"
    token_ids:
    - task13
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 4F955789C7FD114B0B09FD43DDA7D87AC1C0ED90B4867936E2E61EFDC89731DF
```
- s --(2)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task13", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwOCIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task13","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwOCIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 05FB23338C0C438E50F6A2E28D307B0E133E6EB7EA43871AEB7B6213A7FB9A21
```

- 检查：iris地址下面有task13
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/03C876D8704E6202BD830FB24B99D9DE40ABBBDBAB95A8CD5DA668ACD3076739
    token_ids:
    - new1task10
  - denom_id: ibc/660086CBB68FC1BD57F3839865E9FE525889021B9570088C3A80437A51D442AB
    token_ids:
    - newtask10
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/94F60604303084FB1C63A7D6C6B2BC0C654C224BAA64F78A6EC5596874ADF6F4
    token_ids:
    - task13
  - denom_id: ibc/D263639D255AF869577FA40B74E759B510BBB209DC86D5B3A245186A37925076
    token_ids:
    - newtask11
  - denom_id: ibc/D42D85FAA3F6FC33B516900B36B63F85277AA47A89416DFBB1BBDEFB2519F776
    token_ids:
    - task12
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: ibc/FBA04C7966EA860855D8E7EA619E2C375B288280DF8DE5F93969DE377FFEEF26
    token_ids:
    - task11
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"

```

### A14 i --(1)--> u --(1)--> o --(1)--> u --(2)--> i
- i --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task14 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task14
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: C80BC875F49154B38985D91A9286C2000105D19FE58BA7580FD309218BE470C0



ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-17" "uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3" "hashquarklogo" "task14"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-17
    source_port: nft-transfer
    timeout_height:
      revision_height: "2377312"
      revision_number: "2"
    timeout_timestamp: "1678184579532829462"
    token_ids:
    - task14
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: D5B26C9B9EE95F5E881E866C6AFB4149A5CD87ED953ACD8BC65B1B5FE1FA5555
```
- u --(1)--> o 
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-3/hashquarklogo
hash: 855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F" "task14"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-5
    source_port: nft-transfer
    timeout_height:
      revision_height: "434076"
      revision_number: "1"
    timeout_timestamp: "1678184792634858587"
    token_ids:
    - task14
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 4E7F6CA94651A4B213841F9D7DA28BAB72B72095F9A798B07ADDFE211B43022C
```
- o --(1)--> u 
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-3/hashquarklogo
hash: B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-41" $uAddr "ibc/B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6" "task14"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/B2FAEB9A35C41EDC7C7827190478781B1F994E3CFA11D97246A5ED079176A9B6
    memo: ""
    receiver: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-41
    source_port: nft-transfer
    timeout_height:
      revision_height: "2377490"
      revision_number: "2"
    timeout_timestamp: "1678185018074011130"
    token_ids:
    - task14
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 03DF65BF2069BE55AF9F634D88A644EEE4C2195DF1935B1AECF681454CA3C2BD

```
- u --(2)--> i
```
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-3/hashquarklogo
hash: 855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-4" $iAddr "ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F" "task14"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-4
    source_port: nft-transfer
    timeout_height:
      revision_height: "346588"
      revision_number: "1"
    timeout_timestamp: "1678185520736813333"
    token_ids:
    - task14
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 1D05549CE829383DBFE0257B70DCB53B0711BA8D952D3301F305139F4FDCB8C2
```

- 检查：iris地址下面有task14
```shell
ubuntu@10-7-95-161:~/go/bin$ iris q nft owner iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
owner:
  address: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
  id_collections:
  - denom_id: hashquarklogo
    token_ids:
    - thehashquarklog
    - thehashquarklog1
  - denom_id: ibc/03C876D8704E6202BD830FB24B99D9DE40ABBBDBAB95A8CD5DA668ACD3076739
    token_ids:
    - new1task10
  - denom_id: ibc/39D166AA5D97C236FB6817B88D423137A3B1CFC9ED9267C2BE978701E5B9A35E
    token_ids:
    - task14
  - denom_id: ibc/660086CBB68FC1BD57F3839865E9FE525889021B9570088C3A80437A51D442AB
    token_ids:
    - newtask10
  - denom_id: ibc/7B0C278C76A20F82CA4073F56EA86A698841F5F737320564EC11DA7FD6CB9AE0
    token_ids:
    - task9
  - denom_id: ibc/93AD90576BAFA7FC0D00DA77BD3EA1C8426CD52A3E4DCDF2B576C2BA0BB9902A
    token_ids:
    - task7
  - denom_id: ibc/94F60604303084FB1C63A7D6C6B2BC0C654C224BAA64F78A6EC5596874ADF6F4
    token_ids:
    - task13
  - denom_id: ibc/D263639D255AF869577FA40B74E759B510BBB209DC86D5B3A245186A37925076
    token_ids:
    - newtask11
  - denom_id: ibc/D42D85FAA3F6FC33B516900B36B63F85277AA47A89416DFBB1BBDEFB2519F776
    token_ids:
    - task12
  - denom_id: ibc/DF39BAAAAC2A67173F1DFDCBB2F34E74B60FED51273D6EC9BBB0E2EB0865E8D7
    token_ids:
    - task8
  - denom_id: ibc/FBA04C7966EA860855D8E7EA619E2C375B288280DF8DE5F93969DE377FFEEF26
    token_ids:
    - task11
  - denom_id: logo
    token_ids:
    - hashquark1
    - hashquark2
    - hashquark5
pagination:
  next_key: null
  total: "0"
```

### A15 i --(1)--> j --(1)--> u --(1)--> j --(2)--> i
- i --(1)--> j
```shell
ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo task15 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task15
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: FFBADA7DA9490CE6366EF05BB3A06BCB1B05CF360EDDE2B766CA68A2495B2BCE

ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-24" $jAddr "hashquarklogo" "task15"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "424415"
      revision_number: "6"
    timeout_timestamp: "1678190335684827463"
    token_ids:
    - task15
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 3022021B9866662DBB09F8FCCE7FC60722FEB36770762CB13415F010055259AD
```
- j --(1)--> u 
```shell
ubuntu@10-7-95-161:~$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9

ubuntu@10-7-95-161:~$ junod tx wasm execute juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task15", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task15","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: DFA12BF3272CDBEC19EF82D70ED173094F298C7F7EFA34615E9847AACF60A597
```
- u --(1)--> j 
```shell
ubuntu@10-7-95-161:~$ uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo
hash: AF50F2C5AEE12AB30C648324817C04A04E22C4065DD724E61D6AB8F7A5A400C3

ubuntu@10-7-95-161:~$ uptickd tx nft-transfer transfer "nft-transfer" "channel-7" $jAddr "ibc/AF50F2C5AEE12AB30C648324817C04A04E22C4065DD724E61D6AB8F7A5A400C3" "task15"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/AF50F2C5AEE12AB30C648324817C04A04E22C4065DD724E61D6AB8F7A5A400C3
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-7
    source_port: nft-transfer
    timeout_height:
      revision_height: "424503"
      revision_number: "6"
    timeout_timestamp: "1678190843884094124"
    token_ids:
    - task15
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 0E9DDB304F8E6A6BF5A428AF4245AC6F716A75C5A2CC6D821D916DEBE861DABF
```
- j --(2)--> i
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9

ubuntu@10-7-95-161:~$ junod tx wasm execute juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task15", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTkwIiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task15","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTkwIiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: A54777A7DEB0598C5F2DBA397B10204334DCE15F1ED04C4BA542022D0DCFCD70

```

### A16 i --(1)--> j --(1)--> s --(1)--> j --(2)--> i
- i --(1)--> j 
```shell
ubuntu@10-7-95-161:~$ iris tx nft mint hashquarklogo task16 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task16
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 75DB8780482B90E2E87C4C1A450DF0890EF310D99D1DE66B82353BA8323F4C4C

ubuntu@10-7-95-161:~$ iris tx nft-transfer transfer "nft-transfer" "channel-24" $jAddr "hashquarklogo" "task16"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "424444"
      revision_number: "6"
    timeout_timestamp: "1678191635384492663"
    token_ids:
    - task16
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: B6530B009EA6915ABF35A3EC7D726626B582B027BC47B9431578464297D99010
```
- j --(1)--> s 
```shell
ubuntu@10-7-95-161:~$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9

ubuntu@10-7-95-161:~$ junod tx wasm execute juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task16", "msg": "ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDU2MzQKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task16","msg":"ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDU2MzQKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 742935F6D7757EAA41F2427B2647FB8325475F17A972841FA17DE235F7C1C303
```
- s --(1)--> j 
```shell
ubuntu@10-7-95-161:~$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-211/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}'
data: stars1t3ecex7hxlfcnp6tn4n5ugkcgdher6m73uqhulkdpn3xwws27ehqvapc6k

ubuntu@10-7-95-161:~$ starsd tx wasm execute stars1t3ecex7hxlfcnp6tn4n5ugkcgdher6m73uqhulkdpn3xwws27ehqvapc6k '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task16", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDU2MzQKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1t3ecex7hxlfcnp6tn4n5ugkcgdher6m73uqhulkdpn3xwws27ehqvapc6k","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task16","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDU2MzQKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: F2D7A339E20E8D5EC5ACD54CB83707178F1D66EB0B912AAD6C47CE956942BA3A
```
- j --(2)--> i
```shell
ubuntu@10-7-95-161:~$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-89/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9

ubuntu@10-7-95-161:~$ junod tx wasm execute juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task16", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTkwIiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1jqu23ee4zjds8mnc6cnc9de9lv38vg8juqchug5vjcpltkg9uwpqf5apt9","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task16","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTkwIiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: AED96399051AB5388AB630F81228061606F89CD5435CAE777EE6286B519AE4E0
```

### A17 i --(1)--> s --(1)--> j --(1)--> s --(1)--> i
- i --(1)--> s 
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task17 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task17
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: B688B5D07E47C663498A17270EC7BC961A7132EACA4B34D64976EFD1E0FADF80

ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "task17"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3757909"
      revision_number: "1"
    timeout_timestamp: "1678246191437435878"
    token_ids:
    - task17
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: A2E51DE8DB6AF04618D85A4AFD215206CC6A6C75AADE17CCB8C7244D9552DB9B
```
- s --(1)--> j 
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task17", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task17","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 398B093287884E44F79069A107DEEBD8D83A9BB372F0FBBAF0AB762987D66F7E
```
- j --(1)--> s 
```shell
ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task17", "msg": "ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task17","msg":"ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 6E5C7205E7A050672ACBF0924396B7AD42F5CD260DC3AD8BC942D1906ACEC41B
```
- s --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task17", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5OTkyMzY5MAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
password must be at least 8 characters
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task17","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5OTkyMzY5MAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: C90D0F9B2FC9D10F1ECF48F6AB85C1F8839662F4EDC253FB2EFE05482F441600
```

### A18 i --(1)--> o --(1)--> u --(1)--> o --(1)--> i
- i --(1)--> o
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task18 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task18
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 4B8E6E6497DB887D0DE366CC5BF1A39FDC0FA5412B600DE5D1EEF50BE1EB6D13


ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-0" $oAddr "hashquarklogo" "task18"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-0
    source_port: nft-transfer
    timeout_height:
      revision_height: "445608"
      revision_number: "1"
    timeout_timestamp: "1678247302282751292"
    token_ids:
    - task18
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 75F81CB6BF4498616187DFA758230EDE93B664B88751EC878351B7748A6012E4
```
- o --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-24/hashquarklogo
hash: A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-41" $uAddr "ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE" "task18"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE
    memo: ""
    receiver: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-41
    source_port: nft-transfer
    timeout_height:
      revision_height: "2387954"
      revision_number: "2"
    timeout_timestamp: "1678247426519019752"
    token_ids:
    - task18
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 59DF4666C8541303A6B2FB125CC1650B77E071534EEA6067D5EA9046520FDE6E
```
- u --(1)--> o
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-5/nft-transfer/channel-24/hashquarklogo
hash: 26B2C580198F965F5C71834C72187BA022FCB857CA6F681B467C5555A5EDAF5C

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/26B2C580198F965F5C71834C72187BA022FCB857CA6F681B467C5555A5EDAF5C" "task18"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/26B2C580198F965F5C71834C72187BA022FCB857CA6F681B467C5555A5EDAF5C
    memo: ""
    receiver: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-5
    source_port: nft-transfer
    timeout_height:
      revision_height: "445679"
      revision_number: "1"
    timeout_timestamp: "1678247600613839248"
    token_ids:
    - task18
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 8A5261E00B84373250FD06E4A8B5D53D7C17DF6BA79E7D34AE73795218443905
```
- o --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-24/hashquarklogo
hash: A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE


ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" $iAddr "ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE" "task18"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/A68BCAD6B702229DDBCF4EF37870F283AB68F8E10920A423FEE6423A112419CE
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-24
    source_port: nft-transfer
    timeout_height:
      revision_height: "358569"
      revision_number: "1"
    timeout_timestamp: "1678247684619349600"
    token_ids:
    - task18
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 8A77BC84615635EF25B6F0127D8D3E36BC58DC8137F282B4F36AFB69154B3B45
```


### A19 i --(1)--> s --(1)--> j --(1)--> u --(1)--> j --(1)--> s --(1)--> i
- i --(1)--> s
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task19 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task19
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 999742B64BE94EB1BD4A073058CAD986AE2D753490432C9A6D5D1200BEC4FE0C

ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "hashquarklogo" "task19"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-22
    source_port: nft-transfer
    timeout_height:
      revision_height: "3759271"
      revision_number: "1"
    timeout_timestamp: "1678254129352008795"
    token_ids:
    - task19
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 451CFF690FA12F9DA3E625D29FA505872B377812218EF68293AACB561EB4008D
```
- s --(1)--> j
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task19", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task19","msg":"ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMTEiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDMxOTIKICAgICAgICB9CiAgICB9Cn0="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 756A79E8899D8B92DAF2F2A796EED2BEDA281EACF977D19DFCBCF6E4846F041A
```
- j --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62


ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task19", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task19","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 5DCB039F2D6E54DC3BAE2FFB08CB9E02A1F6BE488163F75670CEC795C2613AB8

```
- u --(1)--> j 
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo
hash: 434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-7" $jAddr "ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613" "task19"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/434FA693A2074C9C3EAE33101C1660C9D15DF9A19C6FCBA43164E26B216CD613
    memo: ""
    receiver: juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-7
    source_port: nft-transfer
    timeout_height:
      revision_height: "435600"
      revision_number: "6"
    timeout_timestamp: "1678254788878497683"
    token_ids:
    - task19
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 58E920B5B66102FFFF67B456FD658BBD0EE70E8DEEE596E614EF7401C0AABAD0
```
- j --(1)--> s 
```shell
ubuntu@10-7-95-161:~/go/bin$ junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-93/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6
data: juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62

ubuntu@10-7-95-161:~/go/bin$ junod tx wasm execute juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "task19", "msg": "ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDQ4MzgKICAgICAgICB9CiAgICB9Cn0K"}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"juno15neaqq7kyewa6k3dtdhrgl6dk9pmj7um5fdw6q","contract":"juno1j5h40tf3ukk63yqmc43m9ef9w4kzk7scaf2fqwfk26x3m8gtkrcqhkxw62","msg":{"send_nft":{"contract":"juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a","token_id":"task19","msg":"ewogICAgInJlY2VpdmVyIjogInN0YXJzMTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtazhlZ2tkIiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtOTMiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTM3NDQ4MzgKICAgICAgICB9CiAgICB9Cn0K"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"ujunox","amount":"5000"}],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: DE158908FA42478579E070736FF3BCCFFAC4D9CC9A82BDF46893E61C323E45F6

```
- s --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/hashquarklogo"}}'
data: stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task19", "msg": "ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5OTkyMzY5MAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1qdqqw4uvf0ezkz7enestuq9jw0mtvqutl5cluz35nl3dwnxuqy4skpmz8d","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task19","msg":"ewogICAgInJlY2VpdmVyIjogImlhYTE1bmVhcXE3a3lld2E2azNkdGRocmdsNmRrOXBtajd1bWhld3lsZCIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwNyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5OTkyMzY5MAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: DABB45EE838389CC3812024E14A89C5DF39E20F5699D0C15E16689AEA8E5B1C5
```

### A20 i --(1)--> u --(1)--> s --(1)--> o --(1)--> s --(1)--> u --(1)--> i
- i --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ iris tx nft mint hashquarklogo task20 --name="the hashquark logo" --uri ="https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png" --from hashquark --recipient="iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld" --data="{\"github_username\": \"https://github.com/HashQuark-Research1\", \"discord_handle\": \"James Wang#8021\", \"team_name\": \"hashquark\", \"community\": \"hashquark\"}" --chain-id gon-irishub-1 --fees 20000uiris --node="tcp://34.80.93.133:26657"
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /irismod.nft.MsgMintNFT
    data: '{"github_username": "https://github.com/HashQuark-Research1", "discord_handle":
      "James Wang#8021", "team_name": "hashquark", "community": "hashquark"}'
    denom_id: hashquarklogo
    id: task20
    name: the hashquark logo
    recipient: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: =https://ipfs.io/ipfs/QmZr3L2HdgYh3bWCTmGgSjnvxZsxtkjwKtP5uTL4T7M97W?filename=logo.png
    uri_hash: ""
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 857418CE90D02998739FC468A6A0C31F572B020EB34F783955F2794EDB368F90


ubuntu@10-7-95-161:~/go/bin$ iris tx nft-transfer transfer "nft-transfer" "channel-17" "uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3" "hashquarklogo" "task20"  --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "20000"
      denom: uiris
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: hashquarklogo
    memo: ""
    receiver: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    sender: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    source_channel: channel-17
    source_port: nft-transfer
    timeout_height:
      revision_height: "2389389"
      revision_number: "2"
    timeout_timestamp: "1678255239160340887"
    token_ids:
    - task20
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 6F2E8C060F2751DFD44683B26DFEE6B545998EEE829E49D6989FCD1FEE75107E
```
- u --(1)--> s
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-3/hashquarklogo
hash: 855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-6" $sAddr "ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F" "task20"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-6
    source_port: nft-transfer
    timeout_height:
      revision_height: "3759360"
      revision_number: "1"
    timeout_timestamp: "1678255381170026869"
    token_ids:
    - task20
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 397C63BEE08772E628108B20EF14CB5A7405E92AA8477F89676156D9E512F171
```
- s --(1)--> o
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-203/nft-transfer/channel-3/hashquarklogo"}}'
data: stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task20", "msg": "ewogICAgInJlY2VpdmVyIjogIm9tbmlmbGl4MTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtbDlsdjJ6IiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtMjA5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task20","msg":"ewogICAgInJlY2VpdmVyIjogIm9tbmlmbGl4MTVuZWFxcTdreWV3YTZrM2R0ZGhyZ2w2ZGs5cG1qN3VtbDlsdjJ6IiwKICAgICJjaGFubmVsX2lkIjogImNoYW5uZWwtMjA5IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDkzNzQ1NjM0CiAgICAgICAgfQogICAgfQp9"}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 8FCF811811FA3BCD044CE919CA6FFCF20BB84EF6DE850F2B35AB2435BC365E9D
```
- o --(1)--> s
```shell
ubuntu@10-7-95-161:~/go/bin$ omniflixhubd q nft-transfer class-hash nft-transfer/channel-44/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-203/nft-transfer/channel-3/hashquarklogo
hash: 373358B078678E7DD87D7EFB3F84441CDD507A33B61CA8989F3812BCCA73F1C9

ubuntu@10-7-95-161:~/go/bin$ omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-44" $sAddr "ibc/373358B078678E7DD87D7EFB3F84441CDD507A33B61CA8989F3812BCCA73F1C9" "task20"  --from=hashquark --fees 200uflix
Enter keyring passphrase:
auth_info:
  fee:
    amount:
    - amount: "200"
      denom: uflix
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/373358B078678E7DD87D7EFB3F84441CDD507A33B61CA8989F3812BCCA73F1C9
    memo: ""
    receiver: stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd
    sender: omniflix15neaqq7kyewa6k3dtdhrgl6dk9pmj7uml9lv2z
    source_channel: channel-44
    source_port: nft-transfer
    timeout_height:
      revision_height: "3759697"
      revision_number: "1"
    timeout_timestamp: "1678255992218511161"
    token_ids:
    - task20
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: AFB2173AB19980389BCAA8B0D83DB0191F36AD1F2FC70A90DE3AC79A5A7D9A5B
```
- s --(1)--> u
```shell
ubuntu@10-7-95-161:~/go/bin$ starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-203/nft-transfer/channel-3/hashquarklogo"}}'
data: stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he

ubuntu@10-7-95-161:~/go/bin$ starsd tx wasm execute stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "task20", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}}' --from hashquark --gas 2000000
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/cosmwasm.wasm.v1.MsgExecuteContract","sender":"stars15neaqq7kyewa6k3dtdhrgl6dk9pmj7umk8egkd","contract":"stars1jyvgaugjqz9d2q8l9l0594fuaqluz4577nc6jzu89hkq557lkztqdq74he","msg":{"send_nft":{"contract":"stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh","token_id":"task20","msg":"ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTIwMyIsCiAgICAidGltZW91dCI6IHsKICAgICAgICAiYmxvY2siOiB7CiAgICAgICAgICAgICJyZXZpc2lvbiI6IDksCiAgICAgICAgICAgICJoZWlnaHQiOiA5Mzc0NTYzNAogICAgICAgIH0KICAgIH0KfQ=="}},"funds":[]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"2000000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 580EED1987E2DD7923633ECBCD37C98A8DC796149E6B9B2D9AA5E305D300C794
```
- u --(1)--> i
```shell
ubuntu@10-7-95-161:~/go/bin$ uptickd q nft-transfer class-hash nft-transfer/channel-3/hashquarklogo
hash: 855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F

ubuntu@10-7-95-161:~/go/bin$ uptickd tx nft-transfer transfer "nft-transfer" "channel-3" $iAddr "ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F" "task20"  --from=hashquark
Enter keyring passphrase:
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /ibc.applications.nft_transfer.v1.MsgTransfer
    class_id: ibc/855BF112D0CC5F4527AA906462DD6D40FCD8D46F45D46705900FABA197DF162F
    memo: ""
    receiver: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    sender: uptick1e8c5yx03fshg9wku0sgc2e2xa8hureyr2pz6g3
    source_channel: channel-3
    source_port: nft-transfer
    timeout_height:
      revision_height: "359991"
      revision_number: "1"
    timeout_timestamp: "1678256226431049078"
    token_ids:
    - task20
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 90D9AF6749093C3989FE49E818CE54EE4BA80038C3D986AE32CEAB25067C1154
```