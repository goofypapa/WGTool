package main

import (
	"log"
	"os"
	"io/ioutil"
)

var(
	root string
)

func main(){
	
	dir, err := os.Getwd();
    if err != nil {
        final();
	}
	
	root = dir;
	
	listFile( "" );
	success();
}

func listFile( dir string ){

	files, _ := ioutil.ReadDir( root + dir );

	for _, file := range files {
		if file.IsDir(){
			listFile( dir + "/" + file.Name() )
		} else {

		}
		log.Print( "------>", root + dir + "/" + file.Name() );
	}
}

func final(){
	log.Fatal( "------>")
}

func success(){
	log.Print( "convert success" );
}