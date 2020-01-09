package output

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Load 加载输出器
func Load(confs map[string]string) []Outputer {
	var result []Outputer
	for name, ref := range confs {
		switch name {
		case "postgresql":
			conn, err := sql.Open("postgres", ref)
			if err != nil {
				log.Panic("加载postgresql输出器失败:", err)
			}
			err = conn.Ping()
			if err != nil {
				log.Panic("加载postgresql输出器失败:", err)
			}

			pgOutputer := PGOutputer{
				dsn: ref,
				db:  conn,
			}

			result = append(result, pgOutputer)
			log.Println("加载postgresql输出器，成功")
		case "stdout":
			result = append(result, STDOutOutputer{})
			log.Println("加载stdout输出器，成功")
		}
	}
	return result
}

// All 执行全部相关输出器Output函数
func All(outputers []Outputer, msg string) {
	for _, op := range outputers {
		err := op.Output(msg)
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
}
