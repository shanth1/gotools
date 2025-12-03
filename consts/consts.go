package consts

// --- Environments ---
const (
	EnvLocal  = "local"  // Local development environment
	EnvDev    = "dev"    // Development server
	EnvTest   = "test"   // Testing/CI environment
	EnvQA     = "qa"     // Quality assurance environment
	EnvStage  = "stage"  // Staging/Pre-production environment
	EnvProd   = "prod"   // Production environment
	EnvDocker = "docker" // Running inside a container
)

// --- General Statuses ---
const (
	StatusSuccess  = "success"    // Generic success status
	StatusError    = "error"      // Generic error status
	StatusFail     = "fail"       // Failure (often used in JSend spec)
	StatusPending  = "pending"    // Operation is waiting to start
	StatusProcess  = "processing" // Operation is in progress
	StatusDone     = "done"       // Operation finished
	StatusCanceled = "canceled"   // Operation was stopped manually
	StatusSkipped  = "skipped"    // Operation was intentionally skipped
	StatusUnknown  = "unknown"    // State is not determined
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
const (
	ContentTypeJSON      = "application/json"                  // JSON data
	ContentTypeXML       = "application/xml"                   // XML data
	ContentTypeHTML      = "text/html"                         // HTML content
	ContentTypeText      = "text/plain"                        // Plain text
	ContentTypeForm      = "application/x-www-form-urlencoded" // Form data
	ContentTypeMultipart = "multipart/form-data"               // File uploads
	ContentTypeBinary    = "application/octet-stream"          // Binary data
	ContentTypePDF       = "application/pdf"                   // PDF documents
	ContentTypeCSV       = "text/csv"                          // Comma-separated values
	ContentTypeJS        = "application/javascript"            // JavaScript code
	ContentTypeCSS       = "text/css"                          // CSS styles
	ContentTypePNG       = "image/png"                         // PNG image
	ContentTypeJPEG      = "image/jpeg"                        // JPEG image
	CharsetUTF8          = "charset=utf-8"                     // Standard UTF-8 charset
)

// --- HTTP Methods ---
const (
	MethodGet     = "GET"     // Retrieve resource
	MethodPost    = "POST"    // Create resource
	MethodPut     = "PUT"     // Replace resource
	MethodPatch   = "PATCH"   // Partial update
	MethodDelete  = "DELETE"  // Remove resource
	MethodOptions = "OPTIONS" // Describe communication options
	MethodHead    = "HEAD"    // GET without body
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
const (
	AuthSchemeBearer = "Bearer" // Bearer token scheme
	AuthSchemeBasic  = "Basic"  // Basic auth scheme
	RoleAdmin        = "admin"  // Administrator role
	RoleUser         = "user"   // Standard user role
	RoleGuest        = "guest"  // Unauthenticated user
	RoleSystem       = "system" // Internal system account
)

// --- Database & Pagination ---
const (
	OrderAsc        = "ASC"  // Ascending sort order
	OrderDesc       = "DESC" // Descending sort order
	DefaultPage     = 1      // Default page number
	DefaultPageSize = 20     // Default items per page
	MaxPageSize     = 100    // Safety limit for page size
)

// --- Context Keys (Strings mainly used for Middleware <-> Handler) ---
const (
	CtxKeyUserID    = "user_id"    // Context key for User ID
	CtxKeyRole      = "user_role"  // Context key for User Role
	CtxKeyRequestID = "request_id" // Context key for Request Trace ID
	CtxKeyLogger    = "logger"     // Context key to pass logger instance
	CtxKeyToken     = "auth_token" // Context key for raw token
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
