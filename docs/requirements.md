# Requirements

meshoverは以下を要求する。

## Required

- dockerまたはcontainerd(nerdctl)
- IPv6アドレス(グローバルなIPであり全ノードがNATなしで疎通できること。)
- wireguard-tools
- root権限

## Optional

- iptables (NATを利用するときのみ)
- Kubernetes/Coil/PureLB (コンテナやVMを実行するノードのみ)
