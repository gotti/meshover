# カーネルパラメータ設定

meshoverは、その仕組み上以下のカーネルパラメータを自動で設定する。ノードを動かす際には注意が必要である。

```
net.ipv4.ip_forward=1
net.ipv4.conf.all.rp_filter=0
sysctl", "-w", "net.ipv4.conf.default.rp_filter=0
```
