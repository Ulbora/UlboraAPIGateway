FROM ubuntu

#RUN sudo apt-get update
ADD main /main
ADD entrypoint.sh /entrypoint.sh
WORKDIR /

EXPOSE 3011
ENTRYPOINT ["/entrypoint.sh"]

