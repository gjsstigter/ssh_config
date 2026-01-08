host "my-service" {
  hostname = "instance[1..2].my-service.example.com"
  alias    = "myservice{#1}"
  
  # This is an attribute containing an object
  config = {
    user          = "ubuntu"
    identity_file = "~/.ssh/my_service.pem"
    port          = 22
  }
}

host "my-service2" {
  hostname = "instance[1..10].my-service.example.com"
  alias    = "myservice{#1}"
  
  # This is an attribute containing an object
  config = {
    user          = "ubuntu"
    identity_file = "~/.ssh/my_service.pem"
    port          = 22
  }
}