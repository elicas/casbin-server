// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"errors"

	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

var errDriverName = errors.New("currently supported DriverName: file | mysql | postgres | mssql")

func newAdapter(driverName string, connectionString string, dbSpecified bool) (persist.Adapter, error) {
	var a persist.Adapter
	supportDriverNames := [...]string{"file", "mysql", "postgres", "mssql"}

	switch driverName {
	case "file":
		a = fileadapter.NewAdapter(connectionString)
	default:
		var support = false
		for _, driverName := range supportDriverNames {
			if driverName == driverName {
				support = true
				break
			}
		}
		if !support {
			return nil, errDriverName
		}

		var err error
		a, err = gormadapter.NewAdapter(driverName, connectionString, dbSpecified)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}
