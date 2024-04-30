package deployments

import "strings"

func convertTargetToMap(target string) map[string]string {
	res := make(map[string]string)

	split := strings.Split(target, ",")
	for _, s := range split {
		secondSplit := strings.Split(s, "=")
		if len(secondSplit) != 2 {
			continue
		}

		key, val := secondSplit[0], secondSplit[1]
		res[key] = val
	}

	return res
}
