FROM alpine:latest
MAINTAINER zhongq26@mail2.sysu.edu.cn

RUN apk update \
&& apk upgrade \
&& apk add --no-cache bash bash-doc bash-completion \
&& rm -rf /var/cache/apk/* 

ENV GOROOT=/usr/local/go
WORKDIR /usr/local

RUN apk add --no-cache git
RUN git clone https://github.com/golang/go.git

# build from source code in a nutshell
RUN /bin/bash -c '$GOROOT/src/all.bash'
RUN cd go/
RUN git checkout go1.12.6
# setup env
ENV USERNAME=zhongq
COPY /home/$USERNAME/.bashrc $HOME/
RUN /bin/bash -c 'source $HOME/.bashrc'

# go doc --port=8979
EXPOSE 8979

CMD ["go", "doc", "--port=8979"]
