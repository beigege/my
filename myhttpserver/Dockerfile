FROM ubuntu
ENV MY_SERVICE_PORT=80
LABEL multi.label1="value1" multi.label2="value2" other="value3"
COPY bin/amd64/myhttpserver /myhttpserver
EXPOSE 80
ENTRYPOINT /myhttpserver