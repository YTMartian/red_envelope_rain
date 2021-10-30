FROM centos:7
COPY ./bin/server /root/server
EXPOSE 8080
CMD /root/server