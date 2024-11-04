# Copyright 2022 Google LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Test dataspec utilities for pandas."""

from absl.testing import absltest
import pandas as pd
import polars as pl

from ydf.dataset.io import pandas_io


class PandasIOTest(absltest.TestCase):

  def test_is_pandas(self):
    self.assertTrue(pandas_io.is_pandas_dataframe(pd.DataFrame()))
    self.assertFalse(pandas_io.is_pandas_dataframe({}))

  def test_polars_is_not_pandas(self):
    self.assertFalse(pandas_io.is_pandas_dataframe(pl.DataFrame()))


if __name__ == "__main__":
  absltest.main()
