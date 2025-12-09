package consts

// --- Environments ---
type Env string

const (
	EnvLocal  Env = "local"  // Local development environment
	EnvDev    Env = "dev"    // Development server
	EnvTest   Env = "test"   // Testing/CI environment
	EnvQA     Env = "qa"     // Quality assurance environment
	EnvStage  Env = "stage"  // Staging/Pre-production environment
	EnvProd   Env = "prod"   // Production environment
	EnvDocker Env = "docker" // Running inside a container
)

// --- General Statuses ---
type Status string

const (
	StatusSuccess  Status = "success"    // Generic success status
	StatusError    Status = "error"      // Generic error status
	StatusFail     Status = "fail"       // Failure (often used in JSend spec)
	StatusPending  Status = "pending"    // Operation is waiting to start
	StatusProcess  Status = "processing" // Operation is in progress
	StatusDone     Status = "done"       // Operation finished
	StatusCanceled Status = "canceled"   // Operation was stopped manually
	StatusSkipped  Status = "skipped"    // Operation was intentionally skipped
	StatusUnknown  Status = "unknown"    // State is not determined
)

// --- HTTP Headers ---
const (
	HeaderAuthorization   = "Authorization"    // Auth credentials
	HeaderContentType     = "Content-Type"     // Media type of the resource
	HeaderAccept          = "Accept"           // Media types acceptable for the response
	HeaderUserAgent       = "User-Agent"       // Client application string
	HeaderXRequestID      = "X-Request-ID"     // Trace ID for request tracking
	HeaderXForwardedFor   = "X-Forwarded-For"  // Originating IP of a client
	HeaderXRealIP         = "X-Real-IP"        // Real IP of the client behind proxy
	HeaderCacheControl    = "Cache-Control"    // Caching directives
	HeaderContentEncoding = "Content-Encoding" // Encoding transformation (gzip)
	HeaderOrigin          = "Origin"           // Origin for CORS
	HeaderLocation        = "Location"         // URL to redirect to
)

// --- Content Types (MIME) ---
type ContentType string

const (
	ContentTypeJSON      ContentType = "application/json"                  // JSON data
	ContentTypeXML       ContentType = "application/xml"                   // XML data
	ContentTypeHTML      ContentType = "text/html"                         // HTML content
	ContentTypeText      ContentType = "text/plain"                        // Plain text
	ContentTypeForm      ContentType = "application/x-www-form-urlencoded" // Form data
	ContentTypeMultipart ContentType = "multipart/form-data"               // File uploads
	ContentTypeBinary    ContentType = "application/octet-stream"          // Binary data
	ContentTypePDF       ContentType = "application/pdf"                   // PDF documents
	ContentTypeCSV       ContentType = "text/csv"                          // Comma-separated values
	ContentTypeJS        ContentType = "application/javascript"            // JavaScript code
	ContentTypeCSS       ContentType = "text/css"                          // CSS styles
	ContentTypePNG       ContentType = "image/png"                         // PNG image
	ContentTypeJPEG      ContentType = "image/jpeg"                        // JPEG image
	CharsetUTF8          string      = "charset=utf-8"                     // Standard UTF-8 charset
)

// --- HTTP Methods ---
type Method string

const (
	MethodGet     Method = "GET"     // Retrieve resource
	MethodPost    Method = "POST"    // Create resource
	MethodPut     Method = "PUT"     // Replace resource
	MethodPatch   Method = "PATCH"   // Partial update
	MethodDelete  Method = "DELETE"  // Remove resource
	MethodOptions Method = "OPTIONS" // Describe communication options
	MethodHead    Method = "HEAD"    // GET without body
)

// --- Time & Date Layouts (Go Reference Time: Mon Jan 2 15:04:05 MST 2006) ---
const (
	LayoutDateISO     = "2006-01-02"          // Standard ISO 8601 Date
	LayoutTimeISO     = "15:04:05"            // Standard ISO 8601 Time
	LayoutDateTimeISO = "2006-01-02 15:04:05" // Date and Time space separated
	LayoutDateCompact = "20060102"            // Compact date for filenames
	LayoutYearMonth   = "2006-01"             // Year and Month
	LayoutUSDate      = "01/02/2006"          // US style date
	LayoutEUDates     = "02.01.2006"          // European style date
)

// --- Authentication & Identity ---
type Role string

const (
	AuthSchemeBearer      = "Bearer" // Bearer token scheme
	AuthSchemeBasic       = "Basic"  // Basic auth scheme
	RoleAdmin        Role = "admin"  // Administrator role
	RoleUser         Role = "user"   // Standard user role
	RoleGuest        Role = "guest"  // Unauthenticated user
	RoleSystem       Role = "system" // Internal system account
)

// --- Database & Pagination ---
const (
	OrderAsc        = "ASC"  // Ascending sort order
	OrderDesc       = "DESC" // Descending sort order
	DefaultPage     = 1      // Default page number
	DefaultPageSize = 20     // Default items per page
	MaxPageSize     = 100    // Safety limit for page size
)

// --- Context Keys ---
type contextKey string

const (
	CtxKeyUserID    contextKey = "user_id"    // Context key for User ID
	CtxKeyRole      contextKey = "user_role"  // Context key for User Role
	CtxKeyRequestID contextKey = "request_id" // Context key for Request Trace ID
	CtxKeyLogger    contextKey = "logger"     // Context key to pass logger instance
	CtxKeyToken     contextKey = "auth_token" // Context key for raw token
)

// --- Common Symbols & Chars ---
const (
	EmptyString = ""   // Empty string literal
	Space       = " "  // Single space
	Comma       = ","  // Comma separator
	Dot         = "."  // Dot separator
	Slash       = "/"  // Forward slash
	Underscore  = "_"  // Underscore
	NewLine     = "\n" // New line character
)
