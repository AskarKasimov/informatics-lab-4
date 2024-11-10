package pkg

import "sigs.k8s.io/yaml"

func FirstAdditionalTask(input string) string {
	res, err := yaml.YAMLToJSON([]byte(input))
	if err != nil {
		panic(err)
	}
	return string(res)
}
