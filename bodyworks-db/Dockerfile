FROM mysql
LABEL maintainer="kevyn24ec@gmail.com"

ENV MYSQL_ROOT_PASSWORD @Oct2019+
COPY Functions /home/mysql/bodyworks/Functions
COPY Procedures /home/mysql/bodyworks/Procedures
COPY Tables /home/mysql/bodyworks/Tables
COPY Views /home/mysql/bodyworks/Views
COPY Scripts /home/mysql/bodyworks/Scripts
workdir /home/mysql/bodyworks/
COPY install.sql /docker-entrypoint-initdb.d/