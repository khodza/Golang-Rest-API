# Application Setup Instructions

To start the application, please follow the steps below:

1. Make sure you have Docker and Go installed on your system.

2. Clone the repository:
   https://github.com/khodza/Golang-Rest-API

3. Change to the project directory:
   cd Golang-Rest-API

4. Run the following commands in the terminal:
   make compose-up
   make migrate-up
   make run

Alternatively, you can run the application using Go directly:
go run cmd/main.go

The above commands will set up the necessary dependencies, run any required migrations, and start the application.

5. Open your web browser and visit `http://localhost:8080` to access the application.

Please note that if you encounter any issues during the setup process, make sure to check the project documentation or seek assistance from the project maintainers.

API docs: https://documenter.getpostman.com/view/22439880/2s93z6ejPQ
