package logkeys

// --- System & Infrastructure ---
const (
	Env       = "env"
	App       = "app"
	Service   = "service"
	Instance  = "instance"
	Version   = "version"
	Component = "component"
	PID       = "pid"
)

// --- HTTP Request (Server & Client) ---
const (
	HTTPMethod    = "http_method"
	HTTPRoute     = "http_route"
	HTTPStatus    = "http_status"
	HTTPPath      = "http_path"
	HTTPUserAgent = "http_user_agent"
	HTTPReferer   = "http_referer"
	RemoteIP      = "remote_ip"
	RequestID     = "request_id"
	Latency       = "latency"
	BytesIn       = "bytes_in"
	BytesOut      = "bytes_out"
)

// --- Database & Storage ---
const (
	DBSystem    = "db_system"
	DBOperation = "db_operation"
	DBTable     = "db_table"
	DBQuery     = "db_query"
	DBDuration  = "db_duration"
	DBRows      = "db_rows"
)

// --- Application Logic & Business ---
const (
	Event      = "event"
	Action     = "action"
	TargetID   = "target_id"
	TargetType = "target_type"
)

// --- User & Identity ---
const (
	UserID    = "user_id"
	UserRole  = "user_role"
	TenantID  = "tenant_id"
	SessionID = "session_id"
)

// --- Errors & Stacktrace ---
const (
	Error      = "error"
	ErrorStack = "error_stack"
	ErrorType  = "error_type"
)

// --- Async / Queue (Kafka, RabbitMQ) ---
const (
	MessageID   = "message_id"
	Topic       = "topic"
	Partition   = "partition"
	Offset      = "offset"
	ConsumerGrp = "consumer_grp"
)
