FROM iron/base
WORKDIR /app
# copy binary into image
COPY dimspan /app/
COPY plots /app/
ENTRYPOINT ["./dimspan"]

EXPOSE 8080