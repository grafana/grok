// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
export interface Dockerd {
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	usernsRemap?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	seccompProfile?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ip6tables?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	allowNondistributableArtifacts?: any[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	apiCorsHeader?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	dnsSearch?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	cgroupParent?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	dataRoot?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	debug?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ip?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	maxDownloadAttempts?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	storageDriver?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	group?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	labels?: any[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	bip?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	userlandProxy?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	insecureRegistries?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultShmSize?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	liveRestore?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultIpcMode?: "shareable" | "private";
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	containerdNamespace?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	logLevel?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	icc?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	logOpts?: {
		cacheMaxSize?: string;
		cacheCompress?: string;
		env?: string;
		labels?: string;
		maxFile?: string;
		maxSize?: string;
		cacheDisabled?: string;
		cacheMaxFile?: string;
	};
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	authorizationPlugins?: any[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	dnsOpts?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	tlsverify?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	logDriver?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultCgroupnsMode?: "private" | "host";
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultGateway?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	userlandProxyPath?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	tlscacert?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	clusterStoreOpts?: {
	
	};
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	execOpts?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	runtimes?: {
		custom?: {
			runtimeArgs?: string[];
			path?: string;
		};
		ccRuntime?: {
			path?: string;
		};
	};
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	containerdPluginNamespace?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ipMasq?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ipv6?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	oomScoreAdjust?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	tlskey?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	dns?: any[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	initPath?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	tlscert?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	clusterStore?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultAddressPools?: {
		size?: number;
		base?: string;
	}[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	hosts?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultUlimits?: {
		nofile?: {
			name?: string;
			soft?: number;
			hard?: number;
		};
	};
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	fixedCidrV6?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	tls?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	noNewPrivileges?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	nodeGenericResources?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	experimental?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	selinuxEnabled?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	rawLogs?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	shutdownTimeout?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	maxConcurrentUploads?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	mtu?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultRuntime?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	defaultGatewayV6?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	swarmDefaultAdvertiseAddr?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	bridge?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	clusterAdvertise?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	iptables?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	maxConcurrentDownloads?: number;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	features?: {
	
	};
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	ipForward?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	storageOpts?: any[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	fixedCidr?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	containerd?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	registryMirrors?: string[];
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	execRoot?: string;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	init?: boolean;
	// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
	pidfile?: string;
}

