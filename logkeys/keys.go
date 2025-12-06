package logkeys

// --- Common ---
const (
	Metadata = "metadata" // Generic object/map dump
	Payload  = "payload"  // Request/Message payload dump
	Raw      = "raw"      // Raw data dump (bytes/string)
	Tags     = "tags"     // Array of string tags
)

// --- System, Runtime & Build ---
const (
	Env        = "env"        // Environment name (e.g., prod, dev, staging)
	App        = "app"        // Application name
	Service    = "service"    // Microservice name
	Instance   = "instance"   // Instance ID or hostname
	Version    = "version"    // Application semantic version
	GitHash    = "git_hash"   // Git commit hash
	BuildTime  = "build_time" // Time when binary was built
	Component  = "component"  // Internal component/module name
	PID        = "pid"        // Process ID
	GoVersion  = "go_version" // Go runtime version
	Goroutines = "goroutines" // Number of active goroutines
	Memory     = "memory_mb"  // Allocated memory in megabytes
	Caller     = "caller"     // File and line number of the log call
)

// --- Tracing & Observability (OpenTelemetry) ---
const (
	TraceID  = "trace_id"  // Distributed trace identifier
	SpanID   = "span_id"   // Current span identifier
	ParentID = "parent_id" // Parent span identifier
	Sampled  = "sampled"   // Whether this trace is sampled
)

// --- HTTP Request (Server & Client) ---
const (
	HTTPMethod    = "http_method"     // HTTP request method (GET, POST, etc.)
	HTTPRoute     = "http_route"      // Registered route template (e.g., /users/:id)
	HTTPStatus    = "http_status"     // HTTP response status code
	HTTPPath      = "http_path"       // Actual URL path
	HTTPHost      = "http_host"       // Host header value
	HTTPUserAgent = "http_user_agent" // Client user agent string
	HTTPReferer   = "http_referer"    // Referer URL
	HTTPProto     = "http_proto"      // HTTP protocol version (1.1, 2.0)
	RemoteIP      = "remote_ip"       // IP address of the immediate caller
	ClientIP      = "client_ip"       // True client IP (usually from X-Forwarded-For)
	RequestID     = "request_id"      // Unique ID for the HTTP request
	Latency       = "latency"         // Duration of the operation
	LatencyMS     = "latency_ms"      // Duration in milliseconds
	BytesIn       = "bytes_in"        // Request body size in bytes
	BytesOut      = "bytes_out"       // Response body size in bytes
	ContentType   = "content_type"    // Content-Type header value
)

// --- WebSockets & Real-time ---
const (
	WSConnID  = "ws_conn_id"  // Unique WebSocket connection ID
	WSChannel = "ws_channel"  // Channel, room, or topic name
	WSEvent   = "ws_event"    // Event type (connect, disconnect, message)
	WSMsgType = "ws_msg_type" // Message type (text, binary, ping, pong)
	WSMsgSize = "ws_msg_size" // Size of the payload
)

// --- gRPC & RPC ---
const (
	RPCService = "rpc_service" // gRPC service name
	RPCMethod  = "rpc_method"  // gRPC method name
	RPCCode    = "rpc_code"    // gRPC status code
	RPCDetails = "rpc_details" // Detailed gRPC status message
)

// --- Database (SQL & NoSQL) ---
const (
	DBSystem    = "db_system"    // Database type (postgres, mysql, mongo)
	DBHost      = "db_host"      // Database server address
	DBName      = "db_name"      // Database name
	DBOperation = "db_operation" // Operation type (SELECT, INSERT, UPDATE)
	DBTable     = "db_table"     // Target table or collection
	DBQuery     = "db_query"     // The actual query statement (sanitized)
	DBDuration  = "db_duration"  // Time taken for the DB query
	DBRows      = "db_rows"      // Number of rows affected or returned
	DBTxID      = "db_tx_id"     // Database transaction ID
)

// --- Cache (Redis, Memcached) ---
const (
	CacheSystem = "cache_system" // Cache system name
	CacheKey    = "cache_key"    // Cache key accessed
	CacheHit    = "cache_hit"    // Boolean indicating a cache hit
	CacheOp     = "cache_op"     // Cache operation (GET, SET, DEL)
	CacheTTL    = "cache_ttl"    // Time-to-live value
)

// --- Application Logic & Business ---
const (
	Event       = "event"        // Name of the event occurring
	Action      = "action"       // Specific action taken by user or system
	Category    = "category"     // High-level category of the log
	Resource    = "resource"     // Name of the resource being manipulated
	ResourceID  = "resource_id"  // ID of the resource
	TargetID    = "target_id"    // Target entity ID (legacy alias for ResourceID)
	TargetType  = "target_type"  // Type of the target entity
	Outcome     = "outcome"      // Result of logic (success, failure, skipped)
	Attempt     = "attempt"      // Retry attempt number
	MaxAttempts = "max_attempts" // Maximum allowed attempts
)

// --- User, Auth & Identity ---
const (
	UserID      = "user_id"     // Unique identifier of the user
	UserEmail   = "user_email"  // User's email address (PII warning)
	UserRole    = "user_role"   // Role or permission level of the user
	TenantID    = "tenant_id"   // Multi-tenancy identifier
	AccountID   = "account_id"  // Account or organization ID
	SessionID   = "session_id"  // User session identifier
	AuthMethod  = "auth_method" // Authentication method (jwt, basic, oauth)
	TokenID     = "token_id"    // Unique token identifier (JTI)
	Scopes      = "scopes"      // OAuth scopes granted
	Fingerprint = "fingerprint" // Browser or device fingerprint
)

