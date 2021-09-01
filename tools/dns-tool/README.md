## Generate client certs
```bash
openssl req \
    -new -nodes -x509 \ 
    -out certs/client.pem \
    -keyout certs/client.key \
    -days 3650 \
    -subj "/C=DE/ST=NRW/L=Earth/O=N26/OU=IT/CN=www.n26.com/emailAddress=cookdj0128@gmail.com"
```
## Send a dns request with dig to the proxy 
```bash
dig +tcp -p 5353 @0.0.0.0 www.google.com
```
## tcpdump to check on the encryption 
### show the original client request on the loopback address
```bash
sudo tcpdump -i lo -X port 5353
```
### show the dns-client tcp-tls request to the dns-resolver 

```bash
sudo tcpdump -X dst port 853
```
