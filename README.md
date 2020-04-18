# aligner

Aligner is a perfect multi-line symbol column alignment tool, 

cross-platform cli

## Feature
* TODO 过滤
* 开启代码段控制 aligner disable/enable
* 文件修改监控格式化。后台进程守护。
* 与vscode、idea 集成
* =、+、-、/ 等运算符号对齐
* 自定义符号对齐
* 第一行开启不控制
* 自适应语言代码对齐格式化
* 单行注释与左边的代码间距设置四个空格


## How to get

```
go get github.com/suyanlong/aligner
```

## Usage

install to PATH env
```shell script
aligner -h

NAME:
   aligner - A new CI/CD cli, align multiple lines of single comments or symbols

USAGE:
   aligner [global options] command [command options] [arguments...]

VERSION:
    aligner 3e7cdaa darwin amd64 Sat Apr 18 02:58:46 UTC 2020


COMMANDS:
   comment  align multiple lines of single comments
   check    aligner check whether the symbols are aligned
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --comment value, -c value  annotation symbols
   --path value, -p value     alignment path (default: "/Users/suyanlong/github/aligner")
   --replace, -r              replace file (default: false)
   --ext value, -e value      file extension
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)

```

## Example

### aligner before python.py file
python.py:
```python
class GosTxType:
    gosMallPay = "-10"  # XXXX
    appPay = "-9"  # XXXX
    ecPay = "-8"  # XXXXXXXXX
    addPool = "-7"  # XXXXXXXXXX
    cityFee = "-6"    # XXXXXXXX
    taskFee = "-5"   # XXXXXXX
 
```

### aligner after python.py file
```shell script
aligner  -c "#"  -e ".py" ./python.py
```
python.py:
```python
class GosTxType:
    gosMallPay = "-10"  # XXXX
    appPay = "-9"       # XXXX
    ecPay = "-8"        # XXXXXXXXX
    addPool = "-7"      # XXXXXXXXXX
    cityFee = "-6"      # XXXXXXXX
    taskFee = "-5"      # XXXXXXX
```

## Contribution Welcomed !

Contributors

* [suyanlong](https://github.com/suyanlong) 


Report issue or pull request, or email 592938410@qq.com 
