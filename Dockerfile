FROM scratch

COPY gosrv /

EXPOSE 7777

ENTRYPOINT ["/gosrv"]
