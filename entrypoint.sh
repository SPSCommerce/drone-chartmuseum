#!/bin/sh

export AWS_REGION="us-east-1"
username=$(ssm_get_parameter /techops/cloud-engineering/bdp/dev/chartmuseum_user)
password=$(ssm_get_parameter /techops/cloud-engineering/bdp/dev/chartmuseum_pass)
# may as well hard code the repo url

# invoke plugin
/bin/drone-chartmuseum --repo-url "https://chartmuseum.spsapps.net/" --username $username --password $password --log-level debug
