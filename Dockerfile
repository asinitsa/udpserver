FROM debian

COPY udpserver /bin/udpserver
RUN chmod +x /bin/udpserver
EXPOSE 1234
CMD /bin/udpserver