package constants

// connection information
var (
	MySQLDefaultDSN      string
	MinioEndPoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioSSL             bool
	RedisAddr            string
	RedisPassword        string
	MongoDefaultDSN      string
	Host                 string
	EtcdAddress          string
	SecretKey            string
)

// constants in the project
const (
	MinioImgBucketName = "imagebucket"

	DefaultSign       = "该用户没有留下任何签名"
	DefaultAva        = "imagebucket/mols.jpg"
	DefaultBackground = "imagebucket/marisa.jpg"
	DefaultUserName   = "请设置你的昵称"

	ArticleService  = "aufer-article-service"
	CommentService  = "aufer-comment-service"
	RelationService = "aufer-relation-service"
	UserService     = "aufer-user-service"
	ApiService      = "aufer-api-service"

	CPURateLimit float64 = 80.0
	DefaultLimit         = 10
)
