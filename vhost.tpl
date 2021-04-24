upstream dynamic {
    server {{ansible_hostname}}:{{bindport}};
}

server {
    listen 3000;
    location / {
        proxy_pass http://dynamic;
    }
}

