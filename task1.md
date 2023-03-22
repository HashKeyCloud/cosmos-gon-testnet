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











# i --(1)--> s --(1)--> j --(1)--> u --(1)--> o --(1)--> i
> export classId="gonTeamRace1"
> export nftId="gtr1/HashQuarkResearch1"
- i --(1)--> s
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonTeamRace1" "gtr1/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
533B0A684643583AE848197E55A961E950BED5D38C08FD3E47FDCDDF9CC7F0A8
```
- s --(1)--> j
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace1"}}'


starsd tx wasm execute stars1pn5jfda6vyjaqxz0zmkvdjwujsj0jwczlstuae056un6rx4vpxgqvhgtw8 '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gtr1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
6B00D43F9AA940F0429A3FF61406102FD3F2549EC0FE34811C0661B8487BF266
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace1"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6

junod tx wasm execute juno1uxts9plgp58rscm670zpje5wm9muw3v5ry4d4vr84grcqmj5ld5sherqlv '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gtr1/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
06EA1A5E6B310F97F3C402CF193B3F5224FA3F330892C2A07F9AB69A679D3CCD
```
- u --(1)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace1

uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/20D87745B27FEA686710D22A45B182EEC787214B2C19F3C268D1F67549CF5D78" "gtr1/HashQuarkResearch1"  --from=hashquark
61AA2AC463DDCFD1795841421F31CFA3CD67CBFCCE16EF8E0C05D3494F158484

```

- o --(1)--> i
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace1



omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" $iAddr "ibc/FC543FF67C7ECEFE9961650C59EC6DB908965797C0B8DADDBB5C6C1FBF7486E9" "gtr1/HashQuarkResearch1"  --from=hashquark --fees 200uflix
F51C29FAD279B8A0458207BBC580C21F2076C00801BA0230D990FD95F9C4B194
iris q nft-transfer class-hash nft-transfer/channel-0/nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace1

iris tx nft transfer "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/665C1369FB9374738FB3A960968F1D40693F7FCBE25A169F2066486B02045B87" "gtr1/HashQuarkResearch1" --from=hashquark  --chain-id gon-irishub-1 --fees 20000uiris
1E1DD65128EC2DC0C2984B8B76F9FC77E5D96704840270A8DA85AD0A64D3BEEE
```




# i --(1)--> s --(1)--> j --(1)--> u --(1)--> o --(2)--> i
> export classId="gonTeamRace2"
> export nftId="gtr2/HashQuarkResearch1"
- i --(1)--> s
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonTeamRace2" "gtr2/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
49410F1583E0BB97086847AB00BCEEEB3CF5B749500D3870C2A20F33D2FDFA15
```
- s --(1)--> j
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace2"}}'


starsd tx wasm execute stars1unjgjcyhhx5r8tn96hqs92xdv6d0hy64xw55qu4um8awvy5qdulq0ggj28 '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gtr2/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace2"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6

junod tx wasm execute juno1c2m04cdfvt9atwzc8vf0dhg0w6aygnw84qxkgf87majnt5ggsn2smd43q5 '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gtr2/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
```
- u --(1)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace2

uptickd tx nft-transfer transfer "nft-transfer" "channel-5" $oAddr "ibc/A610F72F743AC5B5521590223CF5FCEE089236A1277ADC99D8E7B3C894243B8E" "gtr2/HashQuarkResearch1"  --from=hashquark
F4D7A39DE7AAC1B66D8C0A7295F56537583A196A11BD83CB13B0DDE27057C21B
```

- o --(2)--> i
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace2



omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-25" $iAddr "ibc/8D4332B5DD00E1BA11E835052B1B6FF6DBC1B3991587C6A658B900069498A32A" "gtr2/HashQuarkResearch1"  --from=hashquark --fees 200uflix
9ABDF05E6A63B51BE1B3640B029A4AE64250896A5527610B3F13EA2D7BC9C19F

iris q nft-transfer class-hash nft-transfer/channel-1/nft-transfer/channel-41/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace2

iris tx nft transfer "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/7168E1ADE6CC7DCFE0AC0B1252B00213479E548865AE0BF104AF958380884A47" "gtr2/HashQuarkResearch1" --from=hashquark  --chain-id gon-irishub-1 --fees 20000uiris
7EDECC1908A9D89D9AE63B956C6EF675D09597A39A381CE2C1D73409964BA0B6
```


# i --(1)--> s --(1)--> j --(1)--> u --(2)--> o --(1)--> i
> export classId="gonTeamRace3"
> export nftId="gtr3/HashQuarkResearch1"
- i --(1)--> s
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonTeamRace3" "gtr3/HashQuarkResearch1" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
89B752F6BFCCDDC13088FB31FC2954FCB8E77305CD9DE62FAF181403FF6F96C0
```
- s --(1)--> j
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace3"}}'


starsd tx wasm execute stars1r0frzufzzft5qkkw959apapjmk8ywlsy5dr2fa6qg8hskemz5fys8ckwj9 '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "gtr3/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
3646825B52B2A617374F9EE10D40A09D18FB9E97BA2F32DFAECF9F25B8015CAD
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace3"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6

junod tx wasm execute juno10vxjy8s9zqmvvjcf8y5j4q4a8w9eendxd4yk56jq7fqhmln7mv4s94qe8r '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "gtr3/HashQuarkResearch1", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
072F3CBD4E829E419719B37838D89B12E377CC99217C4DF5E247DA55BC58D782
```
- u --(2)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace3

