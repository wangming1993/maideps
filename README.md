# maideps

go的依赖管理工具， 参考[【godep】](https://github.com/tools/godep)实现

***support: go1.5+***

## Usage

`go get github.com/wangming/maideps`

通过`go get`安装， 将会在你的`$GOPATH/bin`下面生成**maideps**的可执行文件

如果你的`$GOPATH/bin`被添加到了环境变量里面(`$PATH`), 那么就可以全局使用**maideps**

在你的项目根目录运行`maideps`, 即可以分析你的项目中的依赖， 将import的外部packages管理到`vendor`目录
