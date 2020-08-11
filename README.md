# gograb

Downloads go module dependencies one at a time.

This tool is needed when access to a private repos only allows a small number of concurrent connections.

This problem occurs during CI/CD pipelines. Golang efficiently pulls dependencies in parallel but this efficiency
can break private repos which only allow a small number of concurrent connections per host. Since gograb 
pulls dependencies one at a time, this should minimize the number of concurrent connections per host.
