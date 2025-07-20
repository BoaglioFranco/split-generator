# Split.io Impressions and events generator

## Configuration Options

The script is configured via a `.json` configuration file. The following options are supported:

| Option | Type | Description |
|---|---|---|
| `apiKey` | `string` | Your Split.io API key. |
| `isStaging` | `bool` | Whether to use the staging environment. |
| `flags` | `[]FlagConfig` | A list of feature flags to initialize. |

### `FlagConfig`

| Option | Type | Description |
|---|---|---|
| `name` | `string` | The name of the feature flag. |
| `impressions` | `int` | The number of impressions to generate for this flag. |
| `events` | `[]EventsConfig` | A list of events to associate with this flag. |

### `EventsConfig`

| Option | Type | Description |
|---|---|---|
| `eventType` | `string` | The type of the event. |
| `trafficType` | `string` | The traffic type for the event. |
| `treatments` | `map[string]EventValueSettings` | A map of treatments to their settings. |

### `EventValueSettings`

| Option | Type | Description |
|---|---|---|
| `value` | `*int` | The value of the event. |
| `count` | `*int` | The number of times to send the event. |
| `properties` | `map[string]interface{}` | A map of properties to associate with the event. |

## Example `config.json`

```json
{
  "apiKey": "YOUR_API_KEY",
  "isStaging": true,
  "flags": [
    {
      "name": "gonzalosm_test_flag",
      "impressions": 34,
      "events": [
        {
          "eventType": "page.load.time",
          "trafficType": "user",
          "treatments": {
            "on": { "value": 900 },
            "off": { "value": 200 }
          }
        },
        {
          "eventType": "task_creation",
          "trafficType": "user",
          "treatments": {
            "on": { "count": 2, "properties": { "i_am": "a property" } },
            "off": { "count": 1, "properties": { "i_am": "a different property" } }
          }
        }
      ]
    }
  ]
}
```
