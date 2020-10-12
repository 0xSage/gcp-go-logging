# [START getting_started_gce_startup_script]
set -ex

# Install logging monitor. The monitor will automatically pickup logs sent to syslog.
curl -s "https://storage.googleapis.com/signals-agents/logging/google-fluentd-install.sh" | bash
service google-fluentd restart &

APP_LOCATION=$(curl -s "http://metadata.google.internal/computeMetadata/v1/instance/attributes/app-location" -H "Metadata-Flavor: Google")
gsutil cp "$APP_LOCATION" app.tar.gz
tar -xzf app.tar.gz

# Start the service included in app.tar.gz.
service my-app start
# [END getting_started_gce_startup_script]