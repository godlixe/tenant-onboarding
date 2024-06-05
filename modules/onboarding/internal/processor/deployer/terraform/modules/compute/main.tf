
resource "google_cloud_run_service" "compute" {
  name     = "compute-${var.compute_name}"
  location = "asia-southeast2"

  template {
    spec {
      containers {
        image = "asia-southeast2-docker.pkg.dev/static-booster-418207/saas-todo-list-test/saas-todo-list-test:saas-todo-list-test"
        env {
          name  = "DB_HOST"
          value = var.storage_host
        }
        env {
          name  = "DB_NAME"
          value = var.storage_name
        }
        env {
          name  = "DB_USER"
          value = var.storage_user
        }
        env {
          name  = "DB_PASS"
          value = var.storage_password
        }
        env {
          name  = "DB_PORT"
          value = var.storage_port
        }
      }
    }
  }


  traffic {
    percent         = 100
    latest_revision = true
  }

}

resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.compute.name
  location = google_cloud_run_service.compute.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
