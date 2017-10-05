package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	//"log"
	"fmt"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type ProjectsConfig struct {
	Projects []ProjectConfig `yaml:"projects"`
}

type ProjectConfig struct {
	Name                string              `yaml:"name"`
	QueuesDefaultConfig QueuesDefaultConfig `yaml:"queues_default"`
	Queues              []QueueConfig       `yaml:"queues"`
}

type QueuesDefaultConfig struct {
	NotifyBase      string `yaml:"notify_base"`
	NotifyTimeout   int    `yaml:"notify_timeout"`
	RetryTimes      int    `yaml:"retry_times"`
	RetryDuration   int    `yaml:"retry_duration"`
	BindingExchange string `yaml:"binding_exchange"`
}

type QueueConfig struct {
	QueueName       string   `yaml:"queue_name"`
	RoutingKey      []string `yaml:"routing_key"`
	NotifyPath      string   `yaml:"notify_path"`

	project *ProjectConfig
}


func loadConfig() {
	configFile, err := ioutil.ReadFile("./config.yaml")
	PanicOnError(err)

	projectsConfig := ProjectsConfig{}

	err = yaml.Unmarshal(configFile, &projectsConfig)
	PanicOnError(err)
	//log.Printf("find config: %v", projectsConfig)

	projects := projectsConfig.Projects
	for i, project := range projects {
		fmt.Println("find project:", project.Name)

		fmt.Println("default:",project.QueuesDefaultConfig)

		queues := projects[i].Queues
		for j, queue := range queues {
			fmt.Println("find queue:", queue)

			queues[j].project = &projects[i]

		}
	}

}

func main() {
	loadConfig()
}