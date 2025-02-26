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


app:
	#kind load docker-image product:v1.1.6 gateway:v1.1.6 order:v1.1.6 user:v1.1.6 order:v1.1.6 payment:v1.1.6 email:v1.1.6 checkout:v1.1.6 cart:v1.1.6 llm:v1.1.6 --name mall-dev
	kubectl apply -f app.yaml
	kubectl apply -f gateway.yaml