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

func DnsOpts(dnsOpts []string) Option {
	return func(builder *Builder) error {

		builder.internal.DnsOpts = dnsOpts

		return nil
	}
}

func Tls(tls bool) Option {
	return func(builder *Builder) error {

		builder.internal.Tls = &tls

		return nil
	}
}

func Dns(dns []any) Option {
	return func(builder *Builder) error {

		builder.internal.Dns = dns

		return nil
	}
}

func DefaultIpcMode(defaultIpcMode types.DefaultIpcModeEnum) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultIpcMode = &defaultIpcMode

		return nil
	}
}

func MaxConcurrentDownloads(maxConcurrentDownloads int64) Option {
	return func(builder *Builder) error {

		builder.internal.MaxConcurrentDownloads = &maxConcurrentDownloads

		return nil
	}
}

func LogDriver(logDriver string) Option {
	return func(builder *Builder) error {

		builder.internal.LogDriver = &logDriver

		return nil
	}
}

func DefaultAddressPools(defaultAddressPools []struct {
	Size *int64  `json:"size,omitempty"`
	Base *string `json:"base,omitempty"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultAddressPools = defaultAddressPools

		return nil
	}
}

func ContainerdPluginNamespace(containerdPluginNamespace string) Option {
	return func(builder *Builder) error {

		builder.internal.ContainerdPluginNamespace = &containerdPluginNamespace

		return nil
	}
}

func LogOpts(logOpts struct {
	CacheMaxSize  *string `json:"cache-max-size,omitempty"`
	CacheCompress *string `json:"cache-compress,omitempty"`
	Env           *string `json:"env,omitempty"`
	Labels        *string `json:"labels,omitempty"`
	MaxFile       *string `json:"max-file,omitempty"`
	MaxSize       *string `json:"max-size,omitempty"`
	CacheDisabled *string `json:"cache-disabled,omitempty"`
	CacheMaxFile  *string `json:"cache-max-file,omitempty"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.LogOpts = logOpts

		return nil
	}
}

func Mtu(mtu int64) Option {
	return func(builder *Builder) error {

		builder.internal.Mtu = &mtu

		return nil
	}
}

func ContainerdNamespace(containerdNamespace string) Option {
	return func(builder *Builder) error {

		builder.internal.ContainerdNamespace = &containerdNamespace

		return nil
	}
}

func Init(init bool) Option {
	return func(builder *Builder) error {

		builder.internal.Init = &init

		return nil
	}
}

func Ip6tables(ip6tables bool) Option {
	return func(builder *Builder) error {

		builder.internal.Ip6tables = &ip6tables

		return nil
	}
}

func Group(group string) Option {
	return func(builder *Builder) error {

		builder.internal.Group = &group

		return nil
	}
}

func Tlscert(tlscert string) Option {
	return func(builder *Builder) error {

		builder.internal.Tlscert = &tlscert

		return nil
	}
}

func Iptables(iptables bool) Option {
	return func(builder *Builder) error {

		builder.internal.Iptables = &iptables

		return nil
	}
}

func Bip(bip string) Option {
	return func(builder *Builder) error {

		builder.internal.Bip = &bip

		return nil
	}
}

func UserlandProxy(userlandProxy bool) Option {
	return func(builder *Builder) error {

		builder.internal.UserlandProxy = &userlandProxy

		return nil
	}
}

func DnsSearch(dnsSearch []string) Option {
	return func(builder *Builder) error {

		builder.internal.DnsSearch = dnsSearch

		return nil
	}
}

func Experimental(experimental bool) Option {
	return func(builder *Builder) error {

		builder.internal.Experimental = &experimental

		return nil
	}
}

func CgroupParent(cgroupParent string) Option {
	return func(builder *Builder) error {

		builder.internal.CgroupParent = &cgroupParent

		return nil
	}
}

func Hosts(hosts []string) Option {
	return func(builder *Builder) error {

		builder.internal.Hosts = hosts

		return nil
	}
}

func StorageDriver(storageDriver string) Option {
	return func(builder *Builder) error {

		builder.internal.StorageDriver = &storageDriver

		return nil
	}
}

func DataRoot(dataRoot string) Option {
	return func(builder *Builder) error {

		builder.internal.DataRoot = &dataRoot

		return nil
	}
}

