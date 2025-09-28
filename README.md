# Go Conference 2025: GoのinterfaceとGenericsの内部構造と進化

[@turbofish_](https://x.com/turbofish_)

2025年9月27、28日に開催されたGo Conference 2025の、セッション[「GoのinterfaceとGenericsの内部構造と進化」](https://gocon.jp/2025/talks/951608/)で使用した資料と、調査の際に参考にした情報などをまとめたリポジトリです。

```sh
.
└── go_conference_2025/
    ├── images
    ├── sample_code         アセンブリで遊んでみるために使用したサンプルコード
    ├── base.md / base.pdf  当日使用したスライド（Marp）
    ├── references.md       参考資料のリンク集
    └── README.md
```

- Go version：1.25.1

### アセンブリのダンプ

sample_code に、アセンブリを確認した時に使用したファイルを保存しています。
簡単のため、コンパイラの最適化とインライン化を無効にしてダンプしているものもあります（普通にダンプしているものもあります）。

```sh
go build -gcflags="-N -l" main.go
go tool objdump -s "main\.main" main
```
