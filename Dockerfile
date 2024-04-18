
FROM wordpress:latest


COPY static-pro.zip wp-content/plugins/static-pro.zip

#RUN apt-get update && apt-get install -y unzip
#COPY static-pro.zip /var/www/html/wp-content/plugins/static-pro.zip
#RUN unzip /var/www/html/wp-content/plugins/static-pro.zip -d /var/www/html/wp-content/plugins/ && \
    #rm /var/www/html/wp-content/plugins/static-pro.zip

#COPY static-pro.zip /wp-content/plugins/static-pro.zip
#COPY entrypoint.sh /entrypoint.sh
#RUN chmod +x /entrypoint.sh

ENV WORDPRESS_DB_HOST "example.internal."
ENV WORDPRESS_DB_USER "user"
ENV WORDPRESS_DB_PASSWORD "password"
ENV WORDPRESS_DB_NAME "my-database"

#ENTRYPOINT ["/entrypoint.sh"]
CMD ["apache2-foreground"]

EXPOSE 80
