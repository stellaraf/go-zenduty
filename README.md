![stellar](https://res.cloudinary.com/stellaraf/image/upload/v1604277355/stellar-logo-gradient.png?width=300)

[![Go Reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/go.stellar.af/go-zenduty) [![GitHub Tag](https://img.shields.io/github/v/tag/stellaraf/go-zenduty?style=for-the-badge&label=Version)](https://github.com/stellaraf/go-zenduty/tags)


## `go-zenduty`

A Go SDK for [Zenduty](https://www.zenduty.com).

> [!NOTE]
> This client is autogenerated using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the [Zenduty OpenAPI Spec](https://apidocs.zenduty.com).

### Known Issues

<details>

<summary><h4> 1. Empty Position Parameters in Spec</h4></summary>

The spec has been manually modified to resolve the following error:

```
error generating code: error creating operation definitions: path '/api/account/teams/{}/maintenance/' has 0 positional parameters, but spec has 1 declared
```

Nearly all paths following `/teams` contain a positional parameter of `{}`, but a named parameter is required. For example:

```diff
{
    "paths": {
-        "/api/account/teams/{}/escalation_policies/": {
+        "/api/account/teams/{team_id}/escalation_policies/": {
            "parameters": [
				{
					"name": "team_id",
					"in": "path",
					"description": "unique_id of the Team object",
					"schema": {
						"type": "string"
					},
					"required": true
				}
			]
        }
    }
}
```

The same was required for the `"/api/v2/account/teams/{team_id}/schedules/{}/overrides/` path, but with `schedule_id`.

</details>


![GitHub License](https://img.shields.io/github/license/stellaraf/go-zenduty?style=for-the-badge&color=black)
