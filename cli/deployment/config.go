// Code generated by go generate; DO NOT EDIT.
package deployment

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/coder/coder/codersdk"
)

func Config(vip *viper.Viper) (codersdk.DeploymentConfig, error) {
	cfg := codersdk.DeploymentConfig{}
	return cfg, vip.Unmarshal(cfg)
}

func DefaultViper() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("address", "127.0.0.1:3000")
	v.SetDefault("cache_dir", defaultCacheDir())
	v.SetDefault("provisioner_daemon_count", 3)
	v.SetDefault("ssh_keygen_algorithm", "ed25519")
	v.SetDefault("audit_logging", true)
	v.SetDefault("audit_logging", true)

	return v
}

func AttachFlags(flagset *pflag.FlagSet, vip *viper.Viper) {
	_ = flagset.StringP("access-url", "", vip.GetString("access_url"), `External URL to access your deployment. This must be accessible by all provisioned workspaces.`)
	_ = vip.BindPFlag("access_url", flagset.Lookup("access-url"))
	_ = flagset.StringP("wildcard-access-url", "", vip.GetString("wildcard_access_url"), `Specifies the wildcard hostname to use for workspace applications in the form "*.example.com".`)
	_ = vip.BindPFlag("wildcard_access_url", flagset.Lookup("wildcard-access-url"))
	_ = flagset.StringP("address", "", vip.GetString("address"), `Bind address of the server.`)
	_ = vip.BindPFlag("address", flagset.Lookup("address"))
	_ = flagset.StringP("address", "a", vip.GetString("address"), `Bind address of the server.`)
	_ = vip.BindPFlag("address", flagset.Lookup("address"))
	_ = flagset.StringP("address", "a", vip.GetString("address"), `Bind address of the server.`)
	_ = vip.BindPFlag("address", flagset.Lookup("address"))
	_ = flagset.StringP("cache-dir", "", vip.GetString("cache_dir"), `The directory to cache temporary files. If unspecified and $CACHE_DIRECTORY is set, it will be used for compatibility with systemd.`)
	_ = vip.BindPFlag("cache_dir", flagset.Lookup("cache-dir"))
	_ = flagset.StringP("cache-dir", "", vip.GetString("cache_dir"), `The directory to cache temporary files. If unspecified and $CACHE_DIRECTORY is set, it will be used for compatibility with systemd.`)
	_ = vip.BindPFlag("cache_dir", flagset.Lookup("cache-dir"))
	_ = flagset.BoolP("in-memory", "", vip.GetBool("in_memory_database"), `Controls whether data will be stored in an in-memory database.`)
	_ = vip.BindPFlag("in_memory_database", flagset.Lookup("in-memory"))
	_ = flagset.BoolP("in-memory", "", vip.GetBool("in_memory_database"), `Controls whether data will be stored in an in-memory database.`)
	_ = vip.BindPFlag("in_memory_database", flagset.Lookup("in-memory"))
	_ = flagset.MarkHidden("in-memory")
	_ = flagset.IntP("provisioner-daemons", "", vip.GetInt("provisioner_daemon_count"), `Number of provisioner daemons to create on start. If builds are stuck in queued state for a long time, consider increasing this.`)
	_ = vip.BindPFlag("provisioner_daemon_count", flagset.Lookup("provisioner-daemons"))
	_ = flagset.IntP("provisioner-daemons", "", vip.GetInt("provisioner_daemon_count"), `Number of provisioner daemons to create on start. If builds are stuck in queued state for a long time, consider increasing this.`)
	_ = vip.BindPFlag("provisioner_daemon_count", flagset.Lookup("provisioner-daemons"))
	_ = flagset.StringP("postgres-url", "", vip.GetString("postgres_url"), `URL of a PostgreSQL database. If empty, PostgreSQL binaries will be downloaded from Maven (https://repo1.maven.org/maven2) and store all data in the config root. Access the built-in database with "coder server postgres-builtin-url".`)
	_ = vip.BindPFlag("postgres_url", flagset.Lookup("postgres-url"))
	_ = flagset.BoolP("trace", "", vip.GetBool("trace_enable"), `Whether application tracing data is collected.`)
	_ = vip.BindPFlag("trace_enable", flagset.Lookup("trace"))
	_ = flagset.BoolP("secure-auth-cookie", "", vip.GetBool("secure_auth_cookie"), `Controls if the 'Secure' property is set on browser session cookies.`)
	_ = vip.BindPFlag("secure_auth_cookie", flagset.Lookup("secure-auth-cookie"))
	_ = flagset.StringP("ssh-keygen-algorithm", "", vip.GetString("ssh_keygen_algorithm"), `The algorithm to use for generating ssh keys. Accepted values are "ed25519", "ecdsa", or "rsa4096".`)
	_ = vip.BindPFlag("ssh_keygen_algorithm", flagset.Lookup("ssh-keygen-algorithm"))
	_ = flagset.StringP("ssh-keygen-algorithm", "", vip.GetString("ssh_keygen_algorithm"), `The algorithm to use for generating ssh keys. Accepted values are "ed25519", "ecdsa", or "rsa4096".`)
	_ = vip.BindPFlag("ssh_keygen_algorithm", flagset.Lookup("ssh-keygen-algorithm"))
	_ = flagset.BoolP("verbose", "", vip.GetBool("verbose"), `Enables verbose logging.`)
	_ = vip.BindPFlag("verbose", flagset.Lookup("verbose"))
	_ = flagset.BoolP("verbose", "v", vip.GetBool("verbose"), `Enables verbose logging.`)
	_ = vip.BindPFlag("verbose", flagset.Lookup("verbose"))
	_ = flagset.BoolP("audit-logging", "", vip.GetBool("audit_logging"), `Specifies whether audit logging is enabled.`)
	_ = vip.BindPFlag("audit_logging", flagset.Lookup("audit-logging"))
	_ = flagset.BoolP("audit-logging", "", vip.GetBool("audit_logging"), `Specifies whether audit logging is enabled.`)
	_ = vip.BindPFlag("audit_logging", flagset.Lookup("audit-logging"))
	_ = flagset.BoolP("browser-only", "", vip.GetBool("browser_only"), `Whether Coder only allows connections to workspaces via the browser.`)
	_ = vip.BindPFlag("browser_only", flagset.Lookup("browser-only"))
	_ = flagset.StringP("scim-auth-header", "", vip.GetString("scim_auth_header"), `Enables SCIM and sets the authentication header for the built-in SCIM server. New users are automatically created with OIDC authentication.`)
	_ = vip.BindPFlag("scim_auth_header", flagset.Lookup("scim-auth-header"))
	_ = flagset.IntP("user-workspace-quota", "", vip.GetInt("user_workspace_quota"), `Enables and sets a limit on how many workspaces each user can create.`)
	_ = vip.BindPFlag("user_workspace_quota", flagset.Lookup("user-workspace-quota"))
}

func AttachEnterpriseFlags(flagset *pflag.FlagSet, vip *viper.Viper) {
	_ = flagset.BoolP("audit-logging", "", vip.GetBool("audit_logging"), `Specifies whether audit logging is enabled.`)
	_ = vip.BindPFlag("audit_logging", flagset.Lookup("audit-logging"))
	_ = flagset.BoolP("browser-only", "", vip.GetBool("browser_only"), `Whether Coder only allows connections to workspaces via the browser.`)
	_ = vip.BindPFlag("browser_only", flagset.Lookup("browser-only"))
	_ = flagset.StringP("scim-auth-header", "", vip.GetString("scim_auth_header"), `Enables SCIM and sets the authentication header for the built-in SCIM server. New users are automatically created with OIDC authentication.`)
	_ = vip.BindPFlag("scim_auth_header", flagset.Lookup("scim-auth-header"))
	_ = flagset.IntP("user-workspace-quota", "", vip.GetInt("user_workspace_quota"), `Enables and sets a limit on how many workspaces each user can create.`)
	_ = vip.BindPFlag("user_workspace_quota", flagset.Lookup("user-workspace-quota"))
}

