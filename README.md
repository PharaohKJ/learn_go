Install
=======

[goenv(とgoof)で仮想環境を作ってみた #golang - Qiita](http://qiita.com/knsh14/items/e977b2dbf0efabfc77e1)



1.' OSX に goof を導入
---------------------

※入れるのに時間かかったのでパス。下の1.へ進んでください。

[Go 1.4 の環境構築 Homebrew + Vim 編 (2014.12) - Qiita](http://qiita.com/methane/items/4905f40e4772afec3e60)

```shell-session
brew install go
mkdir ~/go1.4
```

```.bash_profile
export GOPATH=$HOME/go1.4
export GOROOT=/usr/local/opt/go/libexec
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export GOENVHOME=$HOME/.goenvs
source `which goenvwrapper.sh`
```

※上記手順ページを見たが、もうGOROOTは不必要らしい [あなたがGOROOTを本当に設定しなくていい理由 | Androg](http://kwmt27.net/index.php/2013/06/14/you-dont-need-to-set-goroot-really/)


1. OSX に gvm を導入
---------------------

rubyでいうrbenv的なもの。

[最新versionのgolangをぶち込む方法について - Qiita](http://qiita.com/yamadagenki/items/5032cf853f6b136b533f)

```
~/tmp
[pharaohkj]$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
Cloning from https://github.com/moovweb/gvm.git to /Users/pharaohkj/.gvm
Created profile for existing install of Go at /usr/local/opt/go/libexec
Installed GVM v1.0.22

Please restart your terminal session or to get started right away run
   `source /Users/pharaohkj/.gvm/scripts/gvm`
```

```
gvm listall
```

でインストールできるものを確認。今回は1.4.2を選択。(brew install goでも2015-06-10現在、1.4.2が入る)

```
gvm install go1.4.2
gvm use go1.4.2
```

2. emacs の設定
--------------------

[macにgolang入れたりする - Argmax.jp](http://argmax.jp/index.php?mac%E3%81%ABgolang%E5%85%A5%E3%82%8C%E3%81%9F%E3%82%8A%E3%81%99%E3%82%8B)
を参照。

```shell-session
bew install mercurial
go get code.google.com/p/rog-go/exp/cmd/godef
go get -u github.com/nsf/gocode

```

```init.el
(add-hook 'go-mode-hook
          '(lambda()
            (setq c-basic-offset 4)
            (setq indent-tabs-mode t)
            ))
(add-hook 'before-save-hook 'gofmt-before-save)
(add-to-list 'load-path "~/go1.4/src/github.com/nsf/gocode/emacs")
(require 'go-autocomplete)
(require 'auto-complete-config)
```



Server を作ってみる
-------------------

[Go言語でhttpサーバーを立ち上げてHello Worldをする - Qiita](http://qiita.com/taizo/items/bf1ec35a65ad5f608d45)


client を作ってみる
-------------------

[Go net/httpパッケージの概要とHTTPクライアント実装例 - Qiita](http://qiita.com/jpshadowapps/items/463b2623209479adcd88)
[Go言語でPHPのvar_dumpやJavaScriptのconsole.log的なもの - Qiita](http://qiita.com/suin/items/d952fb963956ac31b243)

```
fmt.Printf("%v", value)
fmt.Printf("%+v", value)
fmt.Printf("%#v", value)
fmt.Printf("%T", value)
```

ディレクトリ構成について
------------------------

パッケージ名など、ディレクトリ構成ととても関連するらしいので、ワークスペースの構成はキッチリ配置すること。

[Goコードの書き方 - The Go Programming Language](http://golang-jp.org/doc/code.html#Workspaces)

を参考にして作成すること。

`export GOPATH=$HOME/go` で Go のパスではなく、ワークスペースのトップディレクトリを指定すること。

環境変数を途中で変更すると、上記emacsのflycheckで「ライブラリが見つからない」系のエラーになる
ので注意。

最終的な配置
----------

`export GOPATH=/Users/pharaohkj/gitwork/learn_go`
環境において、

```
~/gitwork/learn_go
[pharaohkj]$ tree
.
├── bin
│   ├── client
│   ├── hello
│   └── server
├── pkg
│   └── darwin_amd64
│       └── github.com
│           └── pharaohkj
│               └── newmath.a
└── src
    └── github.com
        └── pharaohkj
            ├── client
            │   └── client.go
            ├── hello
            │   └── hello.go
            ├── newmath
            │   └── sqrt.go
            └── server
                └── server.go
```

以下は **コードの書き方** ページのサンプル

```
bin/
    streak                         # コマンド実行形式ファイル
    todo                           # コマンド実行形式ファイル
pkg/
    linux_amd64/
        code.google.com/p/goauth2/
            oauth.a                # パッケージオブジェクト
        github.com/nf/todo/
            task.a                 # パッケージオブジェクト
src/
    code.google.com/p/goauth2/
        .hg/                       # mercurialレポジトリメタデータ
        oauth/
            oauth.go               # パッケージソース
            oauth_test.go          # テストソース
    github.com/nf/
        streak/
            .git/                  # gitレポジトリメタデータ
            oauth.go               # コマンドソース
            streak.go              # コマンドソース
        todo/
            .git/                  # gitレポジトリメタデータ
            task/
                task.go            # パッケージソース
            todo.go                # コマンドソース
```

build & install
---------------

前述の最終的な配置において、たとえば `client` をビルドする場合には
`~/gitwork/learn_go` のディレクトリで

`go build github.com/pharaohkj/client`

`go install github.com/pharaohkj/client`

でビルドして`bin`に実行ファイル生成である。
`newmath` の `sqrt.go` を参照した `hello` を `go install github.com/pharaohkj/hello`
でビルドすると、参照されてコンパイルされ、pkgディレクトリに配置されるようだ。

まとめると

```
~/gitwork/learn_go
[pharaohkj]$ go install github.com/pharaohkj/hello
~/gitwork/learn_go
[pharaohkj]$ go install github.com/pharaohkj/server
g~/gitwork/learn_go
~/gitwork/learn_go
[pharaohkj]$ go install github.com/pharaohkj/client
```

```
~/gitwork/learn_go
[pharaohkj]$ bin/hello
Hello, world.  Sqrt(2) = 1.414213562373095
~/gitwork/learn_go
[pharaohkj]$ bin/server &
[1] 26738
~/gitwork/learn_go
[pharaohkj]$ bin/client | jq .
{
  "Message": "Hello, World",
  "Time": "2015-06-10T17:13:59+09:00"
}
~/gitwork/learn_go
```

ほかコマンド
----------

[go - The Go Programming Language](http://golang-jp.org/cmd/go/)

テスト
=====

[Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)


```
$ cp sqrt.go sqrt_test.go
```

して、`func Sqrt` を `func SqrtTest` にする。

ファイル命名ルールは **対象ファイル名_test.go**

以下のようにtestingをimport

```
import (
    "testing"
)
```

関数命名ルールは **func Test対象関数名** (ファイル名はお尻だけれどこちらは逆)
引数は必ず *testing.T を引数に取る

テスト実行
--------

```
~/gitwork/learn_go
[pharaohkj]$ go test github.com/pharaohkj/newmath
--- FAIL: TestSqrt (0.00s)
	sqrt_test.go:12: error! 1.7320508075688774
FAIL
FAIL	github.com/pharaohkj/newmath	0.009s
~/gitwork/learn_go
[pharaohkj]$
```
