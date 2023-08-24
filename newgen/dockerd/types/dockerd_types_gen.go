package types

type DefaultCgroupnsModeEnum string

const (
	DefaultCgroupnsModePrivate DefaultCgroupnsModeEnum = "private"
	DefaultCgroupnsModeHost    DefaultCgroupnsModeEnum = "host"
)

type DefaultIpcModeEnum string

const (
	DefaultIpcModeShareable DefaultIpcModeEnum = "shareable"
	DefaultIpcModePrivate   DefaultIpcModeEnum = "private"
)

// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
type Dockerd struct {
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	UsernsRemap *string `json:"userns-remap,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	SeccompProfile *string `json:"seccomp-profile,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Ip6tables *bool `json:"ip6tables,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	AllowNondistributableArtifacts []any `json:"allow-nondistributable-artifacts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ApiCorsHeader *string `json:"api-cors-header,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DnsSearch []string `json:"dns-search,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	CgroupParent *string `json:"cgroup-parent,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DataRoot *string `json:"data-root,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Debug *bool `json:"debug,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Ip *string `json:"ip,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	MaxDownloadAttempts *int64 `json:"max-download-attempts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	StorageDriver *string `json:"storage-driver,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Group *string `json:"group,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Labels []any `json:"labels,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Bip *string `json:"bip,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	UserlandProxy *bool `json:"userland-proxy,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	InsecureRegistries []string `json:"insecure-registries,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultShmSize *string `json:"default-shm-size,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	LiveRestore *bool `json:"live-restore,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultIpcMode *DefaultIpcModeEnum `json:"default-ipc-mode,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ContainerdNamespace *string `json:"containerd-namespace,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	LogLevel *string `json:"log-level,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Icc *bool `json:"icc,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	LogOpts struct {
		CacheMaxSize  *string `json:"cache-max-size,omitempty"`
		CacheCompress *string `json:"cache-compress,omitempty"`
		Env           *string `json:"env,omitempty"`
		Labels        *string `json:"labels,omitempty"`
		MaxFile       *string `json:"max-file,omitempty"`
		MaxSize       *string `json:"max-size,omitempty"`
		CacheDisabled *string `json:"cache-disabled,omitempty"`
		CacheMaxFile  *string `json:"cache-max-file,omitempty"`
	} `json:"log-opts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	AuthorizationPlugins []any `json:"authorization-plugins,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DnsOpts []string `json:"dns-opts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Tlsverify *bool `json:"tlsverify,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	LogDriver *string `json:"log-driver,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultCgroupnsMode *DefaultCgroupnsModeEnum `json:"default-cgroupns-mode,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultGateway *string `json:"default-gateway,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	UserlandProxyPath *string `json:"userland-proxy-path,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Tlscacert *string `json:"tlscacert,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ClusterStoreOpts struct {
	} `json:"cluster-store-opts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ExecOpts []string `json:"exec-opts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Runtimes struct {
		Custom struct {
			RuntimeArgs []string `json:"runtimeArgs,omitempty"`
			Path        *string  `json:"path,omitempty"`
		} `json:"custom,omitempty"`
		CcRuntime struct {
			Path *string `json:"path,omitempty"`
		} `json:"cc-runtime,omitempty"`
	} `json:"runtimes,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ContainerdPluginNamespace *string `json:"containerd-plugin-namespace,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	IpMasq *bool `json:"ip-masq,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Ipv6 *bool `json:"ipv6,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	OomScoreAdjust *int64 `json:"oom-score-adjust,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Tlskey *string `json:"tlskey,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Dns []any `json:"dns,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	InitPath *string `json:"init-path,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Tlscert *string `json:"tlscert,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ClusterStore *string `json:"cluster-store,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultAddressPools []struct {
		Size *int64  `json:"size,omitempty"`
		Base *string `json:"base,omitempty"`
	} `json:"default-address-pools,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Hosts []string `json:"hosts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultUlimits struct {
		Nofile struct {
			Name *string `json:"Name,omitempty"`
			Soft *int64  `json:"Soft,omitempty"`
			Hard *int64  `json:"Hard,omitempty"`
		} `json:"nofile,omitempty"`
	} `json:"default-ulimits,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	FixedCidrV6 *string `json:"fixed-cidr-v6,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Tls *bool `json:"tls,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	NoNewPrivileges *bool `json:"no-new-privileges,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	NodeGenericResources []string `json:"node-generic-resources,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Experimental *bool `json:"experimental,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	SelinuxEnabled *bool `json:"selinux-enabled,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	RawLogs *bool `json:"raw-logs,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ShutdownTimeout *int64 `json:"shutdown-timeout,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	MaxConcurrentUploads *int64 `json:"max-concurrent-uploads,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Mtu *int64 `json:"mtu,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultRuntime *string `json:"default-runtime,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	DefaultGatewayV6 *string `json:"default-gateway-v6,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	SwarmDefaultAdvertiseAddr *string `json:"swarm-default-advertise-addr,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Bridge *string `json:"bridge,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ClusterAdvertise *string `json:"cluster-advertise,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Iptables *bool `json:"iptables,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	MaxConcurrentDownloads *int64 `json:"max-concurrent-downloads,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Features struct {
	} `json:"features,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	IpForward *bool `json:"ip-forward,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	StorageOpts []any `json:"storage-opts,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	FixedCidr *string `json:"fixed-cidr,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Containerd *string `json:"containerd,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	RegistryMirrors []string `json:"registry-mirrors,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ExecRoot *string `json:"exec-root,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Init *bool `json:"init,omitempty"`
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	Pidfile *string `json:"pidfile,omitempty"`
}