uptickd tx nft-transfer transfer "nft-transfer" "channel-9" $oAddr "ibc/C3AA94AD394BBEF6998033DD2A702978D77D71749DBD3CCC9F2C3AFCC8CCDFA0" "gtr3/HashQuarkResearch1"  --from=hashquark
1393D9D8FE53A7D6D1EF91CE50F487F4F222B92181D6EEEE2322711BDB6B27E6

```

- o --(1)--> i
```shell
omniflixhubd q nft-transfer class-hash nft-transfer/channel-42/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace3



omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-24" $iAddr "ibc/DB716901E3B773E5E50EBBE1A2815BFADCEBFE6C357BD329C78D9545FD6A5478" "gtr3/HashQuarkResearch1"  --from=hashquark --fees 200uflix
B18F6F698D90DD81120E1C14FBB56B4FFD783DA955FB3E79F500509C04C499BB
iris q nft-transfer class-hash nft-transfer/channel-0/nft-transfer/channel-42/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonTeamRace3

iris tx nft transfer "iaa1488wwr235vka7j722hzacpk0plxw33ksqyneuz" "ibc/C8E7EF402A9CA7495BA642524DEC23431977FECC646FD824140D9B6D4D44B1C8" "gtr3/HashQuarkResearch1" --from=hashquark  --chain-id gon-irishub-1 --fees 20000uiris
6A92801471C339B93952A87FF93E36E057E7E30B5CC44D1B59CABA7260BC5978
```






# i --(1)--> s --(1)--> j --(1)--> u --(2)--> o --(2)--> i
- i --(1)--> s
```shell
iris tx nft-transfer transfer "nft-transfer" "channel-22" $sAddr "gonQuiz" "quiz049" --from=hashquark --node="tcp://34.80.93.133:26657" --chain-id gon-irishub-1 --fees 20000uiris
764943C17FE68495157FC78780DE23183FF56DE0AAB20586D9F719EA952838D3
```
- s --(1)--> j
```shell
starsd q wasm contract-state smart stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh '{"nft_contract": {"class_id" : "wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonQuiz"}}'


starsd tx wasm execute stars1ecwr8jj333jm9l6l0l6rqgt478h086yq07z78w9wg6wrdpf2ajgqfk03h9 '{"send_nft": {"contract": "stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh", "token_id": "quiz049", "msg": "ewogICAgInJlY2VpdmVyIjogImp1bm8xNW5lYXFxN2t5ZXdhNmszZHRkaHJnbDZkazlwbWo3dW01ZmR3NnEiLAogICAgImNoYW5uZWxfaWQiOiAiY2hhbm5lbC0yMzAiLAogICAgInRpbWVvdXQiOiB7CiAgICAgICAgImJsb2NrIjogewogICAgICAgICAgICAicmV2aXNpb24iOiA5LAogICAgICAgICAgICAiaGVpZ2h0IjogOTkzNzQzMTkyCiAgICAgICAgfQogICAgfQp9"}}' --from hashquark --gas 2000000
01DA64A680BEDA4B8FD0351DA8925056BA3AC3F55590B7126D8F121174742419
```
- j --(1)--> u
```shell
junod q wasm contract-state smart juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a '{"nft_contract": {"class_id" : "wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonQuiz"}}' --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6

junod tx wasm execute juno1v8juugyz2nym2dhaa0309w98exuxa46geauvw8c8k9e73a5sq64q5j6wlj '{"send_nft": {"contract": "juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a", "token_id": "quiz049", "msg": "ewogICAgInJlY2VpdmVyIjogInVwdGljazFlOGM1eXgwM2ZzaGc5d2t1MHNnYzJlMnhhOGh1cmV5cjJwejZnMyIsCiAgICAiY2hhbm5lbF9pZCI6ICJjaGFubmVsLTg2IiwKICAgICJ0aW1lb3V0IjogewogICAgICAgICJibG9jayI6IHsKICAgICAgICAgICAgInJldmlzaW9uIjogOSwKICAgICAgICAgICAgImhlaWdodCI6IDk0MTk3NjMKICAgICAgICB9CiAgICB9Cn0="}}' --from hashquark  --node="https://rpc.uni.junonetwork.io:443" --chain-id uni-6 --fees 5000ujunox --gas 2000000
6DB8EB73E7FE8CB4C2520445039233923A6FA83B5550CD6E124E4A911821FCEC
```
- u --(2)--> o
```shell
uptickd q nft-transfer class-hash nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonQuiz

uptickd tx nft-transfer transfer "nft-transfer" "channel-9" $oAddr "ibc/C1F6D7B5F4F9CE67F9133F7A96E009378551AB5975D15DC909859423B7D4AF74" "quiz049"  --from=hashquark
BC8142E937934AE67194FDF5C991EEBC5B8EEE0989D41BBFDD3F882F3B427632

```

- o --(2)--> i
```shell
omniflixhubd q nft-transfer class-hash channel-1/nft-transfer/channel-42/nft-transfer/channel-7/wasm.juno1stv6sk0mvku34fj2mqrlyru6683866n306mfv52tlugtl322zmks26kg7a/channel-120/wasm.stars1ve46fjrhcrum94c7d8yc2wsdz8cpuw73503e8qn9r44spr6dw0lsvmvtqh/channel-207/gonQuiz

omniflixhubd tx nft-transfer transfer "nft-transfer" "channel-25" $iAddr "ibc/E572D664BA7341856CDFC3025E377078D1A61BB141F70D1DAB96B6BA3CB76FF5" "quiz049"  --from=hashquark --fees 200uflix
7673A6ECB0A8A7BCCF86C9A9AA16487D3F5107ABA75463CF626507F7A16F4968
```