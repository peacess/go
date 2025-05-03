package mcpsqlite

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"unsafe"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type ServiceSqlite struct {
	db *sql.DB
}

func (c *ServiceSqlite) Register(server *mcp_golang.Server) (err error) {
	err = c.Init()
	if err != nil {
		return
	}
	{
		type ExecuteSqlArguments struct {
			Sql string `json:"Sql" jsonschema:"required,description=SQL to execute"`
		}
		server.RegisterTool("execute_sql", "Execute a SQL on the SQLite database", func(args ExecuteSqlArguments) (res *mcp_golang.ToolResponse, errRe error) {
			res = &mcp_golang.ToolResponse{}
			errRe = nil
			result, errRe := c.db.Exec(args.Sql)
			if errRe != nil {
				return
			}
			var _ = result

			println(args.Sql)
			return
		})
	}
	{
		type QuerySqlArguments struct {
			Sql string `json:"Sql" jsonschema:"required,description=SQL to query"`
		}
		server.RegisterTool("query_sql", "Execute a SQL on the SQLite database", func(args QuerySqlArguments) (res *mcp_golang.ToolResponse, errRe error) {
			res = &mcp_golang.ToolResponse{}
			errRe = nil
			rows, errRe := c.db.Query(args.Sql)
			if errRe != nil {
				return
			}

			colums, errRe := rows.Columns()
			if errRe != nil {
				return
			}
			values := make([]any, len(colums))
			valuePtrs := make([]any, len(colums))
			for i := range len(colums) {
				valuePtrs[i] = &values[i]
			}
			var data []map[string]any
			for rows.Next() {
				if errRe = rows.Scan(valuePtrs...); errRe != nil {
					return
				}
				row := make(map[string]any)
				for i, col := range colums {
					row[col] = values[i]
				}
				data = append(data, row)
			}
			bs, errRe := json.Marshal(data)
			if errRe != nil {
				return
			}
			var dataString string
			{
				p := unsafe.SliceData(bs)
				dataString = unsafe.String(p, len(bs))
			}
			res.Content = append(res.Content, mcp_golang.NewTextContent(dataString))

			defer rows.Close()

			println(args.Sql)
			return
		})
	}
	{
		type ListTablesArguments struct {
			Sql string `json:"Sql" jsonschema:"required,description=SQL to execute"`
		}
		server.RegisterTool("list_tables", "list all tables in the SQLite database", func(args ListTablesArguments) (res *mcp_golang.ToolResponse, errRe error) {
			res = &mcp_golang.ToolResponse{}
			errRe = nil
			rows, errRe := c.db.Query("SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;")
			if errRe != nil {
				return
			}
			defer rows.Close()
			names := make([]string, 0, 64)
			for rows.Next() {
				var tableName string
				if err := rows.Scan(&tableName); err != nil {
					break
				}
				names = append(names, tableName)
			}

			res.Content = append(res.Content, mcp_golang.NewTextContent(strings.Join(names, ", ")))

			if errRe = rows.Err(); errRe != nil {
				return
			}
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
	return
}

func (c *ServiceSqlite) Init() (errRe error) {
	if c.db == nil {
		c.db, errRe = sql.Open("sqlite3", "./mcp_sqlite.db")
		if errRe != nil {
			return
		}
		createTableSQL := `CREATE TABLE IF NOT EXISTS users (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" TEXT,
			"age" INTEGER
		);`
		_, errRe = c.db.Exec(createTableSQL)
		if errRe != nil {
			return
		}
	} else {
		errRe = errors.New("the db is not init")
	}
	return
}

func (c *ServiceSqlite) Uninit() {
	if c.db != nil {
		c.db.Close()
		c.db = nil
	}
}
