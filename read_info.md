# --- functional requirements
users: user sign-up/login, vertification otp/sms, seller/buyer
catalog: product listing, mamage products(CRUD), stock management
payment: purchase (card/online/bank)
notifications: eamil, sms

# --- 
How is Go Ecosystem working
Go Paradime and style of programming
Go convention and functional programming
Build rest API 
Connect with DB using ORM
Usages of Clean architecture 
How GO performs in the Dev and Production environment
Elastic Search with GO
Microservice in GO
GraphQL with GO
Go Deployment and Scalability

# --- architecture
users - route53 -amplity - react app - route 53 - ALB - App Server - S3 - Elasticsearch - RDS - 3rd party

# --- infrastucture

# --- blueprint
graphQL / REST API - Monolithic(users, actalogue, 3rd party(email, sms, payment), transaction, elastich search) - microservices(users, transaction,... - streams - kafka) - storage(rds, s3)

# --- design pattern
database design

# --- aws cloudfornt distribution & s3 operation
backend - webapp - cloudFront Distribution - s3 bucket (private)
(image, url, request)

service endpoints
client - users, seller

# --- prepare our project

# --- clearn architecture & solid princcipal

Framework and device: Devices, DB, Web, UI, external  interfaces
interface adapters: controllers, gateways, presenters
application bussiness: use cases
enterprise bussiness: entities

# --- go serveice & business - Notification
send notification
get status

twilio, SNS, SES

# --- Cart and Orders
create Cart & Order systems like Big E-commerce application

Cart(preserve the cart items in FE) - Collect the payment - create order - Order - ship the order items

Cart - Store all art items in Be/database - Order (cart items transferred to Order)

Collecthe payment - success - Store all cart items

Cart - continue to place order - createing order - collecting payment - order confirmed/ order failed - order detail


User - api gateway - catalog - inventory service - seller service - vendor




