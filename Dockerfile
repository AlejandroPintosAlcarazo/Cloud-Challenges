
FROM wordpress:latest

COPY static-pro.zip /wp-content/plugins/static-pro.zip
#COPY entrypoint.sh /entrypoint.sh
#RUN chmod +x /entrypoint.sh

ENV WORDPRESS_DB_HOST "example.internal."
ENV WORDPRESS_DB_USER "user"
ENV WORDPRESS_DB_PASSWORD "password"
ENV WORDPRESS_DB_NAME "my-database"

#ENTRYPOINT ["/entrypoint.sh"]
#CMD ["apache2-foreground"]

EXPOSE 80
