Projectarium PWA Roadmap:


Phase 1 (Months 1-3): Start Simple
Option 1: All-EC2

Everything on one t2.micro
Understand basics: EC2, security groups, SSH, Docker
Cost: $0


Phase 1.5: ease of use
ECR -> EC2


Phase 1.75: terraform


Phase 2 (Months 4-6): Separate Concerns
Option 2: S3+CloudFront + EC2

Learn CDN concepts
Separate frontend/backend
Static hosting
Cost: $0



Phase 3 (Months 7-9): Managed Database
Option 3: S3+CloudFront + EC2 + RDS

Add managed Postgres
Learn database management, backups
Multi-tier architecture
Cost: $0



Phase 4 (Months 10-12): Container Orchestration
Option 4: ECS on EC2

Learn container orchestration without Fargate costs
Skip ALB (use nginx), keep it free
Task definitions, services
Cost: $0



Ultimate DevOps Portfolio Architecture (After 12 Months):
Route 53 (Custom domain)
         ↓
CloudFront (CDN)
         ↓
S3 (PWA static files)
         ↓ API calls via CloudFront
Application Load Balancer
         ↓
ECS Fargate (Go API containers)
         ↓
RDS Multi-AZ (Postgres)
         
Side: Lambda for async jobs
Side: S3 for file uploads
Side: CloudWatch for monitoring
Side: Secrets Manager for credentials
Side: ECR for Docker images

All defined in Terraform





Cost after free tier: ~$30-40/month
Resume value: 🔥🔥🔥

Quick Decision Matrix:
Goal	                Choose This	            Cost (12mo)
Simplest start	        All-EC2	                $0
Learn AWS basics    	S3+CloudFront+EC2	    $0
Most like production	EC2 + RDS + S3	        $0
Container skills	    ECS on EC2	            $0
Serverless skills	    Lambda + API Gateway	$0
Least effort        	Elastic Beanstalk	    $0-16
