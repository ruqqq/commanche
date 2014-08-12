#
# Commanche Dockerfile
#

FROM        ubuntu
MAINTAINER  Faruq Rasid "faruq.91@gmail.com"

ADD build/commanche-linux-amd64 /bin/commanche
CMD /bin/commanche --help