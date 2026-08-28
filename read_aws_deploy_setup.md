1. Create DB - RDS
standard create
postgreSQL
template - free tier
credential setting - self managed

set name, master password

instance configure - db.t4.micro
public access - public
databse port
password authentication
monitoring - enable trace - database name
backup - 1 day

user VPC custome

avalilable zones


2. Create VPC
create security group

inbound rules:
postgresql - 5432 - tcp - anywhere

outbound rules:

3. Connect table plus with user + password aws RDS

4. Create IAM
create user name
set permission
permission policies: AdminstratorAccess

create access key
commandline CLI

5. Elastic beanstalk
creat appliaction
configure environment - web server environment

aplication information

environemtn name
domain

platform: type, branch, version
application code
presets - single instance

configure service access:
service role
ec2 instance profile

setup networking databse and tags
VPC
instance settings
public Ip address
databse setting

configure instance traffic and scaling
root volume (boot device)
size, IOPS, throughput
Ec2 security groups

configure updates, monitoring, logging
managed platform updates
email notification
rolling update and deployments
instance log streaming to cloudwatch

review

