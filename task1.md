# 第三轮
## 查询iris地址下的NFT
```shell
iris  q nft owner $iAddr
```
## f1
```shell
data: '{"type":"individual round 1","flow":"f01","last_owner":"iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz","start_height":"473000"}'
    id: gir1/HashQuarkResearch1
    name: gir1/HashQuarkResearch1
    owner: iaa15neaqq7kyewa6k3dtdhrgl6dk9pmj7umhewyld
    uri: https://github.com/HashQuarkResearch1
    uri_hash: ba09c44c3ac07366d6ed1689cef91fba1f5f5a927a2430a237c8dc213ec7fe86


{
  "type": "individual race round",
  "flow": "the flow id, check flow with flow-id on https://github.com/game-of-nfts/gon-testnets/blob/main/doc/flow-table.md",
  "last_owner": "the ultimate owner of this NFT",
  "start_height": "transfer before this height are considered invalid"
}
```
> i --(1)--> s --(1)--> j --(1)--> u --(1)--> o --(1)--> i
> export classId="gonIndivRace1"
> export nftId="gir1/HashQuarkResearch1"
- i --(1)--> s 
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonIndivRace1" "gir1/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
```
- s --(1)--> j
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonIndivRace1"}}'











starsd tx wasm execute stars1kd7u4jucpc7q5sm4f4rsy3v7grj9gty6zksz8tezz79uxvse63rqhwvgjp '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gir1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonIndivRace1"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6










junod tx wasm execute juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gir1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
```
- u --(1)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonIndivRace1




uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/AA13DC4FBF98AFE6E3ADFFA35134C402226F4A41E152007BBCC16A10336656C0" "gir1/HashQuarkResearch1"  --from=hashquark

```



- o --(1)--> i
```shell
omniflixhubd q nft-transfer class-hash /nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonIndivRace1



omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/2996F3E187594D0500C2FDE2A1457BCD47C28B03FA01C9BD747955012D53F9DF" "gir1/HashQuarkResearch1"  --from=hashquark --fees 200uflix
```












```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonIndivRace1" "gir1/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris

318B33896BDCCCAED94EB7ACAAB2E7F84CB369C8909BD7360D55F0C73A146BC8

starsd tx wasm execute stars1kd7u4jucpc7q5sm4f4rsy3v7grj9gty6zksz8tezz79uxvse63rqhwvgjp '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gir1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000

EF93778E19A85A0B76D6BADEA83E6D1B0B1275BEEA40E3E7AA20E1A57FC28223

junod tx wasm execute juno1kk0k8f7aq4659zds8072m2pulheeehmcn0u0u8npmvtapffxukkqld6sck '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gir1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000

02DD717DFB18E51E096022946454B34306C10AB912182C4F8BA178B98FBE59AE



uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/AA13DC4FBF98AFE6E3ADFFA35134C402226F4A41E152007BBCC16A10336656C0" "gir1/HashQuarkResearch1"  --from=hashquark

86870664BE9DAE8C4BD0B81F77B1E7BCCBA429B722755104CB283A97A88CDEB4

omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" $iAddr "ibc/2996F3E187594D0500C2FDE2A1457BCD47C28B03FA01C9BD747955012D53F9DF" "gir1/HashQuarkResearch1"  --from=hashquark --fees 200uflix

25F85E8C54CDD1C85DE3396C75350714BE7DB75358EE9BAB446F22C5B0B25084


iris tx nft transfer "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/571895A89A58FFE9FE57C36DC949D647C0B6A900FCA2531905CAECCA30FC86DB" "gir1/HashQuarkResearch1" --from=hashquark  --chain-id gon-irishub-1 --fees 20000uiris
```








## 3.2 i --(2)--> s --(1)--> j --(2)--> u --(2)--> o --(2)--> i
- i --(2)--> s
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-23" $sAddr "gonIndivRace2" "gir2/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris

EEAB8258CB83C1AB3E58DF2374C54E6DF58AB3E7678D368CC6BC59330A0AE88E
```
- s --(1)--> j 
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-208/gonIndivRace2"}}'
stars126vew88r9askv54wdee4rzh3j8squhwyu3rw52emknxf5ztnfpsq49shvt

starsd tx wasm execute stars126vew88r9askv54wdee4rzh3j8squhwyu3rw52emknxf5ztnfpsq49shvt '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gir2/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000

9EFC8ED5EADC07C90DDA04FB3DF310B0D2ABA399AF521011BBCE8444783B2D35
```
- j --(2)--> u 
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-208/gonIndivRace2"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6

juno19df5p5msv0yjnjlljgqupffm7np6cwh0fnj4lk9a6zxjn95qh20sm88jwn

junod tx wasm execute juno19df5p5msv0yjnjlljgqupffm7np6cwh0fnj4lk9a6zxjn95qh20sm88jwn '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gir2/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg4IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
9B79859EABC026E5E97F1A71AE29AAA3CA7ABDDFB037C5B557A42CCD3EFB1A1D
```
- u --(2)--> o 
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-13/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-208/gonIndivRace2
55D4954C669FCB03F63D745DB1B1CB8CBAECAE8CB37156B9DA848D4640D0B2D8



uptickd tx nft-transfer transfer "nft-transfer" "channel-9" $oAddr "ibc/55D4954C669FCB03F63D745DB1B1CB8CBAECAE8CB37156B9DA848D4640D0B2D8" "gir2/HashQuarkResearch1"  --from=hashquark
2C70CEEB855D74F3F03FE4DA654DCE69735F5A41D850824A6A86D5BCCF2BD83F
```
- o --(2)--> i
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-42/nft-transfer/channel-13/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-208/gonIndivRace2
4BFF83EA97A18C3A9B5351477661746023B6D8117D4F55404BD471C0B0DAD8F8



omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-25" $iAddr "ibc/4BFF83EA97A18C3A9B5351477661746023B6D8117D4F55404BD471C0B0DAD8F8" "gir2/HashQuarkResearch1"  --from=hashquark --fees 200uflix
3033EBD11E82D2EE3AB842520FCF2BDA6D0D92397E306DAC713038EA80969AE5
```

i-->i

```shell
iris q nft-transfer class-hash nft-transfer/channel-1/nft-transfer/channel-42/nft-transfer/channel-13/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-208/gonIndivRace2

iris tx nft transfer "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/0CE12429541F00548543A34E4D1B83B4E235E792457DDE117E85CF21459993BB" "gir2/HashQuarkResearch1" --from=hashquark  --chain-id gon-irishub-1 --fees 20000uiris
66BDAF6C1E0391169C7C436018F05D7DA7038137A1FE36FED23E5BF9152ACD47
```