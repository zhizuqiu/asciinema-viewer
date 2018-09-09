# asciinema-viewer

asciinema文件展示

## build

先将前端项目iview-project生成的`dist目录`和`index_prod.html`复制到`static`和`views`目录下

linux:

```
GOOS=linux GOARCH=amd64 go build .
```

windows:

```
go build .
```

## Development

```
bee run
```

## run

linux:

```
./asciinema-viewer
```

windows:

```
.\asciinema-viewer.exe
```

## config
