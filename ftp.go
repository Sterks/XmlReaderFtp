package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"time"
	//"common"
)

func main() {
	c, err := ftp.Dial("ftp.zakupki.gov.ru:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	err = c.Login("free", "free")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	path := "/fcs_regions/Moskva/notifications"
	from, _ := time.Parse("2006-01-02 15:04:05", "2019-03-01 00:00:00")
	to, _ := time.Parse("2006-01-02 15:04:05", "2019-04-10 00:00:00")
	var files []*ftp.Entry
    var exiles []ExtendedEntry

	r := listFiles(c, path, files, exiles, from, to)
	i := 0
	for _, value := range r {
		fmt.Println(value)
		i = i + 1
	}
	fmt.Println(i)
}
	//	dateChange, _ := time.Parse("2006-01-02 15:04:05", "2019-04-04 00:00:00")
	//	if value.Time.Before(dateChange) || value.Time.Equal(dateChange) {


			//fmt.Println(value.Time, " меньше ", dateChange)
			//fmt.Println(to_mb(to_size(int64(value.Size))))
			//fmt.Println(fmt.Sprintf("%.2f", float64(value.Size)/1024/1024), "Mb")
		//}
	//}

type Entry struct {
	Name   string
	Target string // target of symbolic link
	Type   EntryType
	Size   uint64
	Time   time.Time
	Fullpath string
}
type EntryType int

type ExtendedEntry struct {
	Entry
	Fullpath string
}

func listFiles(conn *ftp.ServerConn, path string, files []*ftp.Entry, extf []ExtendedEntry, from time.Time, to time.Time) []*ftp.Entry {
	// Выполнил список
	list, err := conn.List(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range list {
		if value.Type == 0 {
			if value.Time.After(from) && value.Time.Before(to) {
				//fullpath := path + "/" + value.Name
				files = append(files, value)
			}
		} else {
			l := path + "/" +value.Name
			files = listFiles(conn, l, files, extf, from, to)
		}
	}
	return files
}


