---
page_title: "dynatrace_custom_service Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# Resource `dynatrace_custom_service`





## Schema

### Required

- **enabled** (Boolean) Custom service enabled/disabled
- **name** (String) The name of the custom service, displayed in the UI
- **queue_entry_point** (Boolean) The queue entry point flag. Set to `true` for custom messaging services
- **technology** (String) Matcher applying to the file name (ENDS_WITH, EQUALS or STARTS_WITH). Default value is ENDS_WITH (if applicable)

### Optional

- **id** (String) The ID of this resource.
- **order** (String) The order string. Sorting custom services alphabetically by their order string determines their relative ordering. Typically this is managed by Dynatrace internally and will not be present in GET responses
- **process_groups** (List of String) The list of process groups the custom service should belong to
- **queue_entry_point_type** (String) The queue entry point type (IBM_MQ, JMS, KAFKA, MSMQ or RABBIT_MQ)
- **rule** (Block List) The list of rules defining the custom service (see [below for nested schema](#nestedblock--rule))

<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Required:

- **enabled** (Boolean) Rule enabled/disabled
- **method** (Block List, Min: 1) methods to instrument (see [below for nested schema](#nestedblock--rule--method))

Optional:

- **annotations** (List of String) Additional annotations filter of the rule. Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented. Not applicable to PHP
- **class** (Block List, Max: 1) The fully qualified class or interface to instrument (or a substring if matching to a string). Required for Java and .NET custom services. Not applicable to PHP (see [below for nested schema](#nestedblock--rule--class))
- **file** (Block List, Max: 1) The PHP file containing the class or methods to instrument. Required for PHP custom service. Not applicable to Java and .NET (see [below for nested schema](#nestedblock--rule--file))

Read-only:

- **id** (String) The ID of the detection rule

<a id="nestedblock--rule--method"></a>
### Nested Schema for `rule.method`

Required:

- **name** (String) The method to instrument
- **returns** (String) Fully qualified type the method returns

Optional:

- **arguments** (List of String) Fully qualified types of argument the method expects

Read-only:

- **id** (String) The ID of the method rule


<a id="nestedblock--rule--class"></a>
### Nested Schema for `rule.class`

Required:

- **name** (String) The full name of the file / the name to match the file name with

Optional:

- **match** (String) Matcher applying to the class name (ENDS_WITH, EQUALS or STARTS_WITH). STARTS_WITH can only be used if there is at least one annotation defined. Default value is EQUALS


<a id="nestedblock--rule--file"></a>
### Nested Schema for `rule.file`

Required:

- **name** (String) The full name of the file / the name to match the file name with

Optional:

- **match** (String) Matcher applying to the file name (ENDS_WITH, EQUALS or STARTS_WITH). Default value is ENDS_WITH (if applicable)


