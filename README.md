# file-lines-cal


用于统计不同文件的总行数，比如统计代码总行数。

编译

```bash
go build -o file-lines-cal main.go
```


运行效果

```bash
find .  -name "*.scala" -or -name "*.java" -or -name "*.go" -or -name "*.py" -or -name "*.js" -or -name "*.php" -or -name "*.html" -or -name "*.css" -or -name "*.less" -or -name "*.sass" -or -name "*.ts" -or -name "*.lua" | grep -v -e '/vendor/' -e '/node_modules/' | file-lines-cal
FileCount: 43233, Total: 5563561
🚩 .java   : 1585781
🚩 .php    : 1493100
🚩 .js     : 1007976
🚩 .css    : 778902
🚩 .html   : 400378
🚩 .py     : 117507
🚩 .go     : 64302
🚩 .less   : 53847
🚩 .scala  : 36674
🚩 .ts     : 24451
🚩 .lua    : 643
```
