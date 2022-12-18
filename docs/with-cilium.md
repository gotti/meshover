# meshover with cilium

CiliumはeBPFでパケットを転送するCNIです。meshoverと連携するには、最低限以下の設定を入れる必要があります。

values.yaml

```yaml
tunnel: disabled # Cilium自身のVXLANによるオーバーレイネットワークを切る
ipv4NativeRoutingCIDR: 10.0.0.0/8 # Ciliumによる諸々はdstが10.0.0.0/8で切るように
enableRuntimeDeviceDetection: true # meshoverはトンネルを作成したり消したりするので
enableIPv6Masquerade: false
```
