# Databases

## **Instance stores**

Block-level storage volumes behave like physical hard drives.

An [**instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html)** provides temporary block-level storage for an Amazon EC2 instance. An instance store is disk storage that is physically attached to the host computer for an EC2 instance, and therefore has the same lifespan as the instance. When the instance is terminated, you lose any data in the instance store.

To review an example of how instance stores work, choose the arrow buttons to display each step.

## **Amazon Elastic Block Store (Amazon EBS)**

[**Amazon Elastic Block Store (Amazon EBS)**](https://aws.amazon.com/ebs) is a service that provides block-level storage volumes that you can use with Amazon EC2 instances. If you stop or terminate an Amazon EC2 instance, all the data on the attached EBS volume remains available.

To create an EBS volume, you define the configuration (such as volume size and type) and provision it. After you create an EBS volume, it can attach to an Amazon EC2 instance.

Because EBS volumes are for data that needs to persist, it’s important to back up the data. You can take incremental backups of EBS volumes by creating Amazon EBS snapshots.

An Amazon EBS volume stores data in a **single** Availability Zone.

To attach an Amazon EC2 instance to an EBS volume, both the Amazon EC2 instance and the EBS volume must reside within the same Availability Zone.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/UsHkzt64KAl_06th_ruyKsXvVP8ZbeHC1.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/UsHkzt64KAl_06th_ruyKsXvVP8ZbeHC1.png)

## **Object storage**

In **object storage**, each object consists of data, metadata, and a key.

The data might be an image, video, text document, or any other type of file. Metadata contains information about what the data is, how it is used, the object size, and so on. An object’s key is its unique identifier.

## **Amazon Simple Storage Service**

