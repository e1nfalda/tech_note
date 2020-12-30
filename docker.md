# docker 
<!-- :Tech:application: -->
### 语法

`FROM base_image`

`RUN  ["可执行文件", "参数1", "参数2"]`

docker run 时运行

`CMD ["<可执行文件或命令>","<param1>","<param2>",...] `

docker build

`WORKDIR dirctory`

`ADD `

- ADD 的优点：在执行 <源文件> 为 tar 压缩文件的话，压缩格式为 gzip, bzip2 以及 xz 的情况下，会自动复制并解压到 <目标路径>。
- ADD 的缺点：在不解压的前提下，无法复制 tar 压缩文件。会令镜像构建缓存失效，从而可能会令镜像构建变得比较缓慢。具体是否使用，可以根据是否需要自动解压来决定。

`COPY [--chown=<user>:<group>] ["<源路径1>",...  "<目标路径>"]`

`ENTRYPOINT ["<executeable>","<param1>","<param2>",...] ` 

类似于 CMD 指令，但其不会被 docker run 的命令行参数指定的指令所覆盖，而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序。

`EXPOSE <端口1> [<端口2>...]`

`VOLUME ["<路径1>", "<路径2>"...]`

定义匿名数据卷。在启动容器时忘记挂载数据卷，会自动挂载到匿名卷。

作用：

- 避免重要的数据，因容器重启而丢失，这是非常致命的。
- 避免容器不断变大。

`ARG <参数名>[=<默认值>]`:docker build 的过程。

`ENV <key1>=<value1> <key2>=<value2>...`

`USER <端口1> [<端口2>...]`

----

[docker file command](https://www.runoob.com/docker/docker-dockerfile.html)
