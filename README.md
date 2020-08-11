# gograb

Downloads go module dependencies one at a time.

This tool is needed for when access to a private repo that only allows a small number of conncurrent connections.

This problem occurs during CI/CD pipelines. Golang efficiently pulls dependencies in parallel but this efficiency
can break private repos which only allow a small number of concurrent connections per client host. Since this
pull dependencies one at a time there will only be one client connection per host at a time.