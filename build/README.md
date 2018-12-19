# 淘宝SDK生成
基于淘宝ApiMetadata.xml自动生成

默认生成 ../tbapi.go ../tbmodel.go

../tbapi.go 基于 api.tmpl 模板生成
../tbmodel.go 基于 struct.tmpl 模板生成

build/main中有两个可供调用方法

Create(data) 只生成tbapi.go 与tbmodel.go

ScatteredCreate(data) 会为每一个Struct创建一个文件，路径在../tbapi ../tbmodel