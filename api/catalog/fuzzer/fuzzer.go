/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzzer

import (
	v1alpha2 "go.bytebuilders.dev/catalog/api/catalog/v1alpha1"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/randfill"
)

// Funcs returns the fuzzer functions for this api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []any {
	return []any{
		// v1alpha1
		func(s *v1alpha2.ElasticsearchBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.KafkaBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.MariaDBBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.MemcachedBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.MongoDBBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.MySQLBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.PerconaXtraDBBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.PgBouncerBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.PostgresBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.ProxySQLBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
		func(s *v1alpha2.RedisBinding, c randfill.Continue) {
			c.Fill(s) // fuzz self without calling this function again
		},
	}
}
