# Getting Started

## はじめに

meshoverはオーバーレイネットワーク構築を自動化する実験用ツールです。以下の操作を自動で行います。

- ノードへのIPアドレス割り当て(ノードIPをloopbackに、トンネルIPをwireguardにそれぞれ割り当て)
- FRRコンテナの立ち上げと経路の交換
- 通信を可能にするためのカーネルパラメータの設定
- 接続通信状況のモニタリング

## 要求

- コントローラ用サーバ
  - 全ノードからアクセス可能であること
- wireguard-tools
- GREトンネルを張るためのモジュール(必要あれば)

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

## 通常ノードの構築

手元のパソコンなどで`cmd/admin`をビルドし`./admin gen`を実行してください。エージェントトークンが発行されます。

```
$ go run cmd/admin/main.go gen
1 < hexadecimal agent token >
```

以下は対象ノードでの設定です。

各ノードでFRRの設定ファイルテンプレートを`./conf/frr.conf`に準備してください。標準的なテンプレートは以下のものです。このテンプレートではBGPピアを確立し、カーネルのテーブルなどから`192.168.0.0/16`に含まれない経路を広報します。


https://github.com/gotti/meshover/blob/main/conf/frr.conf

Coilと連携する場合は以下のテンプレートを参考にしてください。VRFを作成し経路を読み込んで広報しています。

https://github.com/gotti/meshover/blob/main/conf/frr.conf.withcoil

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
- CNI: Cilium https://github.com/cybozu-go/coil
- LB: PureLB https://purelb.gitlab.io/docs/

#### Kubernetesノードの必要ツールのインストール

Kubernetesを導入する準備をします。

要求に書いてあるものに加えcontainerdを導入してください。

#### Kubernetesのインストール

kubeadmにて通常の手順でインストールしてください。`control-plane-endpoint`や`apiserver-advertise-address`などはmeshoverのものを設定してください。

#### CNIのインストール

CNIはmeshoverと機能が被らない設定ができるものを選択してください。私はCilium(Native Routing)を利用しているのでCiliumの設定を説明します。

Ciliumのvaluesはこんな感じです。

```yaml
tunnel: disabled # meshoverがトンネルを作成するため必要ありません
ipv4NativeRoutingCIDR: 10.0.0.0/8 # meshoverに管理させたIPを指定してください。IPマスカレードなどを切ることができます。
enableRuntimeDeviceDetection: true # meshoverは動的にGREデバイスを作成します。Ciliumは初回しかデバイスを検出しないためこの設定が必要です。
enableIPv6Masquerade: false # お好みで
cleanBpfState: false # デバッグ用、お好みで
cleanState: false # デバッグ用、お好みで
ipam:
  operator:
    clusterPoolIPv4PodCIDRList:
      - 10.228.0.0/16
    clusterPoolIPv6PodCIDRList:
      - xxxx::/104
```






以下のオプションでCoilと連携することができます。

```
-coiladvertise # Enable Coil advertising
-coilnatsources < Replace with your Coil pod private ip range with CIDR > # Set pod IPs, natting outgoing packets from containers
-frr nerdctl
```
