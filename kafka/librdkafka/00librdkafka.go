package librdkafka

/**
 * Copyright 2019 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* This is a dummy subpackage which is imported by ../kafka/build_static.go
 * to make sure this directory is not filtered out (due to no Go files)
 * when using go modules.
 * Something needs to be exported from this package and used by librdkafka.go
 * to avoid unused-errors. */

// #include "rdkafka.h"
import "C"

// LibrdkafkaVersion is a numeric representation of the librdkafka version
// linked statically with this build.
var LibrdkafkaVersion int = C.RD_KAFKA_VERSION
