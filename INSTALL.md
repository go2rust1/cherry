# linux平台安装db2驱动
```shell
# 1、下载odbc驱动
wget https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz

# 2、解压到指定目录
tar -xvzf linuxx64_odbc_cli.tar.gz -C /data/driver/db2

# 3、添加环境变量到/etc/profile
export IBM_DB_HOME=/data/driver/db2/clidriver
export CGO_CFLAGS=-I$IBM_DB_HOME/include
export CGO_LDFLAGS=-L$IBM_DB_HOME/lib
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$IBM_DB_HOME/lib

# 4、运行程序(两种方式2选1)
## 4.1、直接运行
go run main.go
## 4.2、编译运行
go build main.go && ./main
```

# macos平台安装db2驱动
```shell
# 1、下载odbc驱动
wget https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/macos64_odbc_cli.tar.gz

# 2、解压到指定目录
tar -xvzf macos64_odbc_cli.tar.gz -C /Users/rustman/Documents/driver/db2/macos

# 3、添加环境变量到~/.zshrc
export IBM_DB_HOME=/Users/rustman/Documents/driver/db2/macos/clidriver
export CGO_CFLAGS=-I$IBM_DB_HOME/include
export CGO_LDFLAGS=-L$IBM_DB_HOME/lib
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:$IBM_DB_HOME/lib

# 4、运行程序(两种方式2选1)
## 4.1、直接运行
go run -exec "env DYLD_LIBRARY_PATH=${IBM_DB_HOME}/lib" main.go
## 4.2、编译运行
go build main.go && ./main
```

# linux平台安装oracle驱动
```shell
# 1、下载驱动: https://www.oracle.com/database/technologies/instant-client/linux-x86-64-downloads.html
https://download.oracle.com/otn/linux/instantclient/121020/instantclient-basic-linux.x64-12.1.0.2.0.zip
https://download.oracle.com/otn/linux/instantclient/121020/instantclient-sdk-linux.x64-12.1.0.2.0.zip
https://download.oracle.com/otn/linux/instantclient/121020/instantclient-sqlplus-linux.x64-12.1.0.2.0.zip

# 2、创建安装目录
mkdir -p /data/driver/oracle
cd /data/driver/oracle

# 3、移动压缩包到安装目录
mv instantclient-basic-linux.x64-12.1.0.2.0.zip ./
mv instantclient-sdk-linux.x64-12.1.0.2.0.zip ./
mv instantclient-sqlplus-linux.x64-12.1.0.2.0.zip ./

# 4、按顺序解压压缩包
unzip instantclient-basic-linux.x64-12.1.0.2.0.zip
unzip instantclient-sdk-linux.x64-12.1.0.2.0.zip
unzip instantclient-sqlplus-linux.x64-12.1.0.2.0.zip

# 5、解压第一个之后会自动出现一个instantclient_12_1的文件夹, 后面解压的两个文件也会自动放到这个文件夹

# 6、创建软链接
cd instantclient_12_1
ln -s libclntsh.so.12.1 libclntsh.so
ln -s libocci.so.12.1 libocci.so
```

# macos平台安装oracle驱动
```shell
# 1、下载驱动: https://www.oracle.com/id/database/technologies/instant-client/macos-intel-x86-downloads.html
https://download.oracle.com/otn/mac/instantclient/121020/instantclient-basic-macos.x64-12.1.0.2.0.zip
https://download.oracle.com/otn/mac/instantclient/121020/instantclient-sdk-macos.x64-12.1.0.2.0.zip
https://download.oracle.com/otn/mac/instantclient/121020/instantclient-sqlplus-macos.x64-12.1.0.2.0.zip

# 2、创建安装目录
mkdir -p /Users/lujiawei/Documents/driver/oracle/macos
cd /Users/lujiawei/Documents/driver/oracle/macos

# 3、移动压缩包到安装目录
mv ~/Downloads/instantclient-basic-macos.x64-12.1.0.2.0.zip ./
mv ~/Downloads/instantclient-sdk-macos.x64-12.1.0.2.0.zip ./
mv ~/Downloads/instantclient-sqlplus-macos.x64-12.1.0.2.0.zip ./

# 4、按顺序解压压缩包
unzip instantclient-basic-macos.x64-12.1.0.2.0.zip
unzip instantclient-sdk-macos.x64-12.1.0.2.0.zip
unzip instantclient-sqlplus-macos.x64-12.1.0.2.0.zip

# 5、解压第一个之后会自动出现一个instantclient_12_1的文件夹, 后面解压的两个文件也会自动放到这个文件夹

# 6、创建软链接
cd instantclient_12_1
ln -s libclntsh.dylib.12.1 /usr/local/bin/libclntsh.dylib.12.1
```
