# IPアドレス

meshoverは以下の3種類のIPアドレスを割り当て、管理します。

- BaseAddress
  - v4/v6のどちらかで1つのみ
  - ノードに割り当てられるIPアドレスです。一般的な通信においてこれを使います。
- AdditionalAddresses
  - v4/v6で任意の個数
  - ノードに対して追加で割り当てるIPアドレスです。必要があれば説明を添えることができます。
- TunnelAddress
  - v4を1つのみ
  - トンネルに対して割り当てられるIPアドレスです。トンネルの確立に使われますがそれ以外に使われません。
