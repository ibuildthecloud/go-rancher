package client

const (
	REGISTRY_TYPE = "registry"
)

type Registry struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    ServerAddress string `json:"serverAddress,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type RegistryCollection struct {
	Collection
	Data []Registry `json:"data,omitempty"`
}

type RegistryClient struct {
	rancherClient *RancherClient
}

type RegistryOperations interface {
	List(opts *ListOpts) (*RegistryCollection, error)
	Create(opts *Registry) (*Registry, error)
	Update(existing *Registry, updates interface{}) (*Registry, error)
	ById(id string) (*Registry, error)
	Delete(container *Registry) error
    
    ActionActivate (*Registry) (*StoragePool, error)
    
    
    ActionCreate (*Registry) (*StoragePool, error)
    
    
    ActionDeactivate (*Registry) (*StoragePool, error)
    
    
    ActionPurge (*Registry) (*StoragePool, error)
    
    
    ActionRemove (*Registry) (*StoragePool, error)
    
    
    ActionRestore (*Registry) (*StoragePool, error)
    
    
    ActionUpdate (*Registry) (*StoragePool, error)
    
}

func newRegistryClient(rancherClient *RancherClient) *RegistryClient {
	return &RegistryClient{
		rancherClient: rancherClient,
	}
}

func (c *RegistryClient) Create(container *Registry) (*Registry, error) {
	resp := &Registry{}
	err := c.rancherClient.doCreate(REGISTRY_TYPE, container, resp)
	return resp, err
}

func (c *RegistryClient) Update(existing *Registry, updates interface{}) (*Registry, error) {
	resp := &Registry{}
	err := c.rancherClient.doUpdate(REGISTRY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RegistryClient) List(opts *ListOpts) (*RegistryCollection, error) {
	resp := &RegistryCollection{}
	err := c.rancherClient.doList(REGISTRY_TYPE, opts, resp)
	return resp, err
}

func (c *RegistryClient) ById(id string) (*Registry, error) {
	resp := &Registry{}
	err := c.rancherClient.doById(REGISTRY_TYPE, id, resp)
	return resp, err
}

func (c *RegistryClient) Delete(container *Registry) error {
	return c.rancherClient.doResourceDelete(REGISTRY_TYPE, &container.Resource)
}
    
func (c *RegistryClient) ActionActivate (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "activate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionCreate (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "create", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionDeactivate (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "deactivate", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionPurge (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "purge", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionRemove (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "remove", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionRestore (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "restore", &resource.Resource, nil, resp)
    
	return resp, err
}
    
func (c *RegistryClient) ActionUpdate (resource *Registry) (*StoragePool, error) {
    
	resp := &StoragePool{}
    
	err := c.rancherClient.doAction(REGISTRY_TYPE, "update", &resource.Resource, nil, resp)
    
	return resp, err
}
