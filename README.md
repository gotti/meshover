# Meshover

貧者のIP ClosもどきDCを構築するためのツール。

IPv6を使ってWireguardでフルメッシュVPNを構成する。tailscaleやcalicoと似たツール。

tailscaleと違うのはmeshoverを導入したノードに直接接続していないVM/コンテナをネットワークに接続できること。これはBGPを全ノードが喋ることで実現している。VM/コンテナの接続はKubernetesのCoil/PureLBとの連携を想定している。

特徴
- 自動構成IPアドレス
  - 現在はホスト名をキーとして一意なIPアドレスがコントローラから割り当てられる。
- FRRを使ったBGP unnumbered相互接続
- KubernetesのCoil/PureLBとの連携

## 何を解決したい？

- 同一L2の上にあるDHCPなどで割り当てられたアドレスは、そのL2でしか使えない。
- 別の場所へのL2延伸はやりたくない
- そのためサーバ間の接続は全てL3のネットワークを利用するようにしたい。
- tailscaleを使ってみてかなり便利だった
- tailscaleのIPアドレスでKubernetesを動かせばいいのでは？
  - 実装してから気づいたが後述するallowed-ipsの制限のため、この構成はおそらく無理
- BGPを喋らせればCalicoやMetalLBと接続できて完璧なのでは？
  - CalicoとMetalLBは面倒なのでCoilとPureLBになった

## Getting Started

[Getting Started](./docs/getting-started.md)

## 開発状況

- [x] Wireguardでフルメッシュ接続
- [x] GREをWireguardの上に通し後述のallowed-ips制限を回避、またlink local ipv6アドレスが利用可能。
- [x] FRRをdockerで立ち上げて自動生成したBGPの設定を投入する。
  - [x] BGP unnumberedでmeshoverの他ノードとピアリング。
  - [x] テンプレを変更してノード特有の設定(たとえばCoilの提供するPodIPをBGPで広報する)が可能。
- [x] Source IP based routing
  - グローバルIPにSNATされたパケットを特定ノードに転送する。
  - MetalLBでSVCとかVMにグローバルIPを割り当てた際に適切に処理できるようになる。
  - オプション`-gathering <CIDR range>`を指定したノードにその範囲のIPアドレスが集約される。
- [x] Prometheus Exporter
  - Grafana Node Graphを使えばノードと接続状況を見れる
  - [status manager docs](./docs/statusmanager.md)

## 技術

- Go
  - Rustを使ってみたかったがprotoc-gen-validateが非対応で、辛い
- gRPC
  - コントローラへIPアドレス/ASN割り当てリクエストを投げたり接続情報の登録に使う
- FRR
  - ルーティングなど
- Wireguard
  - セキュアなフルメッシュ接続をやる
- GRE
  - Wireguardはallowed-ipsがACLの働きをすると同時にデバイス内でのルーティングにも使われる(1つのデバイスで複数の接続を張れるため)。したがって複数の接続それぞれを0.0.0.0/0で許可できず、BGPでルーティングしてもホストのIPアドレスからの通信以外はその経路を使えない。これをなんとかするためにGREでもう1段トンネルを張っている。
  - link localアドレスが割り当てられるのでBGP unnumberedが可能
  - 参考: https://www.infrastudy.com/?p=1065

## 使い方

### Requirements

- グローバルなIPv6が割り当てられたパソコン n台
  - 検証環境はUbuntu Server 20.04 LTS
  - docker, wireguard-tools

### ビルド

> **Warning**
> 開発中のためお行儀が悪いです。
> 認証やファイアウォールが実装されていません。
> 安全な開発環境でのみ利用してください。


```bash
make # client と server が生成されます
```

サーバを、ノード全てから接続可能なパソコンで動かす。
これはコントローラでありノードの公開鍵と自動構成IPアドレス/ASN、wireguard接続用IPv6アドレスなどを管理します。

```bash
./server -listen 0.0.0.0:12384
```

クライアントをメッシュ接続したいノードで動かします。

```bash
vim conf/frr.conf #FRRの設定テンプレートが存在することを確認し、必要あれば編集
sudo ./client -controlserver <IP Address to server>:12384
```

## TODO

新しく追加したものはIssueにある。ここに残っているものは最初期に作ったTODOのみ。

- wireguard周りのリファクタリング
  - wireguard-tools無しでも動くようにもしたいね
- FRRのListen IPをデフォルトではWireguardのIPだけにする
  - meshover以外とBGP接続するノードのためにファイアウォールもちゃんとやらないと
- BGPのルーティング情報の検証
- テスト整備
- ノードのステータスを監視できるように
- gRPCの認証をやる
  - サーバの認証はSSL/TLS、クライアントはJWTでよさそう
- ノードステータスとノード認証のJWTを発行できるフロントエンド
- コマンドラインオプションで要らないコンポーネントを動かさないように
- 接続性の向上
  - UDPホールパンチングを実装するのが嫌でIPv6必須にしているのでなんとかする
  - UDPホールパンチングをやるか、BGPが動いているのを生かしてIPv4でも繋がるリレーノードを立てるとか
  - どちらにしろ自分のアンダーレイIPアドレスを現在の自己申告ではなく外部のサーバに確認してもらう必要がある
- EVPN
  - せっかくBGPフルメッシュなのでMPLSとかでVPC作りたい
- CNI
  - せっかくEVPN作るならCNI自作してkubevirtとかでVPCに接続できるようにしたい
- VPNをフルメッシュではなく、BGPでいうルートリフレクタのようなノードを自動選出する
  - だいぶ難しそうなので後回し、ひとまずそこまでスケールすることは考えなくてもいいか
  - 経路ごとのコスト(トンネルで全部1ホップになってるから)を計算
  - 実際の通信状況から利用頻度が少ないし別のノード経由でも通信できる接続を落とすとか？
