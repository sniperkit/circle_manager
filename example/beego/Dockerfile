FROM library/golang

# Godep for vendoring
RUN go get -u github.com/tools/godep && \
 go get -u github.com/alecthomas/template && \
 go get -u github.com/go-task/task && \
 go get -u github.com/jirfag/go-queryset/cmd/goqueryset && \
 go get -u github.com/beego/bee && \
 go get -u github.com/jungju/circle_manager && \
 go get -u github.com/jungju/circle_manager/... 

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/jungju/circle_manager/example/beegoapp
RUN mkdir -p $APP_DIR

# Set the entrypoint
# ENTRYPOINT (cd $APP_DIR && ./circle)
ENTRYPOINT (cd $APP_DIR && ./run.sh)
ADD . $APP_DIR
RUN chmod +x $APP_DIR/run.sh

# Compile the binary and statically link
# RUN cd $APP_DIR && CGO_ENABLED=0 godep go build -tags 'bindatafs' -ldflags '-d -w -s'

EXPOSE 8080