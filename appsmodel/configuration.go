package appsmodel

type Configuration struct {
	Port     int    `yaml:"port"`
	Title    string `yaml:"title"`
	Favicon  string `yaml:"favicon"`
	Database struct {
		Name     string `yaml:"name"`
		Server   string `yaml:"server"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Cookie struct {
		Persist  bool `yaml:"persist"`
		Secure   bool `yaml:"secure"`
		LifeTime int  `yaml:"lifetime"`
	} `yaml:"cookie"`
	Template struct {
		Cached  bool     `yaml:"cached"`
		Dir     string   `yaml:"dir"`
		Options []string `yaml:"options"`
	} `yaml:"template"`
	Application struct {
		PageDir    string `yaml:"pagedir"`
		ContentDir string `yaml:"contentdir"`
	} `yaml:"application"`
	Logging struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"logging"`
	ShowServerError bool `yaml:"showservererror"`
	HitTest         bool `yaml:"hittest"`
	ConfigPath      string
}
