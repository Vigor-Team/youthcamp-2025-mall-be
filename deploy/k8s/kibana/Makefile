make:
	kind load docker-image docker.elastic.co/kibana/kibana:8.17.1 --name mall-dev
	kubectl apply -f kibana.yaml

delete:
	kubectl delete -f kibana.yaml