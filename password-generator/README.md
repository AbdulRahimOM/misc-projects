API to generate passwords

Endpoint: "/"
    
Optional query parameters:

┌───────────────────┬───────────────────┬──────────────────┐
│ Query Parameters  │ Expected values   │   Default value  │
╞═══════════════════╪═══════════════════╪══════════════════╡
│     smallAlph     │    true, false    │   true           │
├───────────────────┼───────────────────┼──────────────────┤
│     capAlph       │    true, false    │   true           │
├───────────────────┼───────────────────┼──────────────────┤
│     num           │    true, false    │   true           │
├───────────────────┼───────────────────┼──────────────────┤
│     symbols       │    true/false     │   false          │
├───────────────────┼───────────────────┼──────────────────┤
│     minLen        │    integer        │   8              │
├───────────────────┼───────────────────┼──────────────────┤
│     maxLen        │    integer        │   minLen+10      │
└───────────────────┴───────────────────┴──────────────────┘

eg: url?smallAlph=true&capAlph=true&num=true&symbols=true&minLen=100