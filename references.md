# 参考資料 / References

極力 `golang/go` のソースコードと公式ドキュメントを掲載するようにしていますが、調査の過程で多くの記事を参考にさせていただきました。
ソースコードへのリンクは、登壇準備当時のコミットハッシュ（Goのバージョンは1.25.1）ものです。

## interface{} / any 型、interface
#### 公式ドキュメント、ブログなど
- Go Documentation Frequently Asked Questions (FAQ): [Why is my nil error value not equal to nil?](https://go.dev/doc/faq#nil_error)
- [Interface values](https://go.dev/tour/methods/11) Tour of Go
- [Go Data Structures: Interfaces](https://research.swtch.com/interfaces) 2009-12-01

#### 実装
- [src/internal/abi/type.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/internal/abi/type.go)
- [src/internal/abi/iface.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/internal/abi/iface.go)
  - interface のコンパイル時情報を扱う構造体がある

#### その他
- [GoのTyped-nilの扱い](https://zenn.dev/nobonobo/articles/f554041aea1955) 2022-10-17
- [Go で any型（＝interface{}型）を、「どんな型でも入れられる」型として扱う場合には nil チェックに注意](https://turbofish.hatenablog.com/entry/2025/06/23/071752) 2025-06-23
  - 拙著です
- [Go Internals - Chapter II: Interfaces](https://cmc.gitbook.io/go-internals/chapter-ii-interfaces)
  - [数人で管理している記事のようです](https://github.com/teh-cmc/go-internals)。使用されているGoのバージョンはやや古いものの、大変お世話になりました
  - 第一章の「インターフェース」にて、実質可変長になる配列などについても言及されています

## Generics
#### 公式ドキュメント、ブログなど
- [An Introduction To Generics](https://go.dev/blog/intro-generics) The Go Blog 2022-03-22
  - 型パラメータ、型制約、型セット、型推論など、Generics を理解するために必要な概念が説明されている
- [When To Use Generics](https://go.dev/blog/when-generics) The Go Blog 2022-04-12
- [Featherweight Go](https://arxiv.org/abs/2005.11710) arxiv 2020-05-24
- [Featherweight Go](https://www.youtube.com/watch?v=62xlcsJ0AUs) YouTube 2020-11-12
- [Goodbye core types - Hello Go as we know and love it!](https://go.dev/blog/coretypes)
- [Generics implementation - Dictionaries](https://github.com/golang/proposal/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/design/generics-implementation-dictionaries.md)
- [Go 1.21 is released!](https://go.dev/blog/go1.21)
- [Go 1.24 Release Notes](https://go.dev/doc/go1.24)

#### プロポーザル
- [Go 1.18 Implementation of Generics via Dictionaries and Gcshape Stenciling](https://github.com/golang/proposal/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/design/generics-implementation-dictionaries-go1.18.md)
  - 初期の設計ドキュメント：
    - [Generics implementation - GC Shape Stenciling](https://github.com/golang/proposal/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/design/generics-implementation-gcshape.md)
    - [Generics implementation - Stenciling](https://go.googlesource.com/proposal/+/refs/heads/master/design/generics-implementation-stenciling.md)
    - [Generics implementation - Dictionaries](https://go.googlesource.com/proposal/+/refs/heads/master/design/generics-implementation-dictionaries.md)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) 2021-08-20
- [Generics implementation - Stenciling](https://github.com/golang/proposal/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/design/generics-implementation-stenciling.md)
- [proposal: spec: allow type parameters in methods](https://github.com/golang/go/issues/49085)
  - メソッドでの型パラメータ使用について
- [proposal: Go 2: function overloading](https://github.com/golang/go/issues/21659)
  - 関数オーバーロードについて

#### 実装
- [src/runtime/iface.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/runtime/iface.go)
  - Go ランタイムの中で interface の内部表現や操作を扱う実装 がまとめられているファイル
  - iface（具体的 interface）、eface（空 interface = interface{}）、itab（型情報テーブル、メソッド解決用）関係の関数や、Typed-nil に関係する convXXX 系メソッドがある
- [src/runtime/runtime2.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/runtime/runtime2.go)
  - [iface構造体の定義](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/runtime/runtime2.go#L178-L181)
  - [eface構造体の定義](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/runtime/runtime2.go#L183-L186)
- [cmd/compile/internal/types2](https://github.com/golang/go/tree/master/src/cmd/compile/internal/types2)
  - 型パラメータ、制約 (type sets) の解析とチェック
  - [cmd/compile/internal/types2/instanciate.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/cmd/compile/internal/types2/instantiate.go)で型引数を元に Named 型（定義型）、Alias 型（型エイリアス）、Signature 型（関数型）を具象化
  - [cmd/compile/internal/types2/named.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/cmd/compile/internal/types2/named.go)
- [cmd/compile/internal/noder](https://github.com/golang/go/tree/master/src/cmd/compile/internal/noder)
- [cmd/compile/internal/ir](https://github.com/golang/go/tree/master/src/cmd/compile/internal/ir)
  - AST におけるジェネリック呼び出し／型の具象化
- [cmd/compile/internal/ssa](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/cmd/compile/internal/ssa/README.md)
  - Generics関数本体の静的単一代入 (SSA) 形式の中間表現 (IR) を使用して、Go プログラムの要素 (パッケージ、型、関数、変数、定数) の表現を定義
- [src/go/types/instantiate.go](https://github.com/golang/go/blob/34e67623a81e9e93c3d7d0f0cb257d7d722939f2/src/go/types/instantiate.go)
  - 型推論・制約チェック

#### イシュー
- [proposal: spec: permit referring to a field shared by all elements of a type set #48522](https://github.com/golang/go/issues/48522)
- [cmd/compile: type parameter involving constraint with channels seems like it should be inferrable #69153](https://github.com/golang/go/issues/69153)

#### その他
- [src/runtime](https://pkg.go.dev/runtime)
- [src/cmd/compile](https://pkg.go.dev/cmd/compile)
- [src/internal/abi](https://pkg.go.dev/internal/abi)
- [Go Developer Survey 2024 H1 Results](https://go.dev/blog/survey2024-h1-results)
- [アプリケーションバイナリインターフェース（Wikipedia）](https://ja.wikipedia.org/wiki/%E3%82%A2%E3%83%97%E3%83%AA%E3%82%B1%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%90%E3%82%A4%E3%83%8A%E3%83%AA%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9)
- [静的単一代入 (Wikipedia)](https://ja.wikipedia.org/wiki/%E9%9D%99%E7%9A%84%E5%8D%98%E4%B8%80%E4%BB%A3%E5%85%A5)
- [Dynamic dispatch (動的ディスパッチ)（Wikipedia）](https://en.wikipedia.org/wiki/Dynamic_dispatch)

