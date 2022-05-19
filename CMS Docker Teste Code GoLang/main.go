package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	finance "github.com/piquette/finance-go"
	_ "github.com/piquette/finance-go/etf"
	_ "github.com/piquette/finance-go/options"
	"github.com/piquette/finance-go/quote"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go mod init github.com/ChrisMarSilva/cms.golang.tnb.yahoo
// go mod tidy

// go get github.com/piquette/finance-go
// https://piquette.io/projects/finance-go/

// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.tnb.yahoo"
// go run main.go
// go run .
// docker-compose up -d --build
// docker-compose down

func main() {

	// for {

	db := GetDatabase()

	log.Println("")
	log.Println("YahooFinance.INI")
	var start time.Time = time.Now()

	var wg sync.WaitGroup
	wg.Add(4)

	// go processarAtivo("ACAO-FULL", db, &wg)
	// go processarAtivo("FII-FULL", db, &wg)
	// go processarAtivo("ETF-FULL", db, &wg)
	// go processarAtivo("BDR-FULL", db, &wg)

	go processarAtivo("ACAO", db, &wg)
	go processarAtivo("FII", db, &wg)
	go processarAtivo("ETF", db, &wg)
	go processarAtivo("BDR", db, &wg)

	wg.Wait()

	log.Println("YahooFinance.FIM:", time.Since(start))
	log.Println("")

	// 	sqlDB, err := db.DB()
	// 	if err != nil {
	// 		log.Println("YahooFinance.ERRO.BD", err)
	// 	}
	// 	sqlDB.Close()

	// 	time.Sleep(30 * time.Minute) // time.Sleep(time.Duration(30) * time.Minute)
	// }

	// go func() {
	// 	for now := range time.Tick(time.Minute) {
	// 		fmt.Println(now, statusUpdate())
	// 	}
	// }()

}

func GetDatabase() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Silent // Error // Warn // Info
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	newConfig := &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		QueryFields:                              true,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		DryRun:                                   false,
	}

	// LOCALHOST
	dsn := "root:senha@tcp(localhost:3306)/database?parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), newConfig)
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	sqlDB.SetMaxIdleConns(100)                // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	sqlDB.SetMaxOpenConns(200)                // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	sqlDB.SetConnMaxIdleTime(time.Minute * 5) // SetConnMaxIdleTime define a quantidade máxima de tempo que uma conexão pode ficar inativa.
	sqlDB.SetConnMaxLifetime(time.Minute * 5) // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	return db
}

