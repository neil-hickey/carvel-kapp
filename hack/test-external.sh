#!/bin/bash

set -e -x -u

time kapp app-group deploy -y -g gitops -d examples/gitops/
time kapp app-group delete -y -g gitops

time kapp deploy -y -a istio -f examples/istio-v1.12.2/
time kapp delete -y -a istio

time kapp deploy -y -a cert-manager -f examples/cert-manager-v1.6.1/
time kapp delete -y -a cert-manager

time kapp deploy -y -a knative -f examples/knative-v1.1.0/
time kapp delete -y -a knative

time kapp deploy -y -a cf -f examples/cf-for-k8s-v0.2.0-custom/
time kapp delete -y -a cf

time kapp deploy -y -a gk -f examples/gatekeeper-v3.7.0/config.yml
time kapp delete -y -a gk

time kapp deploy -y -a pinniped -f examples/pinniped-v0.13.0/
time kapp delete -y -a pinniped

echo EXTERNAL SUCCESS
