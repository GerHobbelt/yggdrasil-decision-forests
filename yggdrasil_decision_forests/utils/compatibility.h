/*
 * Copyright 2022 Google LLC.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file contains copy of method that exist in the library we are using,
// but not in the version we are using.
//
// For each method, write what is the source library, and when it will be
// possible to use the source library.
//
#ifndef YGGDRASIL_DECISION_FORESTS_UTILS_COMPATIBILITY_H_
#define YGGDRASIL_DECISION_FORESTS_UTILS_COMPATIBILITY_H_

#include <stdint.h>

#include <cassert>
#include <optional>
#include <string>

namespace yggdrasil_decision_forests {
namespace utils {

// Name of the user.
inline std::optional<std::string> UserName() {
   // TODO: Platform specific implementation.
   return {};
}

#if defined(__GNUC__)
#define PREFETCH(addr) __builtin_prefetch(addr)
#else
#define PREFETCH(addr)
#endif

// Similar as std::accumulate introduced in c++20.
template <class Iter, class Result>
Result accumulate(Iter first, Iter last, Result acc) {
  for (; first != last; ++first) {
    acc = std::move(acc) + *first;
  }
  return acc;
}

}  // namespace utils
}  // namespace yggdrasil_decision_forests

#endif  // THIRD_PARTY_YGGDRASIL_DECISION_FORESTS_UTILS_COMPAT_H_
