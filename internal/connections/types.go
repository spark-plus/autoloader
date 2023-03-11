package connections

// goal here is to create a common interface for connections from where data can be fetched

type ConnectionType string

const (
	// Connection types for different databases
	PostgresConnType  ConnectionType = "postgres"
	MySQLConnType     ConnectionType = "mysql"
	OracleConnType    ConnectionType = "oracle"
	MSSQLConnType     ConnectionType = "mssql"
	RedshiftConnType  ConnectionType = "redshift"
	SnowflakeConnType ConnectionType = "snowflake"
	SQLiteConnType    ConnectionType = "sqlite"
	BigQueryConnType  ConnectionType = "bigquery"
	HiveConnType      ConnectionType = "hive"
	ImpalaConnType    ConnectionType = "impala"
	PrestoConnType    ConnectionType = "presto"
	VerticaConnType   ConnectionType = "vertica"

	// Connection types for cloud providers
	AWSConnType           ConnectionType = "aws"
	GoogleCloudConnType   ConnectionType = "google_cloud_platform"
	AzureConnType         ConnectionType = "azure"
	S3ConnType            ConnectionType = "s3"
	GCSConnType           ConnectionType = "google_cloud_storage"
	AzureBlobConnType     ConnectionType = "azure_blob_storage"
	AzureDataLakeConnType ConnectionType = "azure_data_lake"

	// Connection types for other systems
	FTPConnType            ConnectionType = "ftp"
	HTTPConnType           ConnectionType = "http"
	SSHConnType            ConnectionType = "ssh"
	DatabricksConnType     ConnectionType = "databricks"
	JiraConnType           ConnectionType = "jira"
	KafkaConnType          ConnectionType = "kafka"
	RabbitMQConnType       ConnectionType = "rabbitmq"
	SlackConnType          ConnectionType = "slack"
	SMTPConnType           ConnectionType = "smtp"
	SFTPConnType           ConnectionType = "sftp"
	WinRMConnType          ConnectionType = "winrm"
	SSHHookConnType        ConnectionType = "sshhook"
	DockerConnType         ConnectionType = "docker"
	KubernetesConnType     ConnectionType = "kubernetes"
	MicrosoftTeamsConnType ConnectionType = "microsoft_teams"
	MicrosoftAzureConnType ConnectionType = "microsoft_azure"
)

type LocationType string

const (
	SSMLocation LocationType = "ssm"
)
