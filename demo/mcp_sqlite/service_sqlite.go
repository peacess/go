package mcpsqlite

import mcp_golang "github.com/metoro-io/mcp-golang"

type ServiceSqlite struct {
}

func (c *ServiceSqlite) Register(server *mcp_golang.Server) {
	{
		type ExecuteSqlArguments struct {
			Sql string `json:"Sql" jsonschema:"required,description=SQL to execute"`
		}
		server.RegisterTool("execute_sql", "Execute a SQL on the SQLite database", func(args ExecuteSqlArguments) (res *mcp_golang.ToolResponse, err error) {
			res = &mcp_golang.ToolResponse{}
			err = nil
			println(args.Sql)
			return
		})
	}
	{
		type ListTablesArguments struct {
			Sql string `json:"Sql" jsonschema:"required,description=SQL to execute"`
		}
		server.RegisterTool("list_tables", "list all tables in the SQLite database", func(args ListTablesArguments) (res *mcp_golang.ToolResponse, err error) {
			res = &mcp_golang.ToolResponse{}
			err = nil
			println(args.Sql)
			return
		})
	}
	// {
	// 	type ListTablesArguments struct {
	// 		Sql string `json:"Sql" jsonschema:"required,description=SQL to execute"`
	// 	}
	// 	server.RegisterTool("create_table", "create table in the SQLite database", func(args ListTablesArguments) (res *mcp_golang.ToolResponse, err error) {
	// 		res = &mcp_golang.ToolResponse{}
	// 		err = nil
	// 		println(args.Sql)
	// 		return
	// 	})
	// }
}
