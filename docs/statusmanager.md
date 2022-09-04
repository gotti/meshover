# Status Manager

ノードのメトリクス監視を行うサーバ。クライアントがpushしたデータをprometheus形式に変換してエクスポートする。

実装は`cmd/exporter/main.go`

クライアントから得たデータは30秒で破棄し、prometheus形式への変換は15秒おきに行なっている。

BGPがEstablishedになっているものを集計してGrafana node graph用のメトリクスを作っている

## 表示

![](../img/grafana-example.png)

## 設定例

![](../img/grafana-settings.png)
