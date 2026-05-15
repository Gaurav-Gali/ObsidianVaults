![Source](https://www.youtube.com/watch?v=5pniK1RV_6o&list=PLupK5DK91flV45dkPXyGViMLtHadRr6sp)

![[mlops-diagram.png|306]]

==MLOPS== (Machine Learning Operations) - are the set of practises and automations that are used to automate and simplify ML workflows and deployments.
This makes the development and deployment process more disciplined and structured.

<hr>

### DS Lifecycle 
1. EDA (Exploratory Data Analysis)
2. Pre-Processing
3. Feature Engineering
4. Feature Selection
5. Model Training & Hyper-parameter Tuning
6. Model Evaluation
7. App Building / UI
8. Deployment

<hr>

### Issues with DS practices without MLOPS
1. Low coding standards
	1. OOPS concepts
	2. Modular coding
	3. Logging
	4. Exception Handling
2. No Data Management
	1. Data Ingestion problems
	2. Artifact
3. Versioning
	1. Code
	2. Data
	3. Model
4. Data Pipeline / Experiments
5. No CICD concept (Continuous Integration & Continuous Deployment)
6. Scalability & Monitoring (Production)
	1. Kubernetes
	2. Prometheus
	3. Grafana
7. Cross team friction

<hr>

### SDLC (Software Development Lifecycle)
![[sdlc.png]]

<hr>

### Improvements
1. Code Standards
	1. OOPS
	2. Modular Coding
	3. Logging
2. Code versioning
	1. Git & Github
3. Data / Model versioning
	1. DVC
	2. MLFlow
4. CICD tools
	1. Github actions
	2. Circle CI
	3. Travis CI
5. Containerisation
	1. Docker
	2. Docker hub
6. Scalability & Monitoring
	1. Kubernetes
	2. Prometheus
	3. Grafana
7. AWS Services
	1. IAM User
	2. ECR
	3. S3
	4. EC2
8. All in one services
	1. AWS Sagemaker
	2. Google Vertex AI
	3. Azure ML
