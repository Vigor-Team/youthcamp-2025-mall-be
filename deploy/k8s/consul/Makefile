make:
	kind load docker-image hashicorp/consul:latest --name mall-dev
	kubectl apply -f pv.yaml
	kubectl apply -f pvc.yaml
	kubectl apply -f consul.yaml
