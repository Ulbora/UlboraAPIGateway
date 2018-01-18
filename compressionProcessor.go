/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

func processResponse(resp *http.Response) ([]byte, error) {
	var respbody []byte
	var bdyErr error
	//fmt.Print("Content-Encoding header: ")
	//fmt.Println(resp.Header.Get("Content-Encoding"))
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		//fmt.Println("found body to be gzip")
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Print("gzip error: ")
			fmt.Println(err)
		}
		defer gz.Close()
		resp.Header.Del("Content-Encoding")
		respbody, bdyErr = ioutil.ReadAll(gz)
	case "deflate":
		//fmt.Println("found body to be deflate")
		fz := flate.NewReader(resp.Body)
		defer fz.Close()
		resp.Header.Del("Content-Encoding")
		respbody, bdyErr = ioutil.ReadAll(fz)
	default:
		respbody, bdyErr = ioutil.ReadAll(resp.Body)
	}
	return respbody, bdyErr
}
