---
name: ergo
director_uuid: %%UUID%%
releases:
- name: ergo
  version: latest
resource_pools:
- name: default
  network: default
  stemcell:
    name: bosh-warden-boshlite-ubuntu-trusty-go_agent
    version: latest
networks:
- name: default
  subnets:
  - cloud_properties:
      name: random
    range: 10.244.2.0/30
    reserved:
    - 10.244.2.1
    static:
    - 10.244.2.2
  - cloud_properties:
      name: random
    range: 10.244.2.4/30
    reserved:
    - 10.244.2.5
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.8/30
    reserved:
    - 10.244.2.9
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.12/30
    reserved:
    - 10.244.2.13
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.16/30
    reserved:
    - 10.244.2.17
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.20/30
    reserved:
    - 10.244.2.21
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.24/30
    reserved:
    - 10.244.2.25
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.28/30
    reserved:
    - 10.244.2.29
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.32/30
    reserved:
    - 10.244.2.33
    static: []
  - cloud_properties:
      name: random
    range: 10.244.2.36/30
    reserved:
    - 10.244.2.37
    static: []
compilation:
  workers: 1
  network: default
  reuse_compilation_vms: true
  cloud_properties: {}
update:
  canaries: 1
  max_in_flight: 3
  canary_watch_time: 15000-30000
  update_watch_time: 15000-300000
jobs:
- name: deploy-service-broker
  instances: 1
  lifecycle: errand
  templates:
  - name: deploy-service-broker
    release: ergo
  resource_pool: default
  networks:
  - name: default
properties: {}
dev_name: ergo
