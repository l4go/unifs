# fs.FSのUniPath形式文字列でのラッパー関数

fs.FSの機能をUni-pathで楽に呼び出す為のラッパー関数群

## `func Open(fsys fs.FS, name string) (fs.File, error)`

`fsys`の値を使って、Uni-path形式ファイルの[fs.Open()](https://pkg.go.dev/io/fs#FS)を行います。

## `func Sub(fsys fs.FS, name string) (fs.FS, error)`

`fsys`の値を使って、サブディレクトリのfs.FSを生成します。([fs.Sub()](https://pkg.go.dev/io/fs#SubFS)を実行します。)

## `func Stat(fsys fs.FS, name string) (fs.FileInfo, error)`

`fsys`の値を使って、Uni-path形式ファイルの[Stat()](https://pkg.go.dev/io/fs#StatFS)を行います。

## `func Glob(fsys fs.FS, pattern string) ([]string, error)`

`fsys`の値を使って、Uni-path形式ファイルの[Glob()](https://pkg.go.dev/io/fs#GlobFS)を行います。

## `func ReadFile(fsys fs.FS, name string) ([]byte, error)`

`fsys`の値を使って、Uni-path形式ファイルの[ReadFile()](https://pkg.go.dev/io/fs#ReadFileFS)を行います。

## `func ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error)`

`fsys`の値を使って、Uni-path形式ファイルの[ReadDir()](https://pkg.go.dev/io/fs#ReadDirFS)を行います。

