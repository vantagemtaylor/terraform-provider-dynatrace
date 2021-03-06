---
page_title: "dynatrace_notification Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# Resource `dynatrace_notification`





## Schema

### Optional

- **ansible_tower_notification_config** (Block List) (see [below for nested schema](#nestedblock--ansible_tower_notification_config))
- **base_notification_config** (Block List) (see [below for nested schema](#nestedblock--base_notification_config))
- **email_notification_config** (Block List) (see [below for nested schema](#nestedblock--email_notification_config))
- **hip_chat_notification_config** (Block List) (see [below for nested schema](#nestedblock--hip_chat_notification_config))
- **id** (String) The ID of this resource.
- **jira_notification_config** (Block List) (see [below for nested schema](#nestedblock--jira_notification_config))
- **ops_genie_notification_config** (Block List) (see [below for nested schema](#nestedblock--ops_genie_notification_config))
- **pager_duty_notification_config** (Block List) (see [below for nested schema](#nestedblock--pager_duty_notification_config))
- **service_now_notification_config** (Block List) (see [below for nested schema](#nestedblock--service_now_notification_config))
- **slack_notification_config** (Block List) (see [below for nested schema](#nestedblock--slack_notification_config))
- **trello_notification_config** (Block List) (see [below for nested schema](#nestedblock--trello_notification_config))
- **victor_ops_notification_config** (Block List) (see [below for nested schema](#nestedblock--victor_ops_notification_config))
- **web_hook_notification_config** (Block List) (see [below for nested schema](#nestedblock--web_hook_notification_config))
- **xmatters_notification_config** (Block List) (see [below for nested schema](#nestedblock--xmatters_notification_config))

<a id="nestedblock--ansible_tower_notification_config"></a>
### Nested Schema for `ansible_tower_notification_config`

Optional:

- **accept_any_certificate** (Boolean)
- **active** (Boolean)
- **alerting_profile** (String)
- **custom_message** (String)
- **id** (String) The ID of this resource.
- **job_template_id** (Number)
- **job_template_url** (String)
- **name** (String)
- **password** (String)
- **type** (String)
- **username** (String)


<a id="nestedblock--base_notification_config"></a>
### Nested Schema for `base_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **id** (String) The ID of this resource.
- **instance_name** (String)
- **message** (String)
- **name** (String)
- **password** (String)
- **send_events** (Boolean)
- **send_incidents** (Boolean)
- **type** (String)
- **url** (String)
- **username** (String)


<a id="nestedblock--email_notification_config"></a>
### Nested Schema for `email_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **bcc_receivers** (List of String)
- **body** (String)
- **cc_receivers** (List of String)
- **id** (String) The ID of this resource.
- **name** (String)
- **receivers** (List of String)
- **subject** (String)
- **type** (String)


<a id="nestedblock--hip_chat_notification_config"></a>
### Nested Schema for `hip_chat_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **id** (String) The ID of this resource.
- **message** (String)
- **name** (String)
- **type** (String)
- **url** (String)


<a id="nestedblock--jira_notification_config"></a>
### Nested Schema for `jira_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **description** (String)
- **id** (String) The ID of this resource.
- **issue_type** (String)
- **name** (String)
- **password** (String)
- **project_key** (String)
- **summary** (String)
- **type** (String)
- **url** (String)
- **username** (String)


<a id="nestedblock--ops_genie_notification_config"></a>
### Nested Schema for `ops_genie_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **api_key** (String)
- **domain** (String)
- **id** (String) The ID of this resource.
- **message** (String)
- **name** (String)
- **type** (String)


<a id="nestedblock--pager_duty_notification_config"></a>
### Nested Schema for `pager_duty_notification_config`

Optional:

- **account** (String)
- **active** (Boolean)
- **alerting_profile** (String)
- **id** (String) The ID of this resource.
- **name** (String)
- **service_api_key** (String)
- **service_name** (String)
- **type** (String)


<a id="nestedblock--service_now_notification_config"></a>
### Nested Schema for `service_now_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **id** (String) The ID of this resource.
- **instance_name** (String)
- **message** (String)
- **name** (String)
- **password** (String)
- **send_events** (Boolean)
- **send_incidents** (Boolean)
- **type** (String)
- **url** (String)
- **username** (String)


<a id="nestedblock--slack_notification_config"></a>
### Nested Schema for `slack_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **channel** (String)
- **id** (String) The ID of this resource.
- **name** (String)
- **title** (String)
- **type** (String)
- **url** (String)


<a id="nestedblock--trello_notification_config"></a>
### Nested Schema for `trello_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **application_key** (String)
- **authorization_token** (String)
- **board_id** (String)
- **description** (String)
- **id** (String) The ID of this resource.
- **list_id** (String)
- **name** (String)
- **resolved_list_id** (String)
- **text** (String)
- **type** (String)


<a id="nestedblock--victor_ops_notification_config"></a>
### Nested Schema for `victor_ops_notification_config`

Optional:

- **active** (Boolean)
- **alerting_profile** (String)
- **api_key** (String)
- **id** (String) The ID of this resource.
- **message** (String)
- **name** (String)
- **routing_key** (String)
- **type** (String)


<a id="nestedblock--web_hook_notification_config"></a>
### Nested Schema for `web_hook_notification_config`

Optional:

- **accept_any_certificate** (Boolean)
- **active** (Boolean)
- **alerting_profile** (String)
- **headers** (Block List) (see [below for nested schema](#nestedblock--web_hook_notification_config--headers))
- **id** (String) The ID of this resource.
- **name** (String)
- **payload** (String)
- **type** (String)
- **url** (String)

<a id="nestedblock--web_hook_notification_config--headers"></a>
### Nested Schema for `web_hook_notification_config.headers`

Optional:

- **name** (String)
- **value** (String)



<a id="nestedblock--xmatters_notification_config"></a>
### Nested Schema for `xmatters_notification_config`

Optional:

- **accept_any_certificate** (Boolean)
- **active** (Boolean)
- **alerting_profile** (String)
- **headers** (Block List) (see [below for nested schema](#nestedblock--xmatters_notification_config--headers))
- **id** (String) The ID of this resource.
- **name** (String)
- **payload** (String)
- **type** (String)
- **url** (String)

<a id="nestedblock--xmatters_notification_config--headers"></a>
### Nested Schema for `xmatters_notification_config.headers`

Optional:

- **name** (String)
- **value** (String)


