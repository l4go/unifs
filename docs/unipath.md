# UniPath型
Uni-path形式をカプセル化するUniPath型を提供します。  
stringにUni-path形式のデータを保存すると、コード上で区別しにくい場合に、Uni-path形式の変換漏れ防止に利用することを想定しています。

## `Zero`

ゼロ値のUniPathの値です。代入等の複製元としてコピーして利用します。  
ゼロ値の確認には、`IsZero()`メソッドを利用してください。

`New()`メソッドなどに空文字stringを渡すとエラーが発生するので、
ゼロ値でUniPath型を初期化したいときに利用します。

## `func New(uni_name string) (UniPath, error)`

Uni-path形式の文字列から、UniPath型を作成します。

## `func MustNew(uni_name string) UniPath`

Uni-path形式の文字列から、UniPath型を作成します。
作成に失敗すると、panicします。

確実に作成できるときに利用します。

## `func NewFromOSPath(os_name string) (UniPath, error)`

OS形式のパスから、UniPath形を作成します。

## `func MustNewFromOSPath(os_path string) UniPath`

OS形式のパスから、UniPath型を作成します。
作成に失敗すると、panicします。

## `func NewFromFSPath(fs_name string) (UniPath, error)`

fs.FS用のパスから、UniPath形を作成します。

## `func MustNewFromFSPath(fs_name string) UniPath`

fs.FS用のパスから、UniPath形を作成します。
作成に失敗すると、panicします。

## `func (up UniPath) IsZero() bool`

Zero値のUniPathかどうか判定します。

## `func (up UniPath) String() string`

Uni-path形式の文字列を返します。

## `func (up UniPath) FSPath() string`

fs.FS用のパスを返します。

## `func (up UniPath) OSPath() string`

OS形式のパスを返します。

## `func (up UniPath) Join(names ...string) (UniPath, error)`

渡されたパスを結合したUniPath型の値を生成します。

## `func (up UniPath) MustJoin(names ...string) UniPath`

渡されたパスを結合したUniPath型の値を生成します。
生成に失敗するとpanicします。

## `func (up UniPath) Open(fsys fs.FS) (fs.File, error)`

`fsys`の値を使って、ファイルの[fs.Open()](https://pkg.go.dev/io/fs#FS)を行います。

## `func (up UniPath) Sub(fsys fs.FS) (fs.FS, error)`

`fsys`の値を使って、サブディレクトリのfs.FSを生成します。([fs.Sub()](https://pkg.go.dev/io/fs#SubFS)を実行します。)

## `func (up UniPath) Stat(fsys fs.FS) (fs.FileInfo, error)`

`fsys`の値を使って、ファイルの[Stat()](https://pkg.go.dev/io/fs#StatFS)を行います。

## `func (up UniPath) ReadFile(fsys fs.FS) ([]byte, error)`

`fsys`の値を使って、ファイルの[ReadFile()](https://pkg.go.dev/io/fs#ReadFileFS)を行います。

## `func (up UniPath) ReadDir(fsys fs.FS) ([]fs.DirEntry, error)`

`fsys`の値を使って、ファイルの[ReadDir()](https://pkg.go.dev/io/fs#ReadDirFS)を行います。

## `func (up UniPath) Glob(fsys fs.FS) ([]UniPath, error)`

`fsys`の値を使って、ファイルの[Glob()](https://pkg.go.dev/io/fs#GlobFS)を行います。
