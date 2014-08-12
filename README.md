COMMANCHE
---------
Utility tool to help with configuring hipache redis configuration without having to deal with redis command. If a frontend doesn't exit, commanche will add it for you. If there is no backend left for the frontend, commanche will remove it for you.

HOW TO USE WITH DOCKER
----------------------
docker run -it --rm --link redis:redis commanche bash -c 'commanche -h $REDIS_PORT_6379_TCP_ADDR -add -f="domain.com"'

COMMANDS
--------
Run with --help.

-add run commanche in add mode
-rm run commanche in remove mode
-h redis host
-p redis port
-f domain name to add/rm backend to/from
-b the backend address separated by ','

TODO
----
- Integration with Cloudflare?