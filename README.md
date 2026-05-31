# Distributed Job Scheduler

A platform for scheduling and distributed execution of tasks.

## About
Backend system that orchestrates job execution across time and multiple processes.
Supports cron, one-shot and delayed execution with retry logic.

## Structure

distributed-job-scheduler/
├── services/
│   ├── api/          # REST gateway, entry point for all requests
│   ├── scheduler/    # decides when a job goes into the queue
│   ├── worker/       # consumes and executes jobs from Kafka
│   └── retry/        # retry logic with exponential backoff
├── libs/
│   ├── common/       # logger, errors, tracing
│   └── proto/        # shared gRPC definitions
├── deploy/           # Docker Compose, Kafka, Postgres, Redis
└── docs/             # architecture and flow diagrams