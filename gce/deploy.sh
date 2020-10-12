set -ex

DEPLOY_DIR=gs://go-logging
DEPLOY_FILENAME=app.tar.gz

if [ -z "$DEPLOY_DIR" ]; then
  echo "Must set \$DEPLOY_DIR. For example: DEPLOY_DIR=gs://my-bucket"
  exit 1
fi

if [ -z "$DEPLOY_FILENAME" ]; then
  echo "Must set \$DEPLOY_FILENAME. For example: DEPLOY_FILENAME=app.tar.gz"
  exit 1
fi

gcloud builds submit --substitutions=_DEPLOY_DIR="$DEPLOY_DIR",_DEPLOY_FILENAME="$DEPLOY_FILENAME"

ZONE=us-central1-f

gcloud compute instances create my-app-instance \
    --image-family=debian-10 \
    --image-project=debian-cloud \
    --machine-type=g1-small \
    --scopes userinfo-email,cloud-platform \
    --metadata-from-file startup-script=startup-script.sh \
    --metadata app-location="gs://go-logging/app.tar.gz" \
    --zone us-central1-f \
    --tags http-server

gcloud compute firewall-rules create default-allow-http-80 \
    --allow tcp:80 \
    --source-ranges 0.0.0.0/0 \
    --target-tags http-server \
    --description "Allow port 80 access to http-server"