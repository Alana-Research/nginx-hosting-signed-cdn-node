# Alana Network nginx-hosting-signed-cdn-node

This is a multi service node configured in docker compose too allow any Alana Network user to run a node to share a webpage or a CDN with pre-signed key access.

## Node architecture

//TODO: diagram...

## How to run 

- Docker Compose: 

  ```sh
  #1. Create a local TLS certificate (will be used the network CA when available):
  openssl req -x509 -nodes -days 365 -subj  "/C=XX/ST=XX/O=Alana Inc/CN=*.alananetwork.private" -newkey rsa:2048 -keyout alananetwork.private.key -out alananetwork.private.crt
  #2. Modify the nginx files as you need.
  #3. Create a .env file in the root folder and add yoour signing key if you need it: KEY="xxxx"
  #3. Deploy locally using docker compose:
  make run
  ```
