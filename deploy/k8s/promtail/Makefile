make:
	kind load docker-image grafana/promtail:2.9.2 --name mall-dev
	kubectl apply -f promtail.yaml
	kubectl apply -f config.yaml

delete:
	kubectl delete -f promtail.yaml
	kubectl delete -f config.yaml