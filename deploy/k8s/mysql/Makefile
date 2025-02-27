make:
	kind load docker-image mysql:latest --name mall-dev
	kubectl apply -f pv.yaml
	kubectl apply -f pvc.yaml
	kubectl apply -f mysql.yaml
	kubectl apply -f secret.yaml

delete:
	kubectl delete -f mysql.yaml
	kubectl delete -f pvc.yaml
	kubectl delete -f pv.yaml
	kubectl delete -f secret.yaml