SHELL = /bin/sh -e
CASS_VERSION = 3.11.8
CASS_IMAGE = cassandra:${CASS_VERSION}
CASS_CONFIG_PATH = ${PWD}/db/configs
CASS_USER = cassandra
CASS_SECRET = cassandra

KEYSPACE = development

up: 
	@echo "--> spinning up docker containers"
	${MAKE} copy-base-config
	${MAKE} copy-config-to-cluster
	@docker compose up -d

down: 
	@echo "--> stopping docker containers"
	@docker compose down 

restart: 
	@echo "--> restarting containers"
	@docker compose restart


pause: 
	@echo "--> pausing containers"
	@docker compose pause

unpause: 
	@echo "--> unpausing containers"
	@docker compose start

copy-base-config: 
	@echo "--> pulling configs from cassandra image"
	@rm -rf ${CASS_CONFIG_PATH}/etc_cassandra-${CASS_VERSION}_vanilla
	@if [ $$(docker images --format "{{.Repository}}" | grep -c "${CASS_IMAGE}") -lt 1 ]; then \
		echo "--> found no matching image. pulling."; \
		docker pull ${CASS_IMAGE}; \
	fi
	@docker run --rm -d --name tmp ${CASS_IMAGE}
	@docker cp tmp:/etc/cassandra ./db/configs/etc_cassandra-${CASS_VERSION}_vanilla
	@docker stop tmp

copy-config-to-cluster: 
	@echo "--> copying config into cassandra nodes"
	@rm -rf ./etc
	@mkdir -p db/etc
	@cp -a ${CASS_CONFIG_PATH}/etc_cassandra-${CASS_VERSION}_vanilla db/etc/cass1
	# @cp -a ${CASS_CONFIG_PATH}/etc_cassandra-${CASS_VERSION}_vanilla db/etc/cass2
	# @cp -a ${CASS_CONFIG_PATH}/etc_cassandra-${CASS_VERSION}_vanilla db/etc/cass3

get-keyspaces: 
	@echo "--> getting cassandra keyspaces"
	@docker exec cass1 cqlsh -e "describe keyspaces;"    

init-keyspace: 
	@echo "--> checking for ${KEYSPACE} keyspace"
	@if [ $$(docker  exec cass1 cqlsh -e "describe keyspaces;" | grep -c "${KEYSPACE}") -lt 1 ]; then \
		echo "--> found no matching keyspace"; \
		for i in {1..5}; do $$(docker cp ${PWD}/db/init_keyspace.cql cass1:/ && docker exec cass1 cqlsh -u ${CASS_USER} -p ${CASS_SECRET} -f /init_keyspace.cql) && break || sleep 10; \
		done; \
	fi
	@echo "--> ${KEYSPACE} keyspace ready"
	@docker exec cass1 cqlsh -e "describe keyspace ${KEYSPACE}"
