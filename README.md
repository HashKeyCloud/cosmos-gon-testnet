# cosmos-gon-testnet

Cosmos GoN testing activities are still ongoing. This program is used to monitor the creation of all NFT Collections, NFT mint and track the NFT transaction process. This program currently implements the following functions.
- Monitor NFT collection creation events；
- Monitor all NFT mint events；
- Monitor a specific NFT mint event through the NFT class id；
- Monitor all NFT transaction events；
- Monitor a specific NFT transaction event through the NFT base class id and NFT id；

## 1、Roadmap
- Added dashboard service to more vividly display the NFT transaction process;
- Optimize the event monitoring method to improve the response speed;

## 2. Hardware requirements
We recommennd the following hardware resources:
- CPU: 4 cores
- Memory: 8GM RAM
- Database: PostgreSQL

## 3. configuration file
- `vim /data/config/conf.yaml`
- Configuration example
```yaml
gRpc:
 irisIp: "iris grpc host ip"
 irisPort: "iris grpc host grpc port"

  junoIp: "juno grpc host ip"
  junoPort: "juno grpc host grpc port"

 omniflixIp: "omniflix grpc host ip"
 omniflixPort: "omniflix grpc host grpc port"

 stargazeIp: "stargaze grpc host ip"
 stargazePort: "stargaze grpc host grpc port"

 uptickIp: "uptick grpc host ip"
 upticePort: "uptick grpc host grpc port"

monitor:
  timeInterval: 600  # Monitoring time interval. 600 means monitor transaction section every 600 seconds

log:
  path: "logPath"
  level: "info"
```

## 4. Deployment monitoring
- 4.1 Install go
> TIP
>
> go 1.20++ is required for building and installing the software.

[Golang Official website](https://go.dev/doc/install)

- 4.2 Build monitor binary file
```shell
git clone https://github.com/HashQuark-Research1/cosmos-gon-testnet.git

cd cosmos-gon-testnet
vim main.go
Replace the IP in @host 127.0.0.1:8080 with the IP of the server running the rebate program
go install github.com/swaggo/swag/cmd/swag@latest
swag init –parseInternal depends on –parseDependency -g ./handler/handler.go -o ./docs

go build -o gon
```

- 4.3 Edit configuration file
```shell
vim config/conf.yaml
```

- 4.4 Automating your monitor with systemd
```bash
vim /etc/systemd/system/gon.servic
```
```shell
[Unit]
Description=GoN Daemon
After=network.target
[Service]
User=ubuntu
ExecStart=/home/ubuntu/go/bin/gon -c /data/config/.conf.yaml
KillSignal=SIGINT
Restart=on-failure
RestartSec=5s
LimitNOFILE=1000000
[Install]
WantedBy=multi-user.target
[Manager]
DefaultLimitNOFILE=1000000
```

```shell script
systemctl start gon.servic
```