# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
Helper script to download license files for library dependencies.
Script relies on dependency spec json file to provide destination path to store
license files and list all the libraries and their corresponding url to license
file.
"""

import json
import os
import sys
from urlparse import urlparse
import wget


def _create_local_directory(path):
  try:
    os.mkdir(path)
  except OSError as err:
    print("OS error: {}".format(err))
    sys.exit(1)

def copy_third_party_licenses(dependency_spec):
  if not os.path.isfile(dependency_spec):
    print('dependency spec: {} not found'.format(dependency_spec))
    sys.exit(1)

  with open(dependency_spec, 'r') as f:
    dependencies = json.load(f)
    if not dependencies['target_path']:
      print("Invalid dependency spec. 'target_path' expected")

    _create_local_directory(dependencies['target_path'])

    for dependency in dependencies['libraries']:
      license_dest = '{}/{}'.format(dependencies['target_path'], dependency['library'])
      license_url = dependency['license_url']
      _create_local_directory(license_dest)
      wget.download(license_url, '{}/{}'.format(license_dest, os.path.split(urlparse(license_url).path)[-1]))

if __name__ == '__main__':
  if len(sys.argv) < 2:
    print('script expects path to the dependency spec file as argument')
    sys.exit(1)

  copy_third_party_licenses(sys.argv[1])

