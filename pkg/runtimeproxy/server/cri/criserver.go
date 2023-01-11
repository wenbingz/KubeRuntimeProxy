package cri

import (
	"CRTProxy/cmd/runtimeproxy/conf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"net"
	"time"
)

type RuntimeCriServer struct {
	backendRuntimeServiceClient runtimeapi.RuntimeServiceClient
	backendImageServiceClient   runtimeapi.ImageServiceClient
}

func (c *RuntimeCriServer) ListImages(ctx context.Context, request *runtimeapi.ListImagesRequest) (*runtimeapi.ListImagesResponse, error) {
	fmt.Println("intercepting listimage method")
	return c.backendImageServiceClient.ListImages(ctx, request)
}

func (c *RuntimeCriServer) ImageStatus(ctx context.Context, request *runtimeapi.ImageStatusRequest) (*runtimeapi.ImageStatusResponse, error) {
	fmt.Println("intercepting image status method")
	return c.backendImageServiceClient.ImageStatus(ctx, request)
}

func (c *RuntimeCriServer) PullImage(ctx context.Context, request *runtimeapi.PullImageRequest) (*runtimeapi.PullImageResponse, error) {
	fmt.Println("intercepting pull image method")
	return c.backendImageServiceClient.PullImage(ctx, request)
}

func (c *RuntimeCriServer) RemoveImage(ctx context.Context, request *runtimeapi.RemoveImageRequest) (*runtimeapi.RemoveImageResponse, error) {
	fmt.Println("intercepting remove image method")
	return c.backendImageServiceClient.RemoveImage(ctx, request)
}

func (c *RuntimeCriServer) ImageFsInfo(ctx context.Context, request *runtimeapi.ImageFsInfoRequest) (*runtimeapi.ImageFsInfoResponse, error) {
	fmt.Println("intercepting image fs info method")
	return c.backendImageServiceClient.ImageFsInfo(ctx, request)
}

func (c *RuntimeCriServer) Version(ctx context.Context, request *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	fmt.Println("intercept version method")
	return c.backendRuntimeServiceClient.Version(ctx, request)
}

func (c *RuntimeCriServer) RunPodSandbox(ctx context.Context, request *runtimeapi.RunPodSandboxRequest) (*runtimeapi.RunPodSandboxResponse, error) {
	fmt.Println("intercept runpodsandbox method")
	return c.backendRuntimeServiceClient.RunPodSandbox(ctx, request)
}

func (c *RuntimeCriServer) StopPodSandbox(ctx context.Context, request *runtimeapi.StopPodSandboxRequest) (*runtimeapi.StopPodSandboxResponse, error) {
	fmt.Println("intercept stoppodsandbox method")
	return c.backendRuntimeServiceClient.StopPodSandbox(ctx, request)
}

func (c *RuntimeCriServer) RemovePodSandbox(ctx context.Context, request *runtimeapi.RemovePodSandboxRequest) (*runtimeapi.RemovePodSandboxResponse, error) {
	fmt.Println("intercept removepodsandbox method")
	return c.backendRuntimeServiceClient.RemovePodSandbox(ctx, request)
}

func (c *RuntimeCriServer) PodSandboxStatus(ctx context.Context, request *runtimeapi.PodSandboxStatusRequest) (*runtimeapi.PodSandboxStatusResponse, error) {
	fmt.Println("intercept podsandboxstatus method")
	return c.backendRuntimeServiceClient.PodSandboxStatus(ctx, request)
}

func (c *RuntimeCriServer) ListPodSandbox(ctx context.Context, request *runtimeapi.ListPodSandboxRequest) (*runtimeapi.ListPodSandboxResponse, error) {
	fmt.Println("intercept ListPodSandbox method")
	return c.backendRuntimeServiceClient.ListPodSandbox(ctx, request)
}

func (c *RuntimeCriServer) CreateContainer(ctx context.Context, request *runtimeapi.CreateContainerRequest) (*runtimeapi.CreateContainerResponse, error) {
	fmt.Println("intercept CreateContainer method")
	return c.backendRuntimeServiceClient.CreateContainer(ctx, request)
}

func (c *RuntimeCriServer) StartContainer(ctx context.Context, request *runtimeapi.StartContainerRequest) (*runtimeapi.StartContainerResponse, error) {
	fmt.Println("intercept StartContainer method")
	return c.backendRuntimeServiceClient.StartContainer(ctx, request)
}

func (c *RuntimeCriServer) StopContainer(ctx context.Context, request *runtimeapi.StopContainerRequest) (*runtimeapi.StopContainerResponse, error) {
	fmt.Println("intercept StopContainer method")
	return c.backendRuntimeServiceClient.StopContainer(ctx, request)
}

func (c *RuntimeCriServer) RemoveContainer(ctx context.Context, request *runtimeapi.RemoveContainerRequest) (*runtimeapi.RemoveContainerResponse, error) {
	fmt.Println("intercept RemoveContainer method")
	return c.backendRuntimeServiceClient.RemoveContainer(ctx, request)
}

