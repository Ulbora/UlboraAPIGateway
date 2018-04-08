package cache

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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//CProxy cache proxy
type CProxy struct {
	Host string
}

//Item CacheItem
type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//ResponseValue ResponseValue
type ResponseValue struct {
	Success       bool   `json:"success"`
	Value         string `json:"value"`
	ServiceFailed bool   `json:"serviceFailed"`
}

//Response Response
type Response struct {
	Success bool `json:"success"`
}

// Set sets the cache
func (cp *CProxy) Set(i *Item) *Response {
	var rtn = new(Response)
	var sURL = cp.Host + "/rs/cache/set"
	aJSON, err := json.Marshal(i)
	if err != nil {
		fmt.Println("JSON parse err in cache: ")
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Cache set err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					fmt.Print("Decode cache err: ")
					fmt.Println(error)
					//log.Println(error.Error())
				}
			}
		}
	}
	return rtn
}

//Get get value
func (cp *CProxy) Get(key string) *ResponseValue {
	var rtn = new(ResponseValue)
	var sURL = cp.Host + "/rs/cache/get/" + key
	resp, err := http.Get(sURL)
	//fmt.Println(resp)
	if err != nil {
		log.Println(err.Error())
		rtn.ServiceFailed = true
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(&rtn)
		if error != nil {
			log.Println(error.Error())
		}
	}
	return rtn
}

//Delete delete value
func (cp *CProxy) Delete(key string) *Response {
	var rtn = new(Response)
	var sURL = cp.Host + "/rs/cache/delete/" + key
	//fmt.Print("sURL: ")
	//fmt.Println(sURL)
	req, rErr := http.NewRequest("DELETE", sURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Cache Service delete err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return rtn
}
