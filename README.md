
```shell
  go get github.com/mtmuo/go-pipeline
```



```go
f := new(Factory)
defer f.Release()
firstLine = f.NewLine(50, func(data interface{}) {
    secondLine.Submit(data)
    time.Sleep(1 * time.Second)
    atomic.AddInt64(&aa, 1)
})
secondLine = f.NewLine(1000, func(data interface{}) {
    atomic.AddInt64(&aa, 1)
})
for i := 0; i < 10; i++ {
    firstLine.Submit(i)
}
f.Wait()
```