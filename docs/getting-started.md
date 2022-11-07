# Getting Started

## はじめに

meshoverはオーバーレイネットワーク構築を自動化する実験用ツールです。以下の操作を自動で行います。

- ノードへのIPアドレス割り当て(ノードIPをloopbackに、トンネルIPをwireguardにそれぞれ割り当て)
- FRRコンテナの立ち上げと経路の交換
- 通信を可能にするためのカーネルパラメータの設定
- 接続通信状況のモニタリング


## コントローラの構築

`cmd/server`をビルドし以下の引数で実行してください。コントローラは全ノードからmeshover無しでもアクセス可能である必要があります。

```
-password=< Replace with your strong password for admin cli >
-listen=0.0.0.0:80
-statefile=/pers
```

## ステータスマネージャーの構築

`cmd/exporter`をビルド

TODO

## ノードの構築

手元のパソコンなどで`cmd/admin`をビルドし`./admin gen`を実行してください。エージェントトークンが発行されます。

```
$ go run cmd/admin/main.go gen
1 < hexadecimal agent token >
```

以下は対象ノードでの設定です。

wireguard-toolsをインストールしてください

`cmd/client`をビルドし以下の引数で実行してください

```
-controlserver < Replace with your controller IP >:80
-statusserver < Replace with your statusmanager IP>:80
-agenttoken < Repalce with your agent token >
-frr dockersdk
```

### VMノードの構築

上記の手順でmeshoverをインストールしてください。

Kubernetesをインストールします。ただしコンポーネントとして以下の利用を想定しています。

- コンテナランタイム: containerd
  - nerdctlも必要
- CNI: Coil https://github.com/cybozu-go/coil
- LB: PureLB https://purelb.gitlab.io/docs/

kubeadmでの手順を示します。まずcontainerdをインストールしてください。

```

```

以下のオプションでCoilと連携することができます。

```
-coiladvertise # Enable Coil advertising
-coilnatsources < Replace with your Coil pod private ip range with CIDR > # Set pod IPs, natting outgoing packets from containers
-frr nerdctl
```
