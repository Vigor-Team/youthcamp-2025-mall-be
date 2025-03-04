make:
	kind load docker-image prom/prometheus:latest --name mall-dev
	kubectl apply -f promethues.yaml

delete:
	kubectl delete -f promethues.yaml