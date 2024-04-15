output "host" {
  value = google_sql_database_instance.storage.public_ip_address
}

output "port" {
  value = 5432
}

output "name" {
  value = google_sql_database.storage.name
}

output "user" {
  value = google_sql_user.new-user.name
}