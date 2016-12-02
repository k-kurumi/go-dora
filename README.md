# go-dora
doraで使いそうなものをgolangで作成

### usage

CFでの起動

```
$ glide install
$ cf push -f manifest.yml
```

ローカルでの起動

```
$ glide install
$ export PORT=3333
$ go run main.go
```
