.PHONY: build-app push-app docker-all build-both ratelimit-local \
	ratelimit-local ratelimit-internet ratelimit-remote



build-app:
	docker build --platform linux/amd64 -t 214434260197.dkr.ecr.eu-west-1.amazonaws.com/portfolio:latest .

push-app:
	aws sts get-caller-identity
	aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 214434260197.dkr.ecr.eu-west-1.amazonaws.com
	docker push 214434260197.dkr.ecr.eu-west-1.amazonaws.com/portfolio:latest

docker-all: build-app push-app

deploy-remote:
	ssh portfolio-domain "rm -rf ~/portfolio &> /dev/null && git clone https://github.com/leouduh/portfolio.git --depth 1  && cd ~/portfolio && git fetch && git pull && /home/ubuntu/.local/bin/aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 214434260197.dkr.ecr.eu-west-1.amazonaws.com && docker compose pull && docker compose up -d"

ratelimit-local:
	for i in $$(seq 1 15); do curl -s -o /dev/null -w "%{http_code}\n" http://localhost:8080; done

ratelimit-internet:
	for i in $$(seq 1 15); do curl -s -o /dev/null -w "%{http_code}\n" https://portfolio.leo-sama.com; done

ratelimit-remote:
	ssh portfolio-domain 'for i in $$(seq 1 15); do curl -s -o /dev/null -w "%{http_code}\n" https://portfolio.leo-sama.com; done'
