FROM busybox

MAINTAINER "cherryleo"

ADD ldappasswd /
ADD index.html /

CMD [ "./ldappasswd" ]

