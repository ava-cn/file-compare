# URL文件对比

通过获取给定的URL文件内容判定其MD5值是否一致。

## 使用方法

在命令行执行 `file-compare` 命令，参数跟需要对比的文件地址。

```bash
file-compare https://www.baidu.com https://www.taobao.com
```

## 安装

下载二进制文件 `file-compare` ，并添加到电脑的 `$PATH` 目录下。

[Mac with Intel Chip](https://github.com/ava-cn/file-compare/raw/master/release/macos/amd64/file-compare)
[Mac with Apple Chip](https://github.com/ava-cn/file-compare/raw/master/release/macos/arm64/file-compare)

比如一般会放在 `/usr/local/bin/` 目录下。

```bash
chmod +x file-compare # 添加可执行权限
cp file-compare /usr/local/bin/.
```

