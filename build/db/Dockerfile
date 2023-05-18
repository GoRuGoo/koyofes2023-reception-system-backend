FROM mysql:8-debian


RUN apt-get update && \
    apt-get install -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8 &&\
    apt update; apt -y upgrade &&\
    apt install vim apt-utils -y

ENV LC_ALL ja_JP.UTF-8
EXPOSE 3306
