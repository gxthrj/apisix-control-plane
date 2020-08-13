/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yml

type YmlModel interface {
	ToMem() string
}

type Gateway struct {
	Kind    *string  `json:"kind"`
	Name    *string  `json:"name"`
	Servers []Server `json:"servers"`
}

type Server struct {
	Port  *Port    `json:"port,omitempty"`
	Hosts []string `json:"host,omitempty"`
}

type Port struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

func (g *Gateway) ToMem() string {
	return "gateway"
}

type Rule struct {
	Kind     string   `json:"kind"`
	Name     string   `json:"name"`
	Hosts    []string `json:"hosts"`
	Gateways []string `json:"gateways"`
	HTTP     []HTTP   `json:"http"`
}
type RouteDestination struct {
	Port   int64  `json:"port"`
	Host   string `json:"host"`
	Subset string `json:"subset"`
	Weight int64  `json:"weight"`
}
type Route struct {
	Destination RouteDestination `json:"destination"`
}
type Label map[string]string

type Headers map[string]interface{}

type Match struct {
	Headers Headers `json:"headers"`
}
type HTTP struct {
	Route []Route `json:"route"`
	Label Label   `json:"label"`
	Match []Match `json:"match,omitempty"`
}

func (r *Rule) ToMem() string {
	return "Rule"
}

type Destination struct {
	Kind    string   `json:"kind"`
	Name    string   `json:"name"`
	Host    string   `json:"host"`
	Subsets []Subset `json:"subsets"`
}

type Subset struct {
	Name     string            `json:"name"`
	Ips      []string          `json:"ips,omitempty"`
	Selector map[string]string `json:"selector,omitempty"`
}

func (g *Destination) ToMem() string {
	return "destination"
}

type Plugin struct {
	Kind     string            `yaml:"kind"`
	Selector map[string]string `yaml:"selector"`
	Sets     []Set             `yaml:"sets"`
}

type Set struct {
	Name string                 `yaml:"name"`
	Conf map[string]interface{} `yaml:"conf,omitempty"`
}

func (g *Plugin) ToMem() string {
	return "plugin"
}
