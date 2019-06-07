# 前言
利用plantuml及go-package-plantuml自动生成go程序的类图,主要包括plantuml的安装及使用.

# 步骤
## 安装go-package-plantuml
```sh
mkdir -p $GOPATH/src/git.oschina.net/jscode && cd $GOPATH/src/git.oschina.net/jscode
git clone https://gitee.com/jscode/go-package-plantuml.git
go build git.oschina.net/jscode/go-package-plantuml

# 将二进制可执行文件移到$GOBIN中
mv ./go-package-plantuml ../bin
```

## 使用go-package-plantuml生成uml类图描述文件
```sh
go-package-plantuml --codedir /appdev/gopath/src/github.com/contiv/netplugin \
--gopath /appdev/gopath \
--outputfile  /tmp/result \
--ignoredir /appdev/gopath/src/github.com/contiv/netplugin/vendor

# 参数说明
#--codedir 要分析的代码目录
#--gopath GOPATH环境变量目录
#--outputfile 分析结果保存到该文件
#--ignoredir 不需要进行代码分析的目录（可以不用设置）
```

## 使用plantuml生成类图
### 安装依赖
- [plantuml.jar](https://nchc.dl.sourceforge.net/project/plantuml/plantuml.jar)
- [graphviz](https://graphviz.gitlab.io/download/): `brew install graphviz`

```sh
java -jar plantuml.jar /tmp/uml.txt
# 使用-tsvg参数可以生成svg格式的文件，使用浏览器打开，可以搜索类名。
java -jar plantuml.jar /tmp/uml.txt -tsvg
```

# 参考资料
- [class-diagram](http://plantuml.com/class-diagram)

# TODO
- [ ] 添加整个包级别的分析,实例的使用
- [ ] 降低使用的复杂度,特别是文件名的指定部分
- [ ] 完善整个工具的依赖工作流
