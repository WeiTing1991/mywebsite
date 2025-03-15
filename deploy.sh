# usr/bin/env bash

gcloud auth configure-docker europe-docker.pkg.dev # login to gcloud docker registry
docker buildx build --platform linux/amd64 -t europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest .
gcloud auth login
gcloud config set project mywebproject-420221

docker push europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest
gcloud run deploy myweb --image=europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest \
--region=europe-west1 --allow-unauthenticated --max-instances=3
