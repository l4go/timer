# type Timer
タイムアウトを通知するだけのシンプルなモジュールです。  
go標準のtime.NewTimer()とは違い、短いタイムアウト時間でも安全に使えます。

## import
```go
import "github.com/l4go/timer"
```
vendoringして使うことを推奨します。

## 利用サンプル

[example](../examples/ex_timer/ex_timer.go)

## メソッド概略

### func NewTimer() *Timer
*Timerを生成します。go標準のtime.Timerとは違い、生成時はタイマーは止まっていますので、Start()メゾッドを呼ぶまでは、通知されません。

### func (t *Timer) Start(d time.Duration)
指定した時間でタイムアウト処理を開始します。

### func (t *Timer) Stop()
タイムアウト処理を中止します。重複して呼び出しても飲んだいないので、defer文で実行して、無駄なタイムアウト処理を止めることが出来ます。

### func (t *Timer) Recv() <-chan struct{}
タイムアウトを受信します。
