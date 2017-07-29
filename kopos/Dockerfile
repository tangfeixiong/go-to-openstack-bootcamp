FROM busybox
LABEL "maintainer"="tangfeixiong <tangfx128@gmail.com>" \
    "project"="https://github.com/tangfeixiong/go-to-openstack-bootcamp" \
    "name"="kopos" \
    "version"="0.1"

ARG os_password

COPY bin/kopos /

# RUN mkdir -p /Users/fanhongling/Downloads/tmp

ENV OS_TENANT_NAME='admin' \
  OS_USERNAME='admin' \
  OS_PASSWORD=${os_password:-'secret'} \
  OS_AUTH_URL='http://10.121.198.2:5000/v2.0/' \
  OS_AUTH_STRATEGY='keystone' \
  OS_REGION_NAME='openstack' \
  HV_USERNAME='fake user' \
  HV_PASSWORD='fake secret' \
  HV_KEYPATH='/Users/fanhongling/Downloads/tmp/vagrant' \
  HV_AGENT='false' \
  HV_GRPC='false' \
  MOCK_HOST='127.0.0.1'
  
# VOLUME ["/Users/fanhongling/Downloads/tmp"]
EXPOSE 10000 10001

ENTRYPOINT ["/kopos", "serve"]
CMD ["-k", "/Users/fanhongling/Downloads/tmp/vagrant", "--loglevel=2"]
