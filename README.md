# Tioncon 🚀 🔥

Companies collect data from thousands to millions of sensors and embedded devices in order to track physical assets such as trucks, for improved efficiency in business operations. This collected data can answer important questions such as travel time, distance coverage, fuel consumption and estimation, average speed, real-time location etc. 

## Description 🏠
A service that collects data from IoT sensors and embedded devices such as trucks, stores it in a scalable data store, provides a dashboard for data visualisation, and an API for integrations with different client applications ie. mobile, desktop, web apps, wearable devices, alert systems etc.

## Features ⚙️
Below are the features that the service provides.
1. Speed monitoring, the speed of moving assets being tracked is obtained with the use of changing GPS coordinates and time is taken to make the change. This data is used for speed monitoring. When a device speeds past the limit an event is generated which can be consumed by say alert systems.
2. In real-time tracking, the GPS of the tracked asset is sent to the service thus the current location of the assets can be obtained.
Trip historical data, that trip log data is stored in a Dataware house for future queries and analytics.
3. Location detection, the location of an asset at a given time can be determined by querying the log data stored in the warehouse.
Predictive analysis, data collected is used to perform predictive analysis for efficient future event planning.
## Non-Functional Features
1. Scalable, the service is built on a cloud-managed service to accommodate for spontaneous scale. A lightweight protocol(MQTT) is used to efficiently support data exchange between low-end systems.
2. Low latency, data movement between different services ie. ETL pipelines for analytics are reduced by storing and analysing the data from the same place without moving it around.
3. Security, token-based authorisation is utilised to control the devices sending messages to the service.

## Architectural Design ✒️
This layer has the key components of the server-side applications for processing data from connected devices. This layer consists of the message brokers, consumers(handle data processing), producers and storage that meets system requirements.
## The Flow 
1. The MQTT client connects to the broker, authenticated with the provided username and a password. The password is configured in the broker. This basic security feature helps to protect the broker from the only clients with authentication credentials.
2. The connected client can publish or subscribe to a topic of interest. The sensors publish data to a topic, while the client in the server subscribes to these topics, reads the published data and persists it in the database.
3. The data stored in the database is visualised on a dashboard.
4. Web API servers the data stored for extension of the system, or for client integration.
