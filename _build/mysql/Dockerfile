FROM mysql:latest

#COPY CODE INTO WORKSPACE
COPY ./_build/mysql/demoapp.sql /home/demoapp.sql
COPY ./_build/mysql/dump.sh /docker-entrypoint-initdb.d/dump.sh

RUN chmod -R 775 /docker-entrypoint-initdb.d

EXPOSE 3306