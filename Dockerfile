FROM scratch
COPY app.run /server
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY nsswitch.conf /etc/nsswitch.conf
COPY docs/swagger.json docs/swagger.json
COPY logs logs
CMD [ "/server" ]