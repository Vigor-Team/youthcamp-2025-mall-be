make:
	kind load docker-image grafana/grafana:latest --name mall-dev
	kubectl apply -f grafana.yaml
	kubectl apply -f pv.yaml
	kubectl apply -f pvc.yaml
