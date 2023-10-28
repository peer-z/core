/*
 * Copyright 2020 PeerGum
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

var chessService = Service{
	Info: ServiceInfo{
		Name:        "Chess",
		Path:        "/chess",
		Id:          chessServiceID,
		Version:     0x0100, // version 1.00
		address:     chessServiceAddress,
		Description: "A basic chess game",
		flags:       0,
	},
}
