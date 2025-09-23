2025-09-23T09:42:57.934Z Application Logs v1.0 Development INFO 2025/09/23 09:42:57 Starting auth service...
2025-09-23T09:42:57.940Z Application Logs v1.0 Development INFO 2025/09/23 09:42:57 🔄 Initializing database connection...
2025-09-23T09:42:57.940Z Application Logs v1.0 Development INFO 2025/09/23 09:42:57 Connecting to database...
2025-09-23T09:42:57.940Z Application Logs v1.0 Development INFO 2025/09/23 09:42:57 Resolving hostname: aws-1-ap-southeast-1.pooler.supabase.com
2025-09-23T09:42:58.022Z Application Logs v1.0 Development INFO 2025/09/23 09:42:58 Found IPv4 address: 13.213.241.248
2025-09-23T09:42:58.022Z Application Logs v1.0 Development INFO 2025/09/23 09:42:58 Updated database host to IPv4: 13.213.241.248:6543
2025-09-23T09:42:58.022Z Application Logs v1.0 Development INFO 2025/09/23 09:42:58 Creating database connection pool...
2025-09-23T09:42:58.022Z Application Logs v1.0 Development INFO 2025/09/23 09:42:58 ✅ Database pool created successfully
2025-09-23T09:42:58.022Z Application Logs v1.0 Development INFO 2025/09/23 09:42:58 Testing database connection...
2025-09-23T09:43:08.026Z Application Logs v1.0 Development INFO 2025/09/23 09:43:08 ❌ Failed to ping database: context deadline exceeded
2025-09-23T09:43:08.026Z Application Logs v1.0 Development WARN 2025/09/23 09:43:08 ❌ WARNING: Failed to connect to database: context deadline exceeded
2025-09-23T09:43:08.026Z Application Logs v1.0 Development INFO 2025/09/23 09:43:08 ⚠️  Service will start without database connectivity
2025-09-23T09:43:08.026Z Application Logs v1.0 Development INFO 2025/09/23 09:43:08 ⚠️  Database-dependent endpoints will not work until connection is established
2025-09-23T09:43:08.026Z Application Logs v1.0 Development INFO 2025/09/23 09:43:08 Auth Service starting on port 8080
2025-09-23T09:43:08.026Z Application Logs v1.0 Development INFO 2025/09/23 09:43:08 Health endpoint available at: http://localhost:8080/health
2025-09-23T09:43:14.959Z Application Logs v1.0 Development INFO [GIN] 2025/09/23 - 09:43:14 | 200 |  5.000274981s |   10.100.24.118 | GET      "/auth-service/health"
2025-09-23T09:43:16.476Z Gateway Logs v1.0 Development 200 SERVICE_RESPONSE Method="GET" RequestPath="/sts-test/auth-service/v1.0/health" ServicePath="/auth-service/health" UserAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36" CorrelationID="fe6a110a9ff6b5c4f05c0d26c8132e34" ServiceHost="auth-service-2401972695.dp-development-ststest-75412-4003694047.svc.cluster.local:8080" Duration="23309"
2025-09-23T09:44:00.653Z Application Logs v1.0 Development INFO 2025/09/23 09:44:00 📝 Register endpoint called
2025-09-23T09:44:00.653Z Application Logs v1.0 Development INFO 2025/09/23 09:44:00 🔍 Checking database connectivity...
2025-09-23T09:44:05.657Z Application Logs v1.0 Development INFO 2025/09/23 09:44:05 ❌ Database not connected, returning 503
2025-09-23T09:44:05.657Z Application Logs v1.0 Development INFO [GIN] 2025/09/23 - 09:44:05 | 503 |  5.003951772s |   10.100.38.237 | POST     "/auth-service/api/auth/register"
2025-09-23T09:44:07.307Z Gateway Logs v1.0 Development 503 SERVICE_RESPONSE Method="POST" RequestPath="/sts-test/auth-service/v1.0/api/auth/register" ServicePath="/auth-service/api/auth/register" UserAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36" CorrelationID="7b2698dc2c161df0d4e94ce36846ccb3" ServiceHost="auth-service-2401972695.dp-development-ststest-75412-4003694047.svc.cluster.local:8080" Duration="5041"