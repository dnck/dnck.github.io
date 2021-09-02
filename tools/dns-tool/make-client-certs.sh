mkdir ./certs && \
  openssl req \
    -new -nodes -x509 -out certs/client.pem \
    -keyout certs/client.key -days 3650 \
    -subj "/C=DE/ST=NRW/L=Earth/O=N26/OU=IT/CN=www.n26.com/emailAddress=cookdj0128@gmail.com"