func DefaultUlimits(defaultUlimits struct {
	Nofile struct {
		Hard *int64  `json:"Hard,omitempty"`
		Name *string `json:"Name,omitempty"`
		Soft *int64  `json:"Soft,omitempty"`
	} `json:"nofile,omitempty"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultUlimits = defaultUlimits

		return nil
	}
}

func ExecOpts(execOpts []string) Option {
	return func(builder *Builder) error {

		builder.internal.ExecOpts = execOpts

		return nil
	}
}

func UserlandProxyPath(userlandProxyPath string) Option {
	return func(builder *Builder) error {

		builder.internal.UserlandProxyPath = &userlandProxyPath

		return nil
	}
}

func RegistryMirrors(registryMirrors []string) Option {
	return func(builder *Builder) error {

		builder.internal.RegistryMirrors = registryMirrors

		return nil
	}
}

func IpForward(ipForward bool) Option {
	return func(builder *Builder) error {

		builder.internal.IpForward = &ipForward

		return nil
	}
}

func Debug(debug bool) Option {
	return func(builder *Builder) error {

		builder.internal.Debug = &debug

		return nil
	}
}

func Labels(labels []any) Option {
	return func(builder *Builder) error {

		builder.internal.Labels = labels

		return nil
	}
}

func Features(features struct {
}) Option {
	return func(builder *Builder) error {

		builder.internal.Features = features

		return nil
	}
}

func SelinuxEnabled(selinuxEnabled bool) Option {
	return func(builder *Builder) error {

		builder.internal.SelinuxEnabled = &selinuxEnabled

		return nil
	}
}

func Ipv6(ipv6 bool) Option {
	return func(builder *Builder) error {

		builder.internal.Ipv6 = &ipv6

		return nil
	}
}

func ShutdownTimeout(shutdownTimeout int64) Option {
	return func(builder *Builder) error {

		builder.internal.ShutdownTimeout = &shutdownTimeout

		return nil
	}
}

func Containerd(containerd string) Option {
	return func(builder *Builder) error {

		builder.internal.Containerd = &containerd

		return nil
	}
}

func Pidfile(pidfile string) Option {
	return func(builder *Builder) error {

		builder.internal.Pidfile = &pidfile

		return nil
	}
}

func RawLogs(rawLogs bool) Option {
	return func(builder *Builder) error {

		builder.internal.RawLogs = &rawLogs

		return nil
	}
}

func SwarmDefaultAdvertiseAddr(swarmDefaultAdvertiseAddr string) Option {
	return func(builder *Builder) error {

		builder.internal.SwarmDefaultAdvertiseAddr = &swarmDefaultAdvertiseAddr

		return nil
	}
}

func FixedCidr(fixedCidr string) Option {
	return func(builder *Builder) error {

		builder.internal.FixedCidr = &fixedCidr

		return nil
	}
}

func MaxDownloadAttempts(maxDownloadAttempts int64) Option {
	return func(builder *Builder) error {

		builder.internal.MaxDownloadAttempts = &maxDownloadAttempts

		return nil
	}
}

func UsernsRemap(usernsRemap string) Option {
	return func(builder *Builder) error {

		builder.internal.UsernsRemap = &usernsRemap

		return nil
	}
}

func ClusterAdvertise(clusterAdvertise string) Option {
	return func(builder *Builder) error {

		builder.internal.ClusterAdvertise = &clusterAdvertise

		return nil
	}
}

func ClusterStore(clusterStore string) Option {
	return func(builder *Builder) error {

		builder.internal.ClusterStore = &clusterStore

		return nil
	}
}

func DefaultShmSize(defaultShmSize string) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultShmSize = &defaultShmSize

		return nil
	}
}

func MaxConcurrentUploads(maxConcurrentUploads int64) Option {
	return func(builder *Builder) error {

		builder.internal.MaxConcurrentUploads = &maxConcurrentUploads

		return nil
	}
}

func DefaultGateway(defaultGateway string) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultGateway = &defaultGateway

		return nil
	}
}

func SeccompProfile(seccompProfile string) Option {
	return func(builder *Builder) error {

		builder.internal.SeccompProfile = &seccompProfile

		return nil
	}
}

func AuthorizationPlugins(authorizationPlugins []any) Option {
	return func(builder *Builder) error {

		builder.internal.AuthorizationPlugins = authorizationPlugins

		return nil
	}
}

func ClusterStoreOpts(clusterStoreOpts struct {
}) Option {
	return func(builder *Builder) error {

		builder.internal.ClusterStoreOpts = clusterStoreOpts

		return nil
	}
}

func AllowNondistributableArtifacts(allowNondistributableArtifacts []any) Option {
	return func(builder *Builder) error {

		builder.internal.AllowNondistributableArtifacts = allowNondistributableArtifacts

		return nil
	}
}

func StorageOpts(storageOpts []any) Option {
	return func(builder *Builder) error {

		builder.internal.StorageOpts = storageOpts

		return nil
	}
}

func InsecureRegistries(insecureRegistries []string) Option {
	return func(builder *Builder) error {

		builder.internal.InsecureRegistries = insecureRegistries

		return nil
	}
}

func ExecRoot(execRoot string) Option {
	return func(builder *Builder) error {

		builder.internal.ExecRoot = &execRoot

		return nil
	}
}

func Icc(icc bool) Option {
	return func(builder *Builder) error {

		builder.internal.Icc = &icc

		return nil
	}
}

func Tlskey(tlskey string) Option {
	return func(builder *Builder) error {

		builder.internal.Tlskey = &tlskey

		return nil
	}
}

func DefaultCgroupnsMode(defaultCgroupnsMode types.DefaultCgroupnsModeEnum) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultCgroupnsMode = &defaultCgroupnsMode

		return nil
	}
}

func LogLevel(logLevel string) Option {
	return func(builder *Builder) error {

		builder.internal.LogLevel = &logLevel

		return nil
	}
}

func IpMasq(ipMasq bool) Option {
	return func(builder *Builder) error {

		builder.internal.IpMasq = &ipMasq

		return nil
	}
}

func DefaultRuntime(defaultRuntime string) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultRuntime = &defaultRuntime

		return nil
	}
}

func ApiCorsHeader(apiCorsHeader string) Option {
	return func(builder *Builder) error {

		builder.internal.ApiCorsHeader = &apiCorsHeader

		return nil
	}
}

func DefaultGatewayV6(defaultGatewayV6 string) Option {
	return func(builder *Builder) error {

		builder.internal.DefaultGatewayV6 = &defaultGatewayV6

		return nil
	}
}

func NoNewPrivileges(noNewPrivileges bool) Option {
	return func(builder *Builder) error {

		builder.internal.NoNewPrivileges = &noNewPrivileges

		return nil
	}
}

func Bridge(bridge string) Option {
	return func(builder *Builder) error {

		builder.internal.Bridge = &bridge

		return nil
	}
}

func OomScoreAdjust(oomScoreAdjust int64) Option {
	return func(builder *Builder) error {

		builder.internal.OomScoreAdjust = &oomScoreAdjust

		return nil
	}
}

func InitPath(initPath string) Option {
	return func(builder *Builder) error {

		builder.internal.InitPath = &initPath

		return nil
	}
}

func Tlscacert(tlscacert string) Option {
	return func(builder *Builder) error {

		builder.internal.Tlscacert = &tlscacert

		return nil
	}
}

func FixedCidrV6(fixedCidrV6 string) Option {
	return func(builder *Builder) error {

		builder.internal.FixedCidrV6 = &fixedCidrV6

		return nil
	}
}

func Ip(ip string) Option {
	return func(builder *Builder) error {

		builder.internal.Ip = &ip

		return nil
	}
}

func Runtimes(runtimes struct {
	CcRuntime struct {
		Path *string `json:"path,omitempty"`
	} `json:"cc-runtime,omitempty"`
	Custom struct {
		Path        *string  `json:"path,omitempty"`
		RuntimeArgs []string `json:"runtimeArgs,omitempty"`
	} `json:"custom,omitempty"`
}) Option {
	return func(builder *Builder) error {

		builder.internal.Runtimes = runtimes

		return nil
	}
}

func Tlsverify(tlsverify bool) Option {
	return func(builder *Builder) error {

		builder.internal.Tlsverify = &tlsverify

		return nil
	}
}

func LiveRestore(liveRestore bool) Option {
	return func(builder *Builder) error {

		builder.internal.LiveRestore = &liveRestore

		return nil
	}
}

func NodeGenericResources(nodeGenericResources []string) Option {
	return func(builder *Builder) error {

		builder.internal.NodeGenericResources = nodeGenericResources

		return nil
	}
}

func defaults() []Option {
	return []Option{}
}
