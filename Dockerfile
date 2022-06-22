FROM golang:1.10.7

#创建工作目录
RUN mkdir -p /go/src/web_db
#进入工作目录
WORKDIR /go/src/web_db

#将当前目录下的所有文件复制到指定位置
COPY . /go/src/web_db

#下载beego和bee
RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get github.com/go-sql-driver/mysql

#端口
EXPOSE 8080

#运行
CMD ["bee", "run"]