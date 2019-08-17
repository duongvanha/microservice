package obj

type BuddyYml []BuddyAction

type BuddyAction struct {
	Pipeline         string    `yaml:"pipeline"`
	TriggerMode      string    `yaml:"trigger_mode"`
	RefName          string    `yaml:"ref_name"`
	RefType          string    `yaml:"ref_type"`
	TriggerCondition string    `yaml:"trigger_condition"`
	Actions          []Actions `yaml:"actions"`
}

type Actions struct {
	Action           string `yaml:"action"`
	Type             string `yaml:"type"`
	Login            string `yaml:"login"`
	Password         string `yaml:"password"`
	Disabled         bool   `yaml:"disabled"`
	DockerImageTag   string `yaml:"docker_image_tag"`
	DockerfilePath   string `yaml:"dockerfile_path"`
	ContextPath      string `yaml:"context_path"`
	Repository       string `yaml:"repository"`
	TriggerCondition string `yaml:"trigger_condition"`
}