func (c *RuntimeCriServer) ListContainers(ctx context.Context, request *runtimeapi.ListContainersRequest) (*runtimeapi.ListContainersResponse, error) {
	fmt.Println("intercept ListContainers method")
	return c.backendRuntimeServiceClient.ListContainers(ctx, request)
}

func (c *RuntimeCriServer) ContainerStatus(ctx context.Context, request *runtimeapi.ContainerStatusRequest) (*runtimeapi.ContainerStatusResponse, error) {
	fmt.Println("intercept ContainerStatus method")
	return c.backendRuntimeServiceClient.ContainerStatus(ctx, request)
}

func (c *RuntimeCriServer) UpdateContainerResources(ctx context.Context, request *runtimeapi.UpdateContainerResourcesRequest) (*runtimeapi.UpdateContainerResourcesResponse, error) {
	fmt.Println("intercept UpdateContainerResources method")
	return c.backendRuntimeServiceClient.UpdateContainerResources(ctx, request)
}

func (c *RuntimeCriServer) ReopenContainerLog(ctx context.Context, request *runtimeapi.ReopenContainerLogRequest) (*runtimeapi.ReopenContainerLogResponse, error) {
	fmt.Println("intercept ReopenContainerLog method")
	return c.backendRuntimeServiceClient.ReopenContainerLog(ctx, request)
}

func (c *RuntimeCriServer) ExecSync(ctx context.Context, request *runtimeapi.ExecSyncRequest) (*runtimeapi.ExecSyncResponse, error) {
	fmt.Println("intercept ExecSync method")
	return c.backendRuntimeServiceClient.ExecSync(ctx, request)
}

func (c *RuntimeCriServer) Exec(ctx context.Context, request *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	fmt.Println("intercept Exec method")
	return c.backendRuntimeServiceClient.Exec(ctx, request)
}

func (c *RuntimeCriServer) Attach(ctx context.Context, request *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error) {
	fmt.Println("intercept Attach method")
	return c.backendRuntimeServiceClient.Attach(ctx, request)
}

func (c *RuntimeCriServer) PortForward(ctx context.Context, request *runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error) {
	fmt.Println("intercept PortForward method")
	return c.backendRuntimeServiceClient.PortForward(ctx, request)
}

func (c *RuntimeCriServer) ContainerStats(ctx context.Context, request *runtimeapi.ContainerStatsRequest) (*runtimeapi.ContainerStatsResponse, error) {
	fmt.Println("intercept ContainerStats method")
	return c.backendRuntimeServiceClient.ContainerStats(ctx, request)
}

func (c *RuntimeCriServer) ListContainerStats(ctx context.Context, request *runtimeapi.ListContainerStatsRequest) (*runtimeapi.ListContainerStatsResponse, error) {
	fmt.Println("intercept ListContainerStats method")
	return c.backendRuntimeServiceClient.ListContainerStats(ctx, request)
}

func (c *RuntimeCriServer) UpdateRuntimeConfig(ctx context.Context, request *runtimeapi.UpdateRuntimeConfigRequest) (*runtimeapi.UpdateRuntimeConfigResponse, error) {
	fmt.Println("intercept UpdateRuntimeConfig method")
	return c.backendRuntimeServiceClient.UpdateRuntimeConfig(ctx, request)
}

func (c *RuntimeCriServer) Status(ctx context.Context, request *runtimeapi.StatusRequest) (*runtimeapi.StatusResponse, error) {
	fmt.Println("intercept Status method")
	return c.backendRuntimeServiceClient.Status(ctx, request)
}

func (c *RuntimeCriServer) Name() string {
	return "RuntimeCriServer"
}

func (c *RuntimeCriServer) Run() error {
	if err := c.initBackendServer(conf.RemoteRuntimeServiceEndpoint, conf.RemoteImageServiceEndpoint); err != nil {
		return err
	}
	listener, err := net.Listen("unix", conf.RuntimeProxyEndpoint)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	runtimeapi.RegisterRuntimeServiceServer(grpcServer, c)
	runtimeapi.RegisterImageServiceServer(grpcServer, c)
	err = grpcServer.Serve(listener)
	return err
}

func dialer(ctx context.Context, addr string) (net.Conn, error) {
	return (&net.Dialer{}).DialContext(ctx, "unix", addr)
}
func (c *RuntimeCriServer) initBackendServer(runtimeSocketPath, imageSocketPath string) error {
	generateGrpcConn := func(socketPath string) (*grpc.ClientConn, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return grpc.DialContext(ctx, socketPath, grpc.WithInsecure(), grpc.WithContextDialer(dialer))
	}
	if conn, err := generateGrpcConn(runtimeSocketPath); err != nil {
		return err
	} else {
		c.backendRuntimeServiceClient = runtimeapi.NewRuntimeServiceClient(conn)
	}

	if conn, err := generateGrpcConn(imageSocketPath); err != nil {
		return err
	} else {
		c.backendImageServiceClient = runtimeapi.NewImageServiceClient(conn)
	}
	return nil
}
