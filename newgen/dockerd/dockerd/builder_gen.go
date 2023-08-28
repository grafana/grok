package dockerd

import (
	"encoding/json"

	"github.com/grafana/grok/newgen/dockerd/types"
)

type Option func(builder *Builder) error

type Builder struct {
	internal *types.Dockerd
}
func New(options ...Option) (Builder, error) {
	resource := &types.Dockerd{}
	builder := &Builder{internal: resource}

	for _, opt := range append(defaults(), options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}
// MarshalJSON implements the encoding/json.Marshaler interface.
//
// This method can be used to render the resource as JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalJSON() ([]byte, error) {
	return json.Marshal(builder.internal)
}

// MarshalIndentJSON renders the resource as indented JSON
// which your configuration management tool of choice can then feed into
// Grafana.
func (builder *Builder) MarshalIndentJSON() ([]byte, error) {
	return json.MarshalIndent(builder.internal, "", "  ")
}

func (builder *Builder) Internal() *types.Dockerd {
	return builder.internal
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func LogOpts(logOpts struct {
	CacheMaxSize string `json:"cache-max-size,omitempty"`
	CacheCompress string `json:"cache-compress,omitempty"`
	Env string `json:"env,omitempty"`
	Labels string `json:"labels,omitempty"`
	MaxFile string `json:"max-file,omitempty"`
	MaxSize string `json:"max-size,omitempty"`
	CacheDisabled string `json:"cache-disabled,omitempty"`
	CacheMaxFile string `json:"cache-max-file,omitempty"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.LogOpts = logOpts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Tlsverify(tlsverify bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tlsverify = tlsverify

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func FixedCidr(fixedCidr string) Option {
	return func(builder *Builder) error {
		
		builder.internal.FixedCidr = fixedCidr

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ClusterAdvertise(clusterAdvertise string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ClusterAdvertise = clusterAdvertise

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Debug(debug bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Debug = debug

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Ip(ip string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Ip = ip

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Init(init bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Init = init

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Tlscacert(tlscacert string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tlscacert = tlscacert

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultGatewayV6(defaultGatewayV6 string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultGatewayV6 = defaultGatewayV6

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func CgroupParent(cgroupParent string) Option {
	return func(builder *Builder) error {
		
		builder.internal.CgroupParent = cgroupParent

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Tls(tls bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tls = tls

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Runtimes(runtimes struct {
	CcRuntime struct {
	Path string `json:"path,omitempty"`
} `json:"cc-runtime,omitempty"`
	Custom struct {
	RuntimeArgs []string `json:"runtimeArgs,omitempty"`
	Path string `json:"path,omitempty"`
} `json:"custom,omitempty"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.Runtimes = runtimes

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultUlimits(defaultUlimits struct {
	Nofile struct {
	Hard int64 `json:"Hard,omitempty"`
	Name string `json:"Name,omitempty"`
	Soft int64 `json:"Soft,omitempty"`
} `json:"nofile,omitempty"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultUlimits = defaultUlimits

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Iptables(iptables bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Iptables = iptables

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Tlscert(tlscert string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tlscert = tlscert

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ShutdownTimeout(shutdownTimeout int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.ShutdownTimeout = shutdownTimeout

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Ip6tables(ip6tables bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Ip6tables = ip6tables

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ApiCorsHeader(apiCorsHeader string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ApiCorsHeader = apiCorsHeader

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Pidfile(pidfile string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Pidfile = pidfile

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Features(features struct {
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.Features = features

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Group(group string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Group = group

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultCgroupnsMode(defaultCgroupnsMode enum here) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultCgroupnsMode = defaultCgroupnsMode

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ClusterStoreOpts(clusterStoreOpts struct {
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.ClusterStoreOpts = clusterStoreOpts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func RawLogs(rawLogs bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.RawLogs = rawLogs

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DataRoot(dataRoot string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DataRoot = dataRoot

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Hosts(hosts []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Hosts = hosts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func LiveRestore(liveRestore bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.LiveRestore = liveRestore

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func FixedCidrV6(fixedCidrV6 string) Option {
	return func(builder *Builder) error {
		
		builder.internal.FixedCidrV6 = fixedCidrV6

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ContainerdPluginNamespace(containerdPluginNamespace string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ContainerdPluginNamespace = containerdPluginNamespace

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Ipv6(ipv6 bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Ipv6 = ipv6

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ExecRoot(execRoot string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ExecRoot = execRoot

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultAddressPools(defaultAddressPools []struct {
	Size int64 `json:"size,omitempty"`
	Base string `json:"base,omitempty"`
}) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultAddressPools = defaultAddressPools

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func OomScoreAdjust(oomScoreAdjust int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.OomScoreAdjust = oomScoreAdjust

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func SwarmDefaultAdvertiseAddr(swarmDefaultAdvertiseAddr string) Option {
	return func(builder *Builder) error {
		
		builder.internal.SwarmDefaultAdvertiseAddr = swarmDefaultAdvertiseAddr

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Bridge(bridge string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Bridge = bridge

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func RegistryMirrors(registryMirrors []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.RegistryMirrors = registryMirrors

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DnsSearch(dnsSearch []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DnsSearch = dnsSearch

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ClusterStore(clusterStore string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ClusterStore = clusterStore

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func AuthorizationPlugins(authorizationPlugins []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.AuthorizationPlugins = authorizationPlugins

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ExecOpts(execOpts []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ExecOpts = execOpts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func MaxDownloadAttempts(maxDownloadAttempts int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.MaxDownloadAttempts = maxDownloadAttempts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultGateway(defaultGateway string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultGateway = defaultGateway

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func SelinuxEnabled(selinuxEnabled bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.SelinuxEnabled = selinuxEnabled

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func UsernsRemap(usernsRemap string) Option {
	return func(builder *Builder) error {
		
		builder.internal.UsernsRemap = usernsRemap

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Bip(bip string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Bip = bip

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func LogLevel(logLevel string) Option {
	return func(builder *Builder) error {
		
		builder.internal.LogLevel = logLevel

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Experimental(experimental bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Experimental = experimental

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func NoNewPrivileges(noNewPrivileges bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.NoNewPrivileges = noNewPrivileges

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultIpcMode(defaultIpcMode enum here) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultIpcMode = defaultIpcMode

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func IpMasq(ipMasq bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.IpMasq = ipMasq

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func IpForward(ipForward bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.IpForward = ipForward

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Labels(labels []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Labels = labels

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func SeccompProfile(seccompProfile string) Option {
	return func(builder *Builder) error {
		
		builder.internal.SeccompProfile = seccompProfile

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Containerd(containerd string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Containerd = containerd

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Icc(icc bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.Icc = icc

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Mtu(mtu int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.Mtu = mtu

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func MaxConcurrentDownloads(maxConcurrentDownloads int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.MaxConcurrentDownloads = maxConcurrentDownloads

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Dns(dns []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.Dns = dns

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DnsOpts(dnsOpts []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DnsOpts = dnsOpts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func UserlandProxyPath(userlandProxyPath string) Option {
	return func(builder *Builder) error {
		
		builder.internal.UserlandProxyPath = userlandProxyPath

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultRuntime(defaultRuntime string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultRuntime = defaultRuntime

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func NodeGenericResources(nodeGenericResources []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.NodeGenericResources = nodeGenericResources

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func DefaultShmSize(defaultShmSize string) Option {
	return func(builder *Builder) error {
		
		builder.internal.DefaultShmSize = defaultShmSize

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func Tlskey(tlskey string) Option {
	return func(builder *Builder) error {
		
		builder.internal.Tlskey = tlskey

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func InsecureRegistries(insecureRegistries []string) Option {
	return func(builder *Builder) error {
		
		builder.internal.InsecureRegistries = insecureRegistries

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func LogDriver(logDriver string) Option {
	return func(builder *Builder) error {
		
		builder.internal.LogDriver = logDriver

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func UserlandProxy(userlandProxy bool) Option {
	return func(builder *Builder) error {
		
		builder.internal.UserlandProxy = userlandProxy

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func StorageDriver(storageDriver string) Option {
	return func(builder *Builder) error {
		
		builder.internal.StorageDriver = storageDriver

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func AllowNondistributableArtifacts(allowNondistributableArtifacts []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.AllowNondistributableArtifacts = allowNondistributableArtifacts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func InitPath(initPath string) Option {
	return func(builder *Builder) error {
		
		builder.internal.InitPath = initPath

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func ContainerdNamespace(containerdNamespace string) Option {
	return func(builder *Builder) error {
		
		builder.internal.ContainerdNamespace = containerdNamespace

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func StorageOpts(storageOpts []any) Option {
	return func(builder *Builder) error {
		
		builder.internal.StorageOpts = storageOpts

		return nil
	}
}
// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon
func MaxConcurrentUploads(maxConcurrentUploads int64) Option {
	return func(builder *Builder) error {
		
		builder.internal.MaxConcurrentUploads = maxConcurrentUploads

		return nil
	}
}

func defaults() []Option {
return []Option{
}
}
