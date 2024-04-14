# usr/bin/env bash

docker buildx build --platform linux/amd64 -t europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest .
docker push europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest

gcloud run deploy myweb --image=europe-docker.pkg.dev/mywebproject-420221/mywebsite/myserver:lastest \
--region=europe-west1 --allow-unauthenticated --max-instances=3
