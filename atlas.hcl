data "external_schema" "ent" {
  program = [
    "go", "run", "-mod=mod",
    "entgo.io/ent/cmd/ent", "schema",
    "./internal/data/ent/schema",
    "--dialect", "postgres",
    "--version", "16",
  ]
}

env "local" {
  src  = data.external_schema.ent.url
  dev = "docker://postgres/16/test?search_path=public"
  migration {
    dir = "file://internal/data/ent/migrate/migrations"
  }
}
