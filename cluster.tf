data "ec_stack" "latest" {
  version_regex = "latest"
  region        = "us-east-1"
}

# Create an Elastic Cloud deployment
resource "ec_deployment" "example_minimal" {
  # Optional name.
  name = "my_example_deployment"

  region                 = "us-east-1"
  version                = data.ec_stack.latest.version
  deployment_template_id = "aws-io-optimized-v2"

  elasticsearch = {
    autoscale = "false"
    hot = {
        autoscaling = {
            max_size = "1g"
        }
        size = "1g"
        zone_count = 1
    }
  }

  kibana = {}

  integrations_server = {}
}