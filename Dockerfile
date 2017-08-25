FROM iron/base
WORKDIR /app
# copy binary into image
COPY dimspan /app/
ENTRYPOINT ["./dimspan"]

EXPOSE 8080