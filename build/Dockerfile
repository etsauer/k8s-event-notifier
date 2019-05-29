FROM centos:7

ENV OPERATOR=/usr/local/bin/k8s-notify \
    USER_UID=1001 \
    USER_NAME=k8s-notify

# install operator binary
COPY build/_output/bin/k8s-notify ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
