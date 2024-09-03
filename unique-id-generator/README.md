API to generate unique id

Endpoint: "/"
    
Optional query parameters:

┌───────────────────┬───────────────────┬──────────────────┐
│ Query Parameters  │ Expected values   │   Default value  │
╞═══════════════════╪═══════════════════╪══════════════════╡
│     len           │    integer        │   36             │
├───────────────────┼───────────────────┼──────────────────┤
│     allowHyphen   │    true, false    │   true           │
└───────────────────┴───────────────────┴──────────────────┘

eg: <project_base_url>?len=20&allowHyphen=true