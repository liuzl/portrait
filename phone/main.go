package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/golang/glog"
	"github.com/jszwec/csvutil"
	"github.com/liuzl/goutil"
	gopilosa "github.com/pilosa/go-pilosa"
	"github.com/pilosa/pdk"
)

var (
	pilosaHost  = flag.String("pilosa", "localhost:10101", "the pilosa host")
	pilosaIndex = flag.String("index", "phone", "the pilosa index")
	addr        = flag.String("addr", ":8080", "bind address")
	input       = flag.String("i", "../data/ecommerce_users-20181124.csv", "raw data")
)

type fRecord struct {
	Number string `csv:"msisdn"`
	Domain string `csv:"domain"`
}

func main() {
	flag.Parse()
	defer glog.Flush()

	go func() {
		fmt.Println(http.ListenAndServe(*addr, nil))
	}()

	schema := gopilosa.NewSchema()
	index := schema.Index(*pilosaIndex)
	pdk.NewRankedField(index, "domain", 10000)
	pdk.NewRankedField(index, "usertype", 100)
	pdk.NewRankedField(index, "day", 10000)
	pdk.NewRankedField(index, "mday", 10000)
	pdk.NewRankedField(index, "month", 13)
	pdk.NewRankedField(index, "year", 100)

	indexer, err := pdk.SetupPilosa([]string{*pilosaHost}, *pilosaIndex, schema, uint(1000000))
	if err != nil {
		glog.Fatal(err)
	}
	defer indexer.Close()
	count, err := goutil.FileLineCount(*input)
	if err != nil {
		glog.Fatal(err)
	}
	file, err := os.Open(*input)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		glog.Fatal(err)
	}
	bar := pb.StartNew(count - 1)
	for {
		var item fRecord
		if err := dec.Decode(&item); err == io.EOF {
			break
		} else if err != nil {
			glog.Fatal(err)
		}
		indexer.AddColumn("domain", item.Number, item.Domain)
		bar.Increment()
	}
	bar.FinishPrint("done!")
}
