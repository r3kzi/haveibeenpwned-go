FROM       alpine:3.9
RUN        apk add --update --no-cache ca-certificates
COPY       have-i-been-pwned /bin/have-i-been-pwned
EXPOSE     443
ENTRYPOINT [ "/bin/have-i-been-pwned" ]