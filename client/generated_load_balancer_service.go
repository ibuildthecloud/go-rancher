package client

const (
	LOAD_BALANCER_SERVICE_TYPE = "loadBalancerService"
)

type LoadBalancerService struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    DataVolumesFromService []string `json:"dataVolumesFromService,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    EnvironmentId string `json:"environmentId,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    LaunchConfig Container `json:"launchConfig,omitempty"`
    
    LoadBalancerConfig *LoadBalancerConfig `json:"loadBalancerConfig,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    Scale int `json:"scale,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type LoadBalancerServiceCollection struct {
	Collection
	Data []LoadBalancerService `json:"data,omitempty"`
}

type LoadBalancerServiceClient struct {
	rancherClient *RancherClient
}

type LoadBalancerServiceOperations interface {
	List(opts *ListOpts) (*LoadBalancerServiceCollection, error)
	Create(opts *LoadBalancerService) (*LoadBalancerService, error)
	Update(existing *LoadBalancerService, updates interface{}) (*LoadBalancerService, error)
	ById(id string) (*LoadBalancerService, error)
	Delete(container *LoadBalancerService) error
    
    ActionActivate (*LoadBalancerService) (*Service, error)
    
    
    ActionAddservicelink (*LoadBalancerService, *AddRemoveServiceLinkInput) (*Service, error)
    
    
    ActionCreate (*LoadBalancerService) (*Service, error)
    
    
    ActionDeactivate (*LoadBalancerService) (*Service, error)
    
    
    ActionRemove (*LoadBalancerService) (*Service, error)
    
    
    ActionRemoveservicelink (*LoadBalancerService, *AddRemoveServiceLinkInput) (*Service, error)
    
    
    ActionSetservicelinks (*LoadBalancerService, *SetServiceLinksInput) (*Service, error)
    
    
    ActionUpdate (*LoadBalancerService) (*Service, error)
    
}

func newLoadBalancerServiceClient(rancherClient *RancherClient) *LoadBalancerServiceClient {
	return &LoadBalancerServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *LoadBalancerServiceClient) Create(container *LoadBalancerService) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doCreate(LOAD_BALANCER_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) Update(existing *LoadBalancerService, updates interface{}) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doUpdate(LOAD_BALANCER_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) List(opts *ListOpts) (*LoadBalancerServiceCollection, error) {
	resp := &LoadBalancerServiceCollection{}
	err := c.rancherClient.doList(LOAD_BALANCER_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) ById(id string) (*LoadBalancerService, error) {
	resp := &LoadBalancerService{}
	err := c.rancherClient.doById(LOAD_BALANCER_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *LoadBalancerServiceClient) Delete(container *LoadBalancerService) error {
	return c.rancherClient.doResourceDelete(LOAD_BALANCER_SERVICE_TYPE, &container.Resource)
}
    
func (c *LoadBalancerServiceClient) ActionActivate (resource *LoadBalancerService) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "activate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionAddservicelink (resource *LoadBalancerService, input *AddRemoveServiceLinkInput) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "addservicelink", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionCreate (resource *LoadBalancerService) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionDeactivate (resource *LoadBalancerService) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "deactivate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionRemove (resource *LoadBalancerService) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionRemoveservicelink (resource *LoadBalancerService, input *AddRemoveServiceLinkInput) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "removeservicelink", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionSetservicelinks (resource *LoadBalancerService, input *SetServiceLinksInput) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "setservicelinks", &resource.Resource, input, resp)
    
	return resp, err
}
    
func (c *LoadBalancerServiceClient) ActionUpdate (resource *LoadBalancerService) (*Service, error) {
    
	resp := &Service{}
    
	err := c.rancherClient.doAction(LOAD_BALANCER_SERVICE_TYPE, "update", &resource.Resource, nil, resp)
    
	return resp, err
}
