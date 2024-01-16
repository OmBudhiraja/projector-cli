package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config *Config
	data   *Data
}

func (p *Projector) GetValue(key string) (string, bool) {

	cur := p.config.Pwd
	prev := ""

	out := ""
	found := false

	for cur != prev {

		if dir, ok := p.data.Projector[cur]; ok {
			if val, ok := dir[key]; ok {
				out = val
				found = true
				break
			}
		}
		prev = cur
		cur = path.Dir(cur)
	}

	return out, found
}

func (p *Projector) GetValueAll() map[string]string {

	cur := p.config.Pwd
	prev := ""

	out := make(map[string]string)

	for cur != prev {

		if dir, ok := p.data.Projector[cur]; ok {
			for k, v := range dir {
				if _, exists := out[k]; !exists {
					out[k] = v
				}
			}
		}
		prev = cur
		cur = path.Dir(cur)
	}

	return out
}

func (p *Projector) AddValue(key, value string) {
	if _, ok := p.data.Projector[p.config.Pwd]; !ok {
		p.data.Projector[p.config.Pwd] = make(map[string]string)
	}
	p.data.Projector[p.config.Pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	if _, ok := p.data.Projector[p.config.Pwd]; ok {
		delete(p.data.Projector[p.config.Pwd], key)
	}
}

func (p *Projector) Save() error {
	dir := filepath.Dir(p.config.Config)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	jsonString, err := json.Marshal(p.data)

	if err != nil {
		return err
	}

	os.WriteFile(p.config.Config, jsonString, 0755)

	return nil
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		config: config,
		data:   &Data{Projector: make(map[string]map[string]string)},
	}
}

func NewProjector(config *Config) *Projector {
	if _, err := os.Stat(config.Config); err == nil {
		contents, err := os.ReadFile(config.Config)

		if err != nil {
			return defaultProjector(config)
		} else {
			var data Data
			err = json.Unmarshal(contents, &data)
			if err != nil {
				fmt.Println("error", err)
				return defaultProjector(config)
			}
			return &Projector{
				config: config,
				data:   &data,
			}
		}
	}
	return defaultProjector(config)
}
