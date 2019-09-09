#!/bin/bash
date1=$(kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l)

while [ $date1 != 23 ]; do
  date1=$(kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l)
  echo "${date1}"
done
