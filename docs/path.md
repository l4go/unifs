# l4go/unifsのパス変換関数

string型のOSパス、fs.FSパス、Uni-Path形式のパスを、相互変換するための関数群です。

## `func MustClean(uni_name string) string`
Uni-path形式のパスを最短化します。  
Uni-path形式ではない場合、panicします。

最短化は、[path.Clean()](https://pkg.go.dev/path#Clean)を利用しています。

エラー無く変換できることが分かっている場合に利用します。

## `func Clean(uni_name string) (string, error)`

Uni-path形式のパスを最短化します。  
Uni-path形式ではない場合、エラーを返します。

最短化は、[path.Clean()](https://pkg.go.dev/path#Clean)を利用しています。

## `func MustJoin(names ...string) string`

パス要素を結合して、Uni-path形式の文字列にします。
Uni-path形式ではない場合、panicします。

パス要素の結合には、[path.Join()](https://pkg.go.dev/path#Join)を利用しています。

エラー無く結合できることが分かっている場合に利用します。

## `func Join(names ...string) (string, error)`

パス要素を結合して、Uni-path形式の文字列にします。
Uni-path形式ではない場合、エラーを返します。

パス要素の結合には、[path.Join()](https://pkg.go.dev/path#Join)を利用しています。

## `func ValidSubPath(uni_name string) bool`

Uni-path形式のサブディレクトリ部分(相対パス)として、正常なものか確認します。

## `func ValidPath(uni_name string) bool`

Uni-path形式のパスとして(絶対パス)として、正常なものか確認します。

## `func ToFSPath(uni_name string) (string, error)`

Uni-path形式のパスを、fs.FS用のパスに変換します。
渡されたパスが不適切の場合、エラーを返します。

## `func ToFSPaths(uni_names []string) ([]string, error)`

Uni-path形式のパスのSliceを、fs.FS用のパスのSliceに変換します。
渡されたパスが不適切の場合、エラーを返します。

## `func ToOSPath(uni_name string) (string, error)`

Uni-path形式のパスを、OS用のパスに変換します。(動作してるOSによって結果が変わります。)
渡されたパスが不適切の場合、エラーを返します。

## `func ToOSPaths(uni_names []string) ([]string, error)`

Uni-path形式のパスのSliceを、OS用のパスのSliceに変換します。(動作してるOSによって結果が変わります。)
渡されたパスが不適切の場合、エラーを返します。

## `func FromFSPath(fs_name string) (string, error)`

fs.FS用のパスを、Uni-path形式のパスに変換します。
渡されたパスが不適切の場合、エラーを返します。

## `func FromOSPath(os_name string) (string, error)`

OSのパスを、Uni-path形式のパスに変換します。
渡されたパスが不適切の場合、エラーを返します。
