// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "naming_prefix" {
  description = "Prefix for the provisioned resources."
  type        = string
  default     = "platform"
}

variable "environment" {
  description = "Environment in which the resource should be provisioned like dev, qa, prod etc."
  type        = string
  default     = "dev"
}

variable "environment_number" {
  description = "The environment count for the respective environment. Defaults to 000. Increments in value of 1"
  default     = "000"
}

variable "region" {
  description = "AWS Region in which the infra needs to be provisioned"
  default     = "us-east-2"
}

variable "resource_number" {
  description = "The resource count for the respective resource. Defaults to 000. Increments in value of 1"
  default     = "000"
}

variable "resource_names_map" {
  description = "A map of key to resource_name that will be used by tf-aws-wrapper_module-iam_policy module to generate resource names. If using this wrapper module to create multiple policies within a role, you should adjust this default value to ensure uniqueness across your policies."
  type = map(object(
    {
      name       = string
      max_length = optional(number, 60)
    }
  ))
  default = {
    iam_policy = {
      name       = "plcy"
      max_length = 60
    }
  }
}

variable "policy" {
  description = "A single valid JSON document containing a policy."
  type        = string
}

variable "role_name" {
  description = "Name of the IAM role that should be associated with this policy. This value does not create an attachment but does contribute to tags that are required on the Policy."
  type        = string
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "A map of tags to add to the resources created by the module."
}
