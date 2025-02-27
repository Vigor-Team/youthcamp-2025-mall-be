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



app:
	kind load docker-image product:v1.1.8 gateway:v1.1.8 order:v1.1.8 user:v1.1.8 order:v1.1.8 payment:v1.1.8 email:v1.1.8 checkout:v1.1.8 cart:v1.1.8 llm:v1.1.8 --name mall-dev
	kubectl apply -f app.yaml
	kubectl apply -f gateway.yaml

