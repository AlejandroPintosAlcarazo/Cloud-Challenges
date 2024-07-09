# ETL Weather Map Project Documentation

## Table of Contents
1. [General Info](#general-info)
2. [Technologies and Services](#technologies-and-services)
3. [Collaboration](#collaboration)
4. [FAQs](#faqs)
5. [Problems](#problems)

## General Info

### Introduction
This project aims to create an ETL (Extract, Transform, Load) pipeline to obtain, process, and visualize weather data on a weather map in Looker.

### Project Objective
The main goal is to automate the extraction and processing of weather data to generate accurate and up-to-date visualizations in Looker.

## Technologies and Services
- Google Cloud Scheduler
- Google Cloud Workflows
- Google Cloud Run
- Google BigQuery
- Dataform
- Terraform
- Looker
- Docker

## Collaboration
We welcome contributions from the community. Please read the contributing guidelines before submitting a pull request.

Collaborators:
- Altostratus

## FAQs
**Q: How often is the data updated?**  
A: The data is updated according to the schedule defined in Cloud Scheduler.

**Q: Can I add new weather data sources?**  
A: Yes, you can extend the project by adding new connectors in the fetcher module.

## Problems
There was an issue linking a view to the reporting table from processing due to inexperience with Dataform, which prevented the weather map from being exported to Looker. Further investigation is required to resolve this linkage problem.

## Project Architecture

### Components
- **Terraform**: Deploys the entire project infrastructure on Google Cloud.
- **Connector (Aemet-ELT)**: Divided into several modules to fetch, load, and query weather data.
  - `cmd`: Contains main commands for setup and data update.
  - `fetcher`: Extracts data from the Aemet API.
  - `load`: Uploads extracted data to BigQuery.
  - `query`: Executes database queries.
- **Dataform**: Transforms raw data from the API, corrects typos, and creates the reporting table in BigQuery for weather map visualization.

### Data Flow
1. **Terraform**: Deploys infrastructure (Cloud Run, Cloud Scheduler, BigQuery, Dataform, Looker).
2. **Connector (Aemet-ELT)**: Fetches, processes, and loads weather data into BigQuery.
3. **Dataform**: Transforms and processes raw data into a final format ready for reporting.

## Detailed Description

### Aemet-ELT

This folder contains the core logic for the ETL project, including configuration, data fetching, processing, and storage.

- **cmd**
  - `setup/main.go`: Main script for initial environment and database setup.
  - `updater/main.go`: Main script for updating existing data in the database.
- **configs**
  - Configuration files and constants used in the project.
- **fetcher**
  - Module to extract data from the Aemet API, including functions for making HTTP requests, parsing, and transforming the fetched data.
- **load**
  - Module to load the fetched data into BigQuery, including functions for storing the data in the staging database.
- **query**
  - Module to create and update tables in BigQuery, including functions for running queries, configuring tables, and setting schemas.
- **utils**
  - Utility functions used throughout the project, such as CloudRun connection.

### Dataform

This folder contains the configuration and scripts for data transformation using Dataform.

**Functionality**: Dataform processes raw data from the API, corrects typographical errors, and creates the reporting table in BigQuery, resulting in a ready-to-visualize weather map.

### Terraform

This folder contains the infrastructure-as-code configuration using Terraform.

**Functionality**: Terraform deploys the entire project infrastructure on Google Cloud, including Cloud Run, Cloud Scheduler, BigQuery, Dataform, and Looker, as per the project architecture.
