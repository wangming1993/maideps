# maideps

go的依赖管理工具， 参考[【godep】](https://github.com/tools/godep)实现

***support: go1.5+***

## Usage

`go get github.com/wangming1993/maideps`

通过`go get`安装， 将会在你的`$GOPATH/bin`下面生成**maideps**的可执行文件

如果你的`$GOPATH/bin`被添加到了环境变量里面(`$PATH`), 那么就可以全局使用**maideps**

在你的项目根目录运行`maideps`, 即可以分析你的项目中的依赖， 将import的外部packages管理到`vendor`目录

```
Usage of maideps:
  -debug
        Debug mode, to show more log
  -delete
        To delete dependency
  -import string
        Specific one import package name, only find its dependency
  -reload
        Forced to copy files, ignores file exists

```

通过指定需要import的package名称，可以只管理指定的package依赖

如： `maideps  -debug -import="github.com/garyburd/redigo/redis"`

查找顺序为：

- 是否为golang内置的package, 或是则结束
- 是否为当前项目下的package, 若是，继续递归搜索依赖，但是不加入到依赖中 
- 是否为当前项目vendor下的package, 若是，继续递归搜索依赖，且加入到依赖中 
- 是否为`$GOPATH/src`下的package, 若是，继续递归搜索依赖，且加入到依赖中
- 终止，提示需要收到 `go get packageName` 
