---
page_title: "dynatrace_request_attribute Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# Resource `dynatrace_request_attribute`





## Schema

### Optional

- **aggregation** (String)
- **confidential** (Boolean)
- **data_sources** (Block List) (see [below for nested schema](#nestedblock--data_sources))
- **data_type** (String)
- **enabled** (Boolean)
- **id** (String) The ID of this resource.
- **metadata** (Block List, Max: 1) (see [below for nested schema](#nestedblock--metadata))
- **name** (String)
- **normalization** (String)
- **skip_personal_data_masking** (Boolean)

<a id="nestedblock--data_sources"></a>
### Nested Schema for `data_sources`

Optional:

- **capturing_and_storage_location** (String)
- **cics_sdkmethod_node_condition** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--cics_sdkmethod_node_condition))
- **enabled** (Boolean)
- **iib_label_method_node_condition** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--iib_label_method_node_condition))
- **iib_method_node_condition** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--iib_method_node_condition))
- **iib_node_type** (String)
- **methods** (Block List) (see [below for nested schema](#nestedblock--data_sources--methods))
- **parameter_name** (String)
- **scope** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--scope))
- **session_attribute_technology** (String)
- **source** (String)
- **technology** (String)
- **value_processing** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--value_processing))

<a id="nestedblock--data_sources--cics_sdkmethod_node_condition"></a>
### Nested Schema for `data_sources.cics_sdkmethod_node_condition`

Optional:

- **negate** (Boolean)
- **operator** (String)
- **value** (String)


<a id="nestedblock--data_sources--iib_label_method_node_condition"></a>
### Nested Schema for `data_sources.iib_label_method_node_condition`

Optional:

- **negate** (Boolean)
- **operator** (String)
- **value** (String)


<a id="nestedblock--data_sources--iib_method_node_condition"></a>
### Nested Schema for `data_sources.iib_method_node_condition`

Optional:

- **negate** (Boolean)
- **operator** (String)
- **value** (String)


<a id="nestedblock--data_sources--methods"></a>
### Nested Schema for `data_sources.methods`

Optional:

- **argument_index** (Number)
- **capture** (String)
- **deep_object_access** (String)
- **method** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--methods--method))

<a id="nestedblock--data_sources--methods--method"></a>
### Nested Schema for `data_sources.methods.method`

Optional:

- **argument_types** (List of String)
- **class_name** (String)
- **file_name** (String)
- **file_name_matcher** (String)
- **method_name** (String)
- **modifiers** (List of String)
- **return_type** (String)
- **visibility** (String)



<a id="nestedblock--data_sources--scope"></a>
### Nested Schema for `data_sources.scope`

Optional:

- **host_group** (String)
- **process_group** (String)
- **service_technology** (String)
- **tag_of_process_group** (String)


<a id="nestedblock--data_sources--value_processing"></a>
### Nested Schema for `data_sources.value_processing`

Optional:

- **extract_substring** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--value_processing--extract_substring))
- **split_at** (String)
- **trim** (Boolean)
- **value_condition** (Block List, Max: 1) (see [below for nested schema](#nestedblock--data_sources--value_processing--value_condition))
- **value_extractor_regex** (String)

<a id="nestedblock--data_sources--value_processing--extract_substring"></a>
### Nested Schema for `data_sources.value_processing.extract_substring`

Optional:

- **delimiter** (String)
- **end_delimiter** (String)
- **position** (String)


<a id="nestedblock--data_sources--value_processing--value_condition"></a>
### Nested Schema for `data_sources.value_processing.value_condition`

Optional:

- **negate** (Boolean)
- **operator** (String)
- **value** (String)




<a id="nestedblock--metadata"></a>
### Nested Schema for `metadata`

Optional:

- **cluster_version** (String)
- **configuration_versions** (List of Number)
- **current_configuration_versions** (List of String)


