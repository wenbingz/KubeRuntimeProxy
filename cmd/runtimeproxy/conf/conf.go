package conf

const (
	DefaultRuntimeProxyEndpoint             = "/var/run/runtimeproxy/runtimeproxy.sock"
	DefaultContainerdRuntimeServiceEndpoint = "/var/run/containerd/containerd.sock"
	DefaultContainerdImageServiceEndpoint   = "/var/run/containerd/containerd.sock"

	BackendRuntimeModeContainerd = "Containerd"
	BackendRuntimeModeDocker     = "Docker"
	DefaultBackendRuntimeMode    = BackendRuntimeModeContainerd
)

var (
	RuntimeProxyEndpoint         string
	RemoteRuntimeServiceEndpoint string
	RemoteImageServiceEndpoint   string

	BackendRuntimeMode string
)
