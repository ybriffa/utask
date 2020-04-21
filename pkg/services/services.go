package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

var (
	servicesImported = make(map[string]*Service)
)

type Service struct {
	Name   string `json:"name"`
	Action Action `json:"action"`

	rawService json.RawMessage
}

type Action struct {
	Type              string          `json:"type"`
	BaseConfiguration string          `json:"base_configuration,omitempty"`
	Configuration     json.RawMessage `json:"configuration"`
	BaseOutput        json.RawMessage `json:"base_output"`
}

func (s *Service) Exec(stepName string, baseConfig json.RawMessage, config json.RawMessage, ctx interface{}) (interface{}, interface{}, map[string]string, error) {
	return nil, nil, nil, errors.New("NIY")
}

func (s *Service) ValidConfig(baseConfig json.RawMessage, config json.RawMessage) error {
	fmt.Println("baseConfig:", string(baseConfig))
	fmt.Println("config:", string(config))
	return nil
}

func (s *Service) Context(stepName string) interface{} {
	return nil
}

func (s *Service) MetadataSchema() json.RawMessage {
	return s.rawService
}

func LoadFromDirectory(directory string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		logrus.Warnf("Ignoring services directory %s: %s", directory, err)
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			if err := LoadFromDirectory(path.Join(directory, file.Name())); err != nil {
				return err
			}
			continue
		}

		if !strings.HasSuffix(file.Name(), ".yaml") || strings.HasPrefix(file.Name(), ".") {
			continue
		}

		content, err := ioutil.ReadFile(path.Join(directory, file.Name()))
		if err != nil {
			return err
		}

		var service Service
		if err = yaml.Unmarshal(content, &service); err != nil {
			return err
		}

		servicesImported[service.Name] = &service
	}

	return nil
}

func List() []string {
	var result = []string{}

	for k := range servicesImported {
		result = append(result, k)
	}
	return result
}

func Get(name string) (*Service, bool) {
	s, exists := servicesImported[name]
	return s, exists
}
