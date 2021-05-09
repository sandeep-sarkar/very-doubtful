FROM grpc/go

WORKDIR /usr/src/app
COPY statisticsProcessing /usr/src/app/statisticsProcessing

RUN mkdir -p /usr/src/app/result &&\
    chmod +x /usr/src/app/statisticsProcessing

EXPOSE 50061 50071

CMD ["/usr/src/app/statisticsProcessing"]
