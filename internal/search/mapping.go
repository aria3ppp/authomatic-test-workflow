package search

import "strings"

var (
	seriesMappings = strings.NewReader(`{
		"mappings": {
			"properties": {
				"id": { "type": "keyword", "index": false },
				"title": { "type": "text" },
				"descriptions": { "type": "text" },
				"date_started": { "type": "date", "index": false },
				"date_ended": { "type": "date", "index": false },
				"contributed_by": { "type": "keyword", "index": false },
				"contributed_at": { "type": "date", "index": false },
				"invalidation": { "type": "keyword", "index": false }
			}
		}
	}`)

	movieMappings = strings.NewReader(`{
		"mappings": {
			"properties": {
				"id": { "type": "keyword", "index": false },
				"title": { "type": "text" },
				"descriptions": { "type": "text" },
				"date_released": { "type": "date", "index": false },
				"duration": { "type": "short", "index": false },
				"contributed_by": { "type": "keyword", "index": false },
				"contributed_at": { "type": "date", "index": false },
				"invalidation": { "type": "keyword", "index": false }
			}
		}
	}`)
)
