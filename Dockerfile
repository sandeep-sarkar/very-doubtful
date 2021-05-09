FROM grpc/go

WORKDIR /usr/src/app
COPY statisticsProcessing /usr/src/app/statisticsProcessing
COPY health/grpc-health-probe /bin/grpc-health-probe

RUN mkdir -p /usr/src/app/result &&\
    chmod +x /usr/src/app/statisticsProcessing &&\
    chmod +x /bin/grpc-health-probe

EXPOSE 50061 50071

CMD ["/usr/src/app/statisticsProcessing"]
