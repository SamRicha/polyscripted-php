FROM ubuntu

COPY build-php.sh /php/
WORKDIR /php
ENTRYPOINT ["/bin/bash", "/php/build-php.sh"]
