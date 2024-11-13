FROM migrate/migrate

# Instalar mysql-client usando apk
RUN apk update && apk add mysql-client