[**Amazon Simple Storage Service (Amazon S3)](https://aws.amazon.com/s3/)** is a service that provides object-level storage. Amazon S3 stores data as objects in buckets. Amazon S3 offers unlimited storage space. The maximum file size for an object in Amazon S3 is 5 TB.

When you upload a file to Amazon S3, you can set permissions to control visibility and access to it. You can also use the Amazon S3 versioning feature to track changes to your objects over time.

There are many **S3 storage classes .** we can select it based on these 2 factors

- How often you plan to retrieve your data
- How available you need your data to be

### S3 Standard

- Designed for frequently accessed data
- Stores data in a minimum of three Availability Zones

Amazon S3 Standard provides high availability for objects. This makes it a good choice for a wide range of use cases, such as websites, content distribution, and data analytics. Amazon S3 Standard has a higher cost than other storage classes intended for infrequently accessed data and archival storage.

### S3 Standard - IA

- Ideal for infrequently accessed data
- Similar to Amazon S3 Standard but has a lower storage price and higher retrieval price

Amazon S3 Standard-IA is ideal for data infrequently accessed but requires high availability when needed. Both Amazon S3 Standard and Amazon S3 Standard-IA store data in a minimum of three Availability Zones. Amazon S3 Standard-IA provides the same level of availability as Amazon S3 Standard but with a lower storage price and a higher retrieval price.

### S3 One Zone - IA

- Stores data in a single Availability Zone
- Has a lower storage price than Amazon S3 Standard-IA

Compared to S3 Standard and S3 Standard-IA, which store data in a minimum of three Availability Zones, S3 One Zone-IA stores data in a single Availability Zone. This makes it a good storage class to consider if the following conditions apply:

- You want to save costs on storage.
- You can easily reproduce your data in the event of an Availability Zone failure.

### S3 Intelligent tiering

- Ideal for data with unknown or changing access patterns
- Requires a small monthly monitoring and automation fee per object

In the S3 Intelligent-Tiering storage class, Amazon S3 monitors objects’ access patterns. If you haven’t accessed an object for 30 consecutive days, Amazon S3 automatically moves it to the infrequent access tier, S3 Standard-IA. If you access an object in the infrequent access tier, Amazon S3 automatically moves it to the frequent access tier, S3 Standard.

### S3 Glacier Instant Retrieval

- Works well for archived data that requires immediate access
- Can retrieve objects within a few milliseconds

When you decide between the options for archival storage, consider how quickly you must retrieve the archived objects. You can retrieve objects stored in the S3 Glacier Instant Retrieval storage class within milliseconds, with the same performance as S3 Standard.

### S3 Glacier Flexible Retrieval

S3 Glacier Flexible Retrieval is a low-cost storage class that is ideal for data archiving. For example, you might use this storage class to store archived customer records or older photos and video files. You can retrieve your data from S3 Glacier Flexible Retrieval from 1 minute to 12 hours.

### S3 Glacier Deep Archive

- Lowest-cost object storage class ideal for archiving
- Able to retrieve objects within 12 hours

S3 Deep Archive supports long-term retention and digital preservation for data that might be accessed once or twice in a year. This storage class is the lowest-cost storage in the AWS Cloud, with data retrieval from 12 to 48 hours. All objects from this storage class are replicated and stored across at least three geographically dispersed Availability Zones.

### S3 Outposts

- Creates S3 buckets on Amazon S3 Outposts
- Makes it easier to retrieve, store, and access data on AWS Outposts

Amazon S3 Outposts delivers object storage to your on-premises AWS Outposts environment. Amazon S3 Outposts is designed to store data durably and redundantly across multiple devices and servers on your Outposts. It works well for workloads with local data residency requirements that must satisfy demanding performance needs by keeping data close to on-premises applications.

## **File storage**

In **file storage**, multiple clients (such as users, applications, servers, and so on) can access data that is stored in shared file folders. In this approach, a storage server uses block storage with a local file system to organize files. Clients access data through file paths.

Compared to block storage and object storage, file storage is ideal for use cases in which a large number of services and resources need to access the same data at the same time.

## Amazon Elastic File System (EFS)

is a scalable file system used with AWS Cloud services and on-premises resources. As you add and remove files, Amazon EFS grows and shrinks automatically. It can scale on demand to petabytes without disrupting applications. 

Amazon EFS is a regional service. It stores data in and across **multiple** Availability Zones. 

The duplicate storage enables you to access data concurrently from all the Availability Zones in the Region where a file system is located. Additionally, on-premises servers can access Amazon EFS using AWS Direct Connect.

## **Amazon Relational Database Service**

[**Amazon Relational Database Service (Amazon RDS)](https://aws.amazon.com/rds/)** is a service that enables you to run relational databases in the AWS Cloud.

Amazon RDS is a managed service that automates tasks such as hardware provisioning, database setup, patching, and backups. With these capabilities, you can spend less time completing administrative tasks and more time using data to innovate your applications. You can integrate Amazon RDS with other services to fulfill your business and operational needs, such as using AWS Lambda to query your database from a serverless application.

Amazon RDS is available on six database engines, which optimize for memory, performance, or input/output (I/O). Supported database engines include:

- Amazon Aurora
- PostgreSQL
- MySQL
- MariaDB
- Oracle Database
- Microsoft SQL Server

## **Amazon Aurora**

[**Amazon Aurora**](https://aws.amazon.com/rds/aurora/) is an enterprise-class relational database. It is compatible with MySQL and PostgreSQL relational databases. It is up to five times faster than standard MySQL databases and up to three times faster than standard PostgreSQL databases.

Amazon Aurora helps to reduce your database costs by reducing unnecessary input/output (I/O) operations, while ensuring that your database resources remain reliable and available.

## **Amazon DynamoDB**

[**Amazon DynamoDB](https://aws.amazon.com/dynamodb/)** is a key-value database service. It delivers single-digit millisecond performance at any scale.

DynamoDB is serverless, which means that you do not have to provision, patch, or manage servers.

You also do not have to install, maintain, or operate software.

As the size of your database shrinks or grows, DynamoDB automatically scales to adjust for changes in capacity while maintaining consistent performance.

This makes it a suitable choice for use cases that require high performance while scaling.

## **Amazon Redshift**

[**Amazon Redshift**](https://aws.amazon.com/redshift) is a data warehousing service that you can use for big data analytics. It offers the ability to collect data from many sources and helps you to understand relationships and trends across your data.

## AWS Database Migration Service

[**AWS Database Migration Service (AWS DMS)](https://aws.amazon.com/dms/)** enables you to migrate relational databases, nonrelational databases, and other types of data stores.

With AWS DMS, you move data between a source database and a target database. [The source and target databases](https://aws.amazon.com/dms/resources) can be of the same type or different types. During the migration, your source database remains operational, reducing downtime for any applications that rely on the database.

There are 2 kinds of migration , homogeneous , i.e. when the source and target are same. and heterogeneous where the source and target are different , here we need to build a converter before using DMS to load the data.

## Amazon DocumentDB

[**Amazon DocumentDB](https://aws.amazon.com/documentdb)** is a document database service that supports MongoDB workloads. (MongoDB is a document database program.)

## Amazon Neptune

[**Amazon Neptune](https://aws.amazon.com/neptune)** is a graph database service.

You can use Amazon Neptune to build and run applications that work with highly connected datasets, such as recommendation engines, fraud detection, and knowledge graphs.

## Amazon QLDB

[**Amazon Quantum Ledger Database (Amazon QLDB)](https://aws.amazon.com/qldb)** is a ledger database service. You can use Amazon QLDB to review a complete history of all the changes that have been made to your application data.

## Amazon Managed Blockchain

[**Amazon Managed Blockchain](https://aws.amazon.com/managed-blockchain)** is a service that you can use to create and manage blockchain networks with open-source frameworks. Blockchain is a distributed ledger system that lets multiple parties run transactions and share data without a central authority.

## Amazon ElastiCache

[**Amazon ElastiCache](https://aws.amazon.com/elasticache)** is a service that adds caching layers on top of your databases to help improve the read times of common requests.

It supports two types of data stores: Redis and Memcached

## DAX

[**Amazon DynamoDB Accelerator (DAX)](https://aws.amazon.com/dynamodb/dax/)** is an in-memory cache for DynamoDB.

It helps improve response times from single-digit milliseconds to microseconds.