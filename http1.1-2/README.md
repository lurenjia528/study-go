# HTTP/1.1 HTTP/2

目前市面使用最多的HTTP/1.1

Go的http/2要求必须使用tls,目前浏览器对http/2都是采用TLS的方式

Go使用h2c,提供非tls的http/2

谷歌浏览器f12不显示协议版本,火狐浏览器显示协议版本

```bash
root@HT061:/home/ht061/ygt/11/test/hw#  openssl genrsa -out rootCA.key 2048
Generating RSA private key, 2048 bit long modulus
........+++
...............................+++
e is 65537 (0x010001)
root@HT061:/home/ht061/ygt/11/test/hw#  openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) [Internet Widgits Pty Ltd]:csse
Organizational Unit Name (eg, section) []:csse
Common Name (e.g. server FQDN or YOUR name) []:lurenjia
Email Address []:lurenjia@163.com
root@HT061:/home/ht061/ygt/11/test/hw# openssl genrsa -out server.key 2048^C
root@HT061:/home/ht061/ygt/11/test/hw# ll
总用量 12448
drwxr-xr-x 2 root  root     4096 6月  28 09:41 ./
drwxrwxr-x 4 ht061 ht061    4096 6月  28 09:07 ../
-rw------- 1 root  root     1675 6月  28 09:40 rootCA.key
-rw-r--r-- 1 root  root     1403 6月  28 09:41 rootCA.pem
root@HT061:/home/ht061/ygt/11/test/hw# openssl genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus
.....+++
...............................+++
e is 65537 (0x010001)
root@HT061:/home/ht061/ygt/11/test/hw# openssl req -new -key server.key -out server.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) [Internet Widgits Pty Ltd]:csse
Organizational Unit Name (eg, section) []:csse
Common Name (e.g. server FQDN or YOUR name) []:lurenjia
Email Address []:lurenjia@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:123123
An optional company name []:csse
root@HT061:/home/ht061/ygt/11/test/hw#  openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500
Signature ok
subject=C = CN, ST = Beijing, L = Beijing, O = csse, OU = csse, CN = lurenjia, emailAddress = lurenjia@163.com
Getting CA Private Key
root@HT061:/home/ht061/ygt/11/test/hw# ll
总用量 12464
drwxr-xr-x 2 root  root     4096 6月  28 09:42 ./
drwxrwxr-x 4 ht061 ht061    4096 6月  28 09:07 ../
-rw------- 1 root  root     1675 6月  28 09:40 rootCA.key
-rw-r--r-- 1 root  root     1403 6月  28 09:41 rootCA.pem
-rw-r--r-- 1 root  root       17 6月  28 09:42 rootCA.srl
-rw-r--r-- 1 root  root     1281 6月  28 09:42 server.crt
-rw-r--r-- 1 root  root     1102 6月  28 09:42 server.csr
-rw------- 1 root  root     1679 6月  28 09:41 server.key

```