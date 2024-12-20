# ASAP

ASAP is an IoT-based farm management system developed as part of the Carriot internship bootcamp. Designed to improve farm production, ASAP leverages the CQRS (Command Query Responsibility Segregation) architectural pattern and Domain-Driven Design to provide a robust and scalable solution for monitoring and managing farm environments.

## Features

- **Temperature and Humidity Monitoring**
  - Continuously track and display farm temperature and humidity levels to ensure optimal growing conditions.
  
- **Real-Time Data Consumption**
  - Messages from sensors are consumed via MQTT, enabling reliable and efficient data transmission from IoT devices.

## Technology Stack

- **Backend:**
  - **Golang:** Utilized for building high-performance and scalable server-side applications.
  - **gRPC:** Implemented for efficient communication between services.
  - **CQRS Pattern:** Applied to separate read and write operations, enhancing scalability and maintainability.

- **Frontend:**
  - **React.js:** Developed a responsive and user-friendly interface for real-time monitoring and management.

- **Containerization:**
  - **Docker:** Fully Dockerized the application to ensure consistent deployment across different environments.