func processarAtivo(tipo string, db *gorm.DB, wg *sync.WaitGroup) {

	defer wg.Done()

	var start time.Time
	var sql string = ""

	if tipo == "ACAO" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'ACAO' AS Tipo, A.SITUACAO AS Situacao FROM TBEMPRESA_ATIVO A WHERE A.SITUACAO = 'A' AND ( A.CODIGO IN ('IBOV', 'IDIV', 'IBXX', 'SMLL') OR EXISTS( SELECT 1 FROM TBLANCAMENTO O WHERE O.IDATIVO  = A.ID ) OR EXISTS( SELECT 1 FROM TBUSUARIO_ACOMP_ATIVO O WHERE O.IDATIVO  = A.ID ) ) `
	} else if tipo == "ACAO-FULL" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'ACAO' AS Tipo, A.SITUACAO AS Situacao FROM TBEMPRESA_ATIVO A WHERE A.SITUACAO = 'A' `
	} else if tipo == "FII" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'FII'  AS Tipo, A.SITUACAO AS Situacao FROM TBFII_FUNDOIMOB A WHERE A.SITUACAO = 'A' AND ( EXISTS( SELECT 1 FROM TBFII_LANCAMENTO O WHERE O.IDFUNDO  = A.ID ) OR EXISTS( SELECT 1 FROM TBUSUARIO_ACOMP_FUNDO O WHERE O.IDFUNDO  = A.ID ) ) `
	} else if tipo == "FII-FULL" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'FII'  AS Tipo, A.SITUACAO AS Situacao FROM TBFII_FUNDOIMOB A WHERE A.SITUACAO = 'A'  `
	} else if tipo == "ETF" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'ETF' AS Tipo, A.SITUACAO AS Situacao FROM TBETF_INDICE A WHERE A.SITUACAO = 'A' AND ( EXISTS( SELECT 1 FROM TBETF_LANCAMENTO O WHERE O.IDINDICE = A.ID ) OR EXISTS( SELECT 1 FROM TBUSUARIO_ACOMP_INDICE O WHERE O.IDINDICE  = A.ID ) ) `
	} else if tipo == "ETF-FULL" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'ETF' AS Tipo, A.SITUACAO AS Situacao FROM TBETF_INDICE A WHERE A.SITUACAO = 'A'  `
	} else if tipo == "BDR" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'BDR' AS Tipo, A.SITUACAO AS Situacao FROM TBBDR_EMPRESA A WHERE A.SITUACAO = 'A' AND ( EXISTS( SELECT 1 FROM TBBDR_LANCAMENTO O WHERE O.IDBDR = A.ID ) OR EXISTS( SELECT 1 FROM TBUSUARIO_ACOMP_BDR O WHERE O.IDBDR  = A.ID ) )  `
	} else if tipo == "BDR-FULL" {
		sql = ` SELECT A.ID AS ID, A.CODIGO AS Codigo, 'BDR' AS Tipo, A.SITUACAO AS Situacao FROM TBBDR_EMPRESA A WHERE A.SITUACAO = 'A'  `
	}

	var rows []AtivoLocal
	err := db.Raw(sql).Scan(&rows).Error
	if err != nil {
		log.Println("YahooFinance."+tipo+".QUERY.", err) // err.Error()
		return
	}

	codigos := []string{}
	for _, row := range rows {
		codigo := row.Codigo + ".SA"
		if codigo == "IBOV" || codigo == "IBOV.SA" {
			codigo = "%5EBVSP"
		}
		codigos = append(codigos, codigo)
	}

	start = time.Now()
	var wgLocalQuote sync.WaitGroup
	if len(codigos) > 0 {
		iter := quote.List(codigos)
		for iter.Next() {
			q := iter.Quote()
			wgLocalQuote.Add(1)
			go func(q *finance.Quote, rows []AtivoLocal, ww *sync.WaitGroup) {
				defer ww.Done()
				codigo := strings.Replace(q.Symbol, ".SA", "", 1)
				index := bucarIndiceAtivos(rows, codigo)
				rows[index].Cotacao = q.RegularMarketPrice
				rows[index].Variacao = q.RegularMarketChangePercent
				rows[index].Anterior = q.RegularMarketPreviousClose
			}(q, rows, &wgLocalQuote)
		}
	}
	wgLocalQuote.Wait()
	log.Println("YahooFinance."+tipo+".GetList("+strconv.Itoa(len(codigos))+").Tempo:", time.Since(start))

	if tipo == "ACAO" || tipo == "ACAO-FULL" {
		sql = "UPDATE TBEMPRESA_ATIVOCOTACAO SET DATA = ?, VLRPRECOFECHAMENTO = ?, VLRPRECOANTERIOR = ?, VLRVARIACAO = ?, DATAHORAALTERACO = ? WHERE IDATIVO = ?"
	} else if tipo == "FII" || tipo == "FII-FULL" {
		sql = "UPDATE TBFII_FUNDOIMOB_COTACAO SET DATA = ?, VLRPRECOFECHAMENTO = ?, VLRPRECOANTERIOR = ?, VLRVARIACAO = ?, DATAHORAALTERACO = ? WHERE IDFUNDO = ?"
	} else if tipo == "ETF" || tipo == "ETF-FULL" {
		sql = "UPDATE TBETF_INDICE_COTACAO SET DATA = ?, VLRPRECOFECHAMENTO = ?, VLRPRECOANTERIOR = ?, VLRVARIACAO = ?, DATAHORAALTERACO = ? WHERE IDINDICE = ?"
	} else if tipo == "BDR" || tipo == "BDR-FULL" {
		sql = "UPDATE TBBDR_EMPRESA_COTACAO SET DATA = ?, VLRPRECOFECHAMENTO = ?, VLRPRECOANTERIOR = ?, VLRVARIACAO = ?, DATAHORAALTERACO = ? WHERE IDBDR = ?"
	}

	tx := db.Session(&gorm.Session{PrepareStmt: true})
	dataAtual := PegarDataAtual()
	dataHoraAtual := PegarDataHoraAtual()

	start = time.Now()
	// var wgLocalUpd sync.WaitGroup
	for _, row := range rows {
		// wgLocalUpd.Add(1)
		// go func(tipo string, sql string, data string, cotacao float64, anterior float64, variacao float64, datahora string, id int64, ww *sync.WaitGroup) {
		// go func(tipo string, sql string, data string, cotacao float64, anterior float64, variacao float64, datahora string, id int64, ww *sync.WaitGroup) {
		//	defer ww.Done()
		if row.Cotacao > 0.0 {
			err = tx.Exec(sql, dataAtual, row.Cotacao, row.Anterior, row.Variacao, dataHoraAtual, row.ID).Error
			// err = tx.Exec(sql, data, cotacao, anterior, variacao, datahora, id).Error
			if err != nil {
				log.Println("YahooFinance."+tipo+".UpdList.Erro:", err)
				return // continue
			}
		}
		//}(tipo, sql, dataAtual, row.Cotacao, row.Anterior, row.Variacao, dataHoraAtual, row.ID, &wgLocalUpd)
	}
	// wgLocalUpd.Wait()
	log.Println("YahooFinance."+tipo+".UpdList("+strconv.Itoa(len(codigos))+").Tempo:", time.Since(start))

}

func bucarIndiceAtivos(ativos []AtivoLocal, codigo string) int {
	for i, _ := range ativos {
		if ativos[i].Codigo == codigo || ativos[i].Codigo == codigo+".SA" {
			return i
		}
	}
	return -1
}

func PegarDataAtual() string {
	t := time.Now()
	var strDtAtual string = fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	return strDtAtual
}

func PegarDataHoraAtual() string {
	t := time.Now()
	var strDtHrAtual string = fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return strDtHrAtual
}

type AtivoLocal struct {
	ID       int64
	Codigo   string
	Tipo     string
	Situacao string
	Cotacao  float64 `default0:"0.0"`
	Variacao float64 `default0:"0.0"`
	Anterior float64 `default0:"0.0"`
}

func (row AtivoLocal) ToString() string {
	return fmt.Sprintf("ID: %s; Codigo: %s; Tipo: %s; Situacao: %s", row.ID, row.Codigo, row.Tipo, row.Situacao)
}
