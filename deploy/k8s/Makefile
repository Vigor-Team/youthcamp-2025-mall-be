all:
	#make dashboard
	make -C consul make
	make -C mongodb make
	make -C mysql make
	make -C rabbitmq make
	make -C redis make
	make -C loki make
	make -C es make
	make -C kibana make
	make -C prometheus make
	make -C jaeger make
	make -C nats make
	make -C promtail make
	make -C grafana make



app:
	kind load docker-image product:v1.1.1 gateway:v1.1.1 order:v1.1.1 user:v1.1.1 order:v1.1.1 payment:v1.1.1 email:v1.1.1 checkout:v1.1.1 cart:v1.1.1 llm:v1.1.1 --name mall-dev
	kubectl apply -f app.yaml
	kubectl apply -f gateway.yaml

