// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRedisInstance_redisInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRedisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_redis_instance" "cache" {
  name           = "memory-cache-%{random_suffix}"
  memory_size_gb = 1
}
`, context)
}

func TestAccRedisInstance_redisInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRedisInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceFullExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_redis_instance" "cache" {
  name           = "ha-memory-cache-%{random_suffix}"
  tier           = "STANDARD_HA"
  memory_size_gb = 1

  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  authorized_network = "${google_compute_network.auto-network.self_link}"

  redis_version     = "REDIS_3_2"
  display_name      = "Terraform Test Instance"
  reserved_ip_range = "192.168.0.0/29"

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }
}

resource "google_compute_network" "auto-network" {
  name = "authorized-network-%{random_suffix}"
}
`, context)
}

func testAccCheckRedisInstanceDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_redis_instance" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://redis.googleapis.com/v1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("RedisInstance still exists at %s", url)
		}
	}

	return nil
}
