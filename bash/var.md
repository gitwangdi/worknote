# Var

## #\*,##\*,%\*,%%\*

\# 去掉左边  
\% 去掉右边  
单个符号 最小匹配  
两个符号 最大匹配

eg.

```bash
var=aa.bb.cc.dd.ee
${var#*.} # bb.cc.dd.ee
${var##*.} # ee
${var%.*} # aa.bb.cc.dd
${var%%.*} # aa
```

## :

```bash
var=aa.bb.cc.dd.ee
# var:startIndex:length
${var:5:5} # .cc.d
```

## /

/ 替换第一个匹配字符串
// 替换所有匹配字符串

```bash
var=aa.bb.cc.bb.ee
${var/bb/path} # aa.path.cc.bb.ee
${var//bb/path} # aa.path.cc.path.ee
```
