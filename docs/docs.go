// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/scheduler/func/name": {
            "get": {
                "description": "获取所有的 func name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Func"
                ],
                "summary": "获取所有的 func name",
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            }
        },
        "/scheduler/job": {
            "get": {
                "description": "根据job_id 获取job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "根据job_id 获取job",
                "parameters": [
                    {
                        "type": "string",
                        "description": " job id",
                        "name": "job_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "修改任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "修改任务",
                "parameters": [
                    {
                        "description": "body",
                        "name": "Data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.Job"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "添加任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "添加任务",
                "parameters": [
                    {
                        "description": "body",
                        "name": "Data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.Job"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "删除任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "job id",
                        "name": "job_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            }
        },
        "/scheduler/job/pause": {
            "put": {
                "description": "暂定任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "暂定任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": " job_id",
                        "name": "job_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            }
        },
        "/scheduler/job/resume": {
            "put": {
                "description": "暂定任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "暂定任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": " job_id",
                        "name": "job_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            }
        },
        "/scheduler/jobs": {
            "get": {
                "description": "根据store_name获取jobs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jobs"
                ],
                "summary": "根据tore_name获取jobs",
                "parameters": [
                    {
                        "type": "string",
                        "description": " store_name",
                        "name": "store_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/view.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/job.Job"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/scheduler/store/name": {
            "get": {
                "description": "获取所有的 store name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Store"
                ],
                "summary": "获取所有的 store name",
                "responses": {
                    "200": {
                        "description": "响应值",
                        "schema": {
                            "$ref": "#/definitions/view.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "job.Job": {
            "type": "object",
            "properties": {
                "args": {
                    "description": "Arguments for ` + "`" + `Func` + "`" + `.",
                    "type": "object",
                    "additionalProperties": {}
                },
                "func_name": {
                    "description": "注册函数名",
                    "type": "string"
                },
                "id": {
                    "description": "任务的唯一id.",
                    "type": "string"
                },
                "max_instance": {
                    "description": "改任务可以同时存在的个数 最少1个, 默认 1",
                    "type": "integer"
                },
                "name": {
                    "description": "job name",
                    "type": "string"
                },
                "next_run_time": {
                    "description": "Automatic update, not manual setting.  ms",
                    "type": "integer"
                },
                "replace": {
                    "description": "任务存在是否更新 默认false",
                    "type": "boolean"
                },
                "status": {
                    "description": "Optional: ` + "`" + `STATUS_RUNNING` + "`" + ` | ` + "`" + `STATUS_PAUSED` + "`" + `\nIt should not be set manually.",
                    "type": "string"
                },
                "store_name": {
                    "description": "jobStoreName",
                    "type": "string"
                },
                "timeout": {
                    "description": "The running timeout of ` + "`" + `Func` + "`" + `.\nms Default: 3600*1000",
                    "type": "integer"
                },
                "trigger": {}
            }
        },
        "triggers.CronTrigger": {
            "type": "object",
            "properties": {
                "Jitter": {
                    "description": "时间误差, 超过这个误差时间就忽略本次执行, 默认 0 表示不管误差, 单位 ms time.Millisecond,",
                    "type": "integer"
                },
                "cron_expr": {
                    "type": "string"
                },
                "end_time": {
                    "description": "数据格式 time.DateTime \"2006-01-02 15:04:05\"",
                    "type": "string"
                },
                "start_time": {
                    "description": "数据格式 time.DateTime \"2006-01-02 15:04:05\"",
                    "type": "string"
                },
                "utc_time_zone": {
                    "description": "默认就是 UTC",
                    "type": "string"
                }
            }
        },
        "triggers.DateTrigger": {
            "type": "object",
            "properties": {
                "Jitter": {
                    "description": "时间误差, 超过这个误差时间就忽略本次执行,默认 0 表示不管误差, 单位 ms time.Millisecond",
                    "type": "integer"
                },
                "run_date": {
                    "description": "数据格式 time.DateTime \"2006-01-02 15:04:05\"",
                    "type": "string"
                },
                "utc_time_zone": {
                    "description": "默认UTC",
                    "type": "string"
                }
            }
        },
        "triggers.IntervalTrigger": {
            "type": "object",
            "properties": {
                "Jitter": {
                    "description": "时间误差, 超过这个误差时间就忽略本次执行, 默认 0 表示不管误差, 单位 ms time.Millisecond",
                    "type": "integer"
                },
                "end_time": {
                    "description": "数据格式 time.DateTime  \"2006-01-02 15:04:05\"",
                    "type": "string"
                },
                "interval": {
                    "description": "单位 ms",
                    "type": "integer"
                },
                "start_time": {
                    "description": "数据格式 time.DateTime  \"2006-01-02 15:04:05\"",
                    "type": "string"
                },
                "utc_time_zone": {
                    "description": "默认UTC",
                    "type": "string"
                }
            }
        },
        "view.HttpResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "errorCode": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "view.Job": {
            "type": "object",
            "properties": {
                "args": {
                    "description": "Arguments for ` + "`" + `Func` + "`" + `.",
                    "type": "object",
                    "additionalProperties": {}
                },
                "cron_trigger": {
                    "$ref": "#/definitions/triggers.CronTrigger"
                },
                "date_trigger": {
                    "description": "同时出现的时候优先使用 interval_trigger \u003e date_trigger \u003e cron_trigger",
                    "allOf": [
                        {
                            "$ref": "#/definitions/triggers.DateTrigger"
                        }
                    ]
                },
                "func_name": {
                    "description": "必须和注册的函数名一致",
                    "type": "string"
                },
                "id": {
                    "description": "任务的唯一id.",
                    "type": "string"
                },
                "interval_trigger": {
                    "description": "cron_trigger, date_trigger,interval_trigger 只能有一个生效",
                    "allOf": [
                        {
                            "$ref": "#/definitions/triggers.IntervalTrigger"
                        }
                    ]
                },
                "max_instance": {
                    "description": "改任务可以同时存在的个数 最少1个, 默认 1",
                    "type": "integer"
                },
                "name": {
                    "description": "job name",
                    "type": "string"
                },
                "old_store_name": {
                    "description": "更新的时候需要",
                    "type": "string"
                },
                "replace": {
                    "description": "任务存在是否更新 默认false",
                    "type": "boolean"
                },
                "store_name": {
                    "description": "jobStoreName",
                    "type": "string"
                },
                "timeout": {
                    "description": "任务执行超时时间 ms Default: ` + "`" + `3600*1000` + "`" + `",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