// --- Security & Audit ---
const (
	ClientGeo    = "client_geo"    // Country or region code from IP
	RiskScore    = "risk_score"    // Fraud/Risk score (0.0 - 1.0)
	ThreatType   = "threat_type"   // Type of threat (sql_injection, xss, brute_force)
	AttackVector = "attack_vector" // How the attack was delivered
	ACLPolicy    = "acl_policy"    // Name of the ACL policy evaluated
	Permission   = "permission"    // Specific permission checked
	CipherSuite  = "cipher_suite"  // TLS cipher suite used
	TLSVersion   = "tls_version"   // TLS version (1.2, 1.3)
)

// --- Errors & Security ---
const (
	Error       = "error"        // String representation of the error
	ErrorType   = "error_type"   // Class or type of the error
	ErrorStack  = "error_stack"  // Stack trace related to the error
	ErrorCode   = "error_code"   // Application specific error code
	PanicReason = "panic_reason" // Value recovered from a panic
	Blocked     = "blocked"      // Boolean indicating request was blocked
	Reason      = "reason"       // Human-readable reason for an action/error
)

// --- Data Validation ---
const (
	Field        = "field"        // Name of the field being validated
	Constraint   = "constraint"   // Rule that was violated (e.g., required, min_len)
	InvalidValue = "invalid_val"  // The value that failed validation (sanitize PII!)
	InputSource  = "input_source" // Source of input (query, body, header)
)

// --- Messaging & Async (Kafka, RabbitMQ, SQS) ---
const (
	MsgSystem   = "msg_system"   // Messaging system (kafka, rabbitmq)
	MessageID   = "message_id"   // Unique ID of the message
	Topic       = "topic"        // Topic or exchange name
	Queue       = "queue"        // Queue name
	Partition   = "partition"    // Partition ID
	Offset      = "offset"       // Message offset in the partition
	ConsumerGrp = "consumer_grp" // Consumer group name
	WorkerID    = "worker_id"    // ID of the background worker
	JobType     = "job_type"     // Type of background job
	JobID       = "job_id"       // Unique job execution ID
)

// --- Infrastructure, Cloud & Containerization ---
const (
	Region      = "region"       // Cloud region (e.g., us-east-1)
	Zone        = "zone"         // Availability zone
	Cluster     = "cluster"      // Cluster name
	Node        = "node"         // Node name or IP
	Pod         = "pod"          // Kubernetes pod name
	Namespace   = "namespace"    // Kubernetes namespace
	Container   = "container"    // Container name
	ContainerID = "container_id" // Docker/CRI container ID
	Image       = "image"        // Container image name/tag
)

// --- Feature Flags & Configuration ---
const (
	FlagKey     = "flag_key"     // Feature flag identifier
	FlagVariant = "flag_variant" // Variant of the flag served
	ConfigKey   = "config_key"   // Configuration parameter name
)

// --- External APIs (Third Party) ---
const (
	Provider    = "provider"     // Name of third-party provider (e.g., Stripe, AWS)
	ExtEndpoint = "ext_endpoint" // External API endpoint called
	ExtStatus   = "ext_status"   // HTTP status from external service
)

// --- File System & I/O ---
const (
	FileName = "file_name" // Name of the file being accessed
	FilePath = "file_path" // Full path to the file
	FileSize = "file_size" // Size of the file in bytes
	FileMode = "file_mode" // File permissions mode
)

// --- Finance & Transactions ---
const (
	Currency     = "currency"     // ISO 4217 currency code (USD, EUR)
	Amount       = "amount"       // Monetary amount (decimal or minor units)
	OrderID      = "order_id"     // Unique order identifier
	PaymentID    = "payment_id"   // Payment gateway transaction ID
	Subscription = "subscription" // Subscription ID
	Gateway      = "gateway"      // Payment gateway name (Stripe, PayPal)
	Wallet       = "wallet"       // Wallet address (crypto) or ID
)

// --- Notifications (Email, SMS, Push) ---
const (
	Recipient  = "recipient"   // Identifier of receiver (masked email/phone)
	Sender     = "sender"      // Sender identity (e.g., noreply@...)
	Subject    = "subject"     // Email subject or notification title
	TemplateID = "template_id" // ID of the template used
	Channel    = "channel"     // Delivery channel (email, sms, push, slack)
	DeliveryID = "delivery_id" // ID provided by the delivery provider (SendGrid, Twilio ID)
)

// --- AI & LLM Operations ---
const (
	Model       = "model"       // Model name (e.g., gpt-4, llama-2)
	TokensIn    = "tokens_in"   // Prompt token count
	TokensOut   = "tokens_out"  // Completion token count
	Temperature = "temperature" // Model temperature setting
	VectorID    = "vector_id"   // ID in vector database
)

// --- OS & Lifecycle ---
const (
	Signal   = "signal"    // OS signal received (SIGINT, SIGTERM)
	ExitCode = "exit_code" // Process exit code
	Uptime   = "uptime"    // Application uptime duration
)
