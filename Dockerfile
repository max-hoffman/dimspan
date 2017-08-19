FROM iron/base
WORKDIR /app
# copy binary into image
COPY myapp /app/
ENTRYPOINT ["./myapp"]

EXPOSE 